<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="课程类型">
                    <el-input v-model.number="searchInfo.courseType" placeholder="请输入班级编号" />
                </el-form-item>
                <el-form-item label="课程名称">
                    <el-input v-model="searchInfo.courseName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item>
                    <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table
                ref="multipleTable"
                border
                stripe
                style="width: 100%"
                :header-cell-style="{color: '#224b8f',fontFamily:'MicrosoftYaHeiUI',fontSize:'16px',fontWeight:900}"
                tooltip-effect="dark"
                :data="tableData"
                @selection-change="handleSelectionChange">
            <el-table-column label="日期" prop="CreatedAt" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="课程类型" prop="courseType" align="center"/>
            <el-table-column label="课程名称" prop="courseName" align="center"/>
            <el-table-column label="源码上传" align="center" width="500">
                <template slot-scope="scope">
                    <el-upload
                            class="avatar-uploader"
                            ref="uploadSourceCode"
                            action=""
                            accept=".sb3, .py, .cpp"
                            :before-upload="beforeUploadSourceCode"
                            :show-file-list="false"
                            :on-progress="uploadSourceCode"
                            :auto-upload="false">
                        <el-button slot="trigger" size="mini" type="primary">选取文件</el-button>
                        <el-button v-if="scope.row.isUploadSourceCode === 0" style="margin-left: 10px;" size="mini" type="success" @click="submitUploadSourceCode(scope.row)">上传到服务器</el-button>
                        <el-button v-else-if="scope.row.isUploadSourceCode === 1" style="margin-left: 10px;" size="mini" type="info" @click="submitUploadSourceCode(scope.row)">重新上传</el-button>
                        <el-button v-else style="margin-left: 10px;" size="mini" type="danger">文件状态异常</el-button>
                        <div slot="tip" class="el-upload__tip">上传本节课素材压缩包，仅支持sb3/py/cpp格式</div>
                    </el-upload>
                </template>
            </el-table-column>
            <el-table-column label="素材上传" align="center" width="500">
                <template slot-scope="scope">
                    <el-upload
                            class="avatar-uploader"
                            ref="uploadFodder"
                            action=""
                            accept=".rar, .zip, .7z, .tar"
                            :before-upload="beforeUploadFodder"
                            :show-file-list="false"
                            :on-progress="uploadFodder"
                            :auto-upload="false">
                        <el-button slot="trigger" size="mini" type="primary">选取文件</el-button>
                        <el-button v-if="scope.row.isUploadFodder === 0" style="margin-left: 10px;" size="mini" type="success" @click="submitUploadFodder(scope.row)">上传到服务器</el-button>
                        <el-button v-else-if="scope.row.isUploadFodder === 1" style="margin-left: 10px;" size="mini" type="info" @click="submitUploadFodder(scope.row)">重新上传</el-button>
                        <el-button v-else style="margin-left: 10px;" size="mini" type="danger">文件状态异常</el-button>
                        <div slot="tip" class="el-upload__tip">上传本节课素材压缩包，仅支持rar/zip/7z/tar格式</div>
                    </el-upload>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination
                layout="total, sizes, prev, pager, next, jumper"
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"/>
    </div>
</template>

<script>
    import {
        getTmpSecret,
        updateUpload,
        findUpload,
        getUploadList,
    } from '@/api/z_upload'
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    import COS from 'cos-js-sdk-v5'
    export default {
        name: 'Upload',
        mixins: [infoList],
        data() {
            return {
                listApi: getUploadList,
                dialogFormVisible: false,
                type: '',
                fodderFile: null,
                sourceCodeFile: null,

                formData: {
                    CreatedAt: null,
                    courseType: null,
                    courseName: null,
                    bucket: null,
                    region: null,
                    isUploadSourceCode: null,
                    sourceCodeFolder: null,
                    sourceCodeUrl: null,
                    isUploadFodder: null,
                    fodderFolder: null,
                    fodderUrl: null,
                },
            }
        },
        async created() {
            await this.getTableData()
        },
        methods: {
            async submitUploadFodder(row) {
                await this.$refs.uploadFodder.submit();
                if (this.fodderFile === null) {
                    this.$message({
                        type: 'warning',
                        message: '请先选择文件后再上传'
                    });
                    return
                }
                row.isUploadFodder = 1;
                row.fodderUrl =  'https://' + row.bucket + '.cos.' + row.region + '.myqcloud.com/fodder/' + this.fodderFile.name;
                let res = await updateUpload(row);
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '上传素材成功'
                    });
                    //this.fodderFile = null
                }
            },
            async submitUploadSourceCode(row) {
                await this.$refs.uploadSourceCode.submit();
                if (this.sourceCodeFile === null) {
                    this.$message({
                        type: 'warning',
                        message: '请先选择文件后再上传'
                    });
                    return
                }
                row.isUploadSourceCode = 1;
                row.sourceCodeUrl =  'https://' + row.bucket + '.cos.' + row.region + '.myqcloud.com/source-code/' + this.sourceCodeFile.name;
                let res = await updateUpload(row);
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '上传源码成功'
                    });
                    //this.sourceCodeFile = null
                }
            },
            async uploadSourceCode() {
                let res = await getTmpSecret();
                if (res.code !== 0 ) {return console.error('credentials invalid')}
                let cos = new COS({
                    SecretId: res.data.Credentials.TmpSecretId,
                    SecretKey: res.data.Credentials.TmpSecretKey,
                    SecurityToken: res.data.Credentials.Token,
                    StartTime: res.data.StartTime,
                    ExpiredTime: res.data.ExpiredTime,
                });
                cos.putObject({
                    Bucket: 'saisicode-1304003768',
                    Region: 'ap-chengdu',
                    Key: 'source-code/' + this.sourceCodeFile.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: this.sourceCodeFile, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, function(err, data) {
                    //console.log(err || data);
                });
            },
            async uploadFodder() {
                let res = await getTmpSecret();
                if (res.code !== 0 ) {return console.error('credentials invalid')}
                let cos = new COS({
                    SecretId: res.data.Credentials.TmpSecretId,
                    SecretKey: res.data.Credentials.TmpSecretKey,
                    SecurityToken: res.data.Credentials.Token,
                    StartTime: res.data.StartTime,
                    ExpiredTime: res.data.ExpiredTime,
                });
                cos.putObject({
                    Bucket: 'saisicode-1304003768',
                    Region: 'ap-chengdu',
                    Key: 'fodder/' + this.fodderFile.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: this.fodderFile, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, function(err, data) {
                    //console.log(err || data);
                });
            },
            beforeUploadFodder (file) {
                this.fodderFile = file
                console.log(file.name)
            },
            beforeUploadSourceCode (file) {
                this.sourceCodeFile = file
                console.log(file.name)
            },
            formatDate(row) {
                return formatTimeToStr(row.CreatedAt, 'yyyy-MM-dd');
            },
            // 条件搜索前端看此方法
            onSubmit() {
                this.page = 1;
                this.pageSize = 10;
                this.getTableData()
            },
            handleSelectionChange(val) {
                this.multipleSelection = val
            },
        },
    }
</script>

<style>
</style>


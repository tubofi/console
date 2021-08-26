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
            <el-table-column label="源码上传" align="center">
                <template slot-scope="scope">
                    <el-upload
                            class="avatar-uploader"
                            action=""
                            accept=".sb3, .py, .cpp"
                            :http-request="uploadSourceCode">
                        <el-button v-if="scope.row.isUploadSourceCode === 0" size="mini" type="primary" @click="putFormData(scope.row)">上传源码</el-button>
                        <el-button v-else-if="scope.row.isUploadSourceCode === 1" size="mini" type="warning" @click="putFormData(scope.row)">重新上传</el-button>
                        <el-button v-else size="mini" type="danger">状态异常</el-button>
                        <div slot="tip" class="el-upload__tip">上传本节课源码，仅支持sb3/py/cpp格式</div>
                    </el-upload>
                </template>
            </el-table-column>
            <el-table-column label="素材上传" align="center">
                <template slot-scope="scope">
                    <el-upload
                            class="avatar-uploader"
                            action=""
                            accept=".rar, .zip, .7z, .tar"
                            :http-request="uploadFodder">
                        <el-button v-if="scope.row.isUploadFodder === 0" size="mini" type="primary" @click="putFormData(scope.row)">上传素材</el-button>
                        <el-button v-else-if="scope.row.isUploadFodder === 1" size="mini" type="warning" @click="putFormData(scope.row)">重新上传</el-button>
                        <el-button v-else size="mini" type="danger">状态异常</el-button>
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
                formData: null,

                /*
                formData: {
                    ID: null,
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
                    comment: null,
                },*/
            }
        },
        async created() {
            await this.getTableData()
        },
        methods: {
            putFormData(row){
                this.formData = row;
            },
            async uploadSourceCode(params){
                let row = this.formData;
                let res = await getTmpSecret();
                if (res.code !== 0 ) {console.error('credentials invalid'); return}
                let that = this;            //子函数的this不能指向vue对象的api
                let cos = new COS({
                    SecretId: res.data.Credentials.TmpSecretId,
                    SecretKey: res.data.Credentials.TmpSecretKey,
                    SecurityToken: res.data.Credentials.Token,
                    StartTime: res.data.StartTime,
                    ExpiredTime: res.data.ExpiredTime,
                });
                cos.putObject({
                    Bucket: row.bucket,
                    Region: row.region,
                    Key: row.sourceCodeFolder + params.file.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: params.file, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, async function(err, data) {
                    if (err === null) {
                        row.isUploadSourceCode = 1;
                        row.sourceCodeUrl = data.Location
                    } else {
                        that.$message({
                            type: 'danger',
                            message: '素材上传失败！'
                        });return
                    }
                    let response = await updateUpload(row);
                    if (response.code === 0) {
                        that.$message({
                            type: 'success',
                            message: '源码上传成功！'
                        });
                        that.formData = null
                    }
                });
            },
            async uploadFodder(params){
                let row = this.formData;
                let res = await getTmpSecret();
                if (res.code !== 0 ) {console.error('credentials invalid'); return}
                let that = this;            //子函数的this不能指向vue对象的api
                let cos = new COS({
                    SecretId: res.data.Credentials.TmpSecretId,
                    SecretKey: res.data.Credentials.TmpSecretKey,
                    SecurityToken: res.data.Credentials.Token,
                    StartTime: res.data.StartTime,
                    ExpiredTime: res.data.ExpiredTime,
                });
                cos.putObject({
                    Bucket: row.bucket,
                    Region: row.region,
                    Key: row.fodderFolder + params.file.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: params.file, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, async function(err, data) {
                    if (err === null) {
                        row.isUploadFodder = 1;
                        row.fodderUrl = data.Location
                    } else {
                        that.$message({
                            type: 'danger',
                            message: '素材上传失败！'
                        });return
                    }
                    let response = await updateUpload(row);
                    if (response.code === 0) {
                        that.$message({
                            type: 'success',
                            message: '素材上传成功！'
                        });
                        that.formData = null
                    }
                });
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


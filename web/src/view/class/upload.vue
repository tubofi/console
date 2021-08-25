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
            <el-table-column label="源码上传" align="center" width="300">
                <template slot-scope="scope">
                    <div v-if="scope.row.isUploadSourceCode !== 1">
                        <el-upload
                                class="avatar-uploader"
                                action="#999"
                                accept=".sb3, .py, .cpp"
                                :before-upload="beforeAvatarUpload"
                                :show-file-list="false"
                                :on-progress="uploadSourceCode(scope.row)"
                                :auto-upload="false"
                                :on-success = "finishSourceCode(scope.row)">
                            <el-button size="mini" type="primary">点击上传</el-button>
                            <div slot="tip" class="el-upload__tip">上传本节课源程序</div>
                        </el-upload>
                    </div>
                    <div v-else>
                        <el-upload
                                class="avatar-uploader"
                                action="#999"
                                accept=".sb3, .py, .cpp"
                                :before-upload="beforeAvatarUpload"
                                :show-file-list="false"
                                :on-progress="uploadSourceCode(scope.row)"
                                :auto-upload="false"
                                :on-success = "finishSourceCode(scope.row)">
                            <el-button size="mini" type="info">重新上传</el-button>
                            <div slot="tip" class="el-upload__tip">程序已上传，需要重新上传请点击</div>
                        </el-upload>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="源码上传" align="center" width="300">
                <template slot-scope="scope">
                    <div v-if="scope.row.isUploadFodder !== 1">
                        <el-upload
                                class="avatar-uploader"
                                action="#999"
                                accept=".rar, .zip, .7z, .tar"
                                :before-upload="beforeAvatarUpload"
                                :show-file-list="false"
                                :on-progress="uploadFodder(scope.row)"
                                :auto-upload="false"
                                :on-success = "finishFodder(scope.row)">
                            <el-button size="mini" type="primary">点击上传</el-button>
                            <div slot="tip" class="el-upload__tip">上传本节课素材压缩包</div>
                        </el-upload>
                    </div>
                    <div v-else>
                        <el-upload
                                class="avatar-uploader"
                                action="#999"
                                accept=".rar, .zip, .7z, .tar"
                                :before-upload="beforeAvatarUpload"
                                :show-file-list="false"
                                :on-progress="uploadFodder(scope.row)"
                                :auto-upload="false"
                                :on-success = "finishFodder(scope.row)">
                            <el-button size="mini" type="info">重新上传</el-button>
                            <div slot="tip" class="el-upload__tip">素材已上传，需要重新上传请点击</div>
                        </el-upload>
                    </div>
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
            async uploadSourceCode(row) {
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
                    Bucket: row.bucket,
                    Region: row.region,
                    Key: row.sourceCodeFolder + this.uploadFile.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: this.uploadFile, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, function(err, data) {
                    console.log(err || data);
                });
            },
            async uploadFodder(row) {
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
                    Bucket: row.bucket,
                    Region: row.region,
                    Key: row.fodderFolder + this.uploadFile.name,              /* 必须 */
                    StorageClass: 'STANDARD',
                    Body: this.uploadFile, // 上传文件对象
                    onProgress: function(progressData) {
                        console.log(JSON.stringify(progressData));
                    }
                }, function(err, data) {
                    console.log(err || data);
                });
            },
            beforeAvatarUpload (file) {
                this.uploadFile = file
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
        },
    }
</script>

<style>
</style>


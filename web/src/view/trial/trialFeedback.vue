<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="学生姓名">
                    <el-input v-model="searchInfo.studentName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="教师姓名">
                    <el-input v-model="searchInfo.teacherName" placeholder="请输入模糊关键词" />
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
            <el-table-column label="日期" prop="time" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="学生姓名" prop="studentName" align="center"/>
            <el-table-column label="课程类型" prop="courseType" align="center"/>
            <el-table-column label="课程名称" prop="courseName" align="center"/>
            <el-table-column label="授课教师" prop="teacherName" align="center"/>
            <el-table-column label="教室" prop="room" align="center"/>
            <el-table-column label="反馈详情" align="center" width="100">
                <template slot-scope="scope">
                    <div>
                        <el-popover v-if="scope.row" placement="top-start" trigger="hover">
                            <div class="popover-box">
                                <pre>总分：{{scope.row.total}}</pre>
                                <pre>守时：{{scope.row.punctuality}}</pre>
                                <pre>纪律：{{scope.row.discipline}}</pre>
                                <pre>专注：{{scope.row.concentration}}</pre>
                                <pre>创新：{{scope.row.innovation}}</pre>
                                <pre>逻辑：{{scope.row.logic}}</pre>
                                <pre>数学：{{scope.row.mathematics}}</pre>
                                <pre>独立：{{scope.row.complete}}</pre>
                            </div>
                            <i slot="reference" class="el-icon-view" />
                        </el-popover>
                        <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <div v-if="scope.row.needFeedback === 1 && scope.row.total === 0">
                        <el-button plain type="danger" icon="el-icon-star-off" size="mini" @click="updateTrialCourseRecord(scope.row)">反馈</el-button>
                    </div>
                    <div v-else-if="scope.row.needFeedback === 0 && scope.row.total > 0">
                        <el-button plain size="mini" type="primary" icon="el-icon-edit" @click="updateTrialCourseRecord(scope.row)">编辑</el-button>
                    </div>
                    <div v-else>
                        <el-tag type="danger" size="medium">异常</el-tag>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="试听反馈表单" width="40%">
            <el-form ref="formData" :model="formData" label-position="right" label-width="100px">
                <el-form-item label="学生姓名:" prop="studentName" >
                    <el-input :disabled="true" v-model="formData.studentName" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="课程名称:" prop="courseName">
                    <el-input :disabled="true" v-model="formData.courseName" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="授课教师:" prop="teacherName">
                    <el-input :disabled="true" v-model="formData.teacherName" clearable placeholder="请输入"/>
                </el-form-item>
            </el-form>
            <el-form ref="rateForm" :model="formData" :rules="rules" label-width="0px">
                <el-row class="row-card" type="flex" justify="space-between">
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">守时(5%)：</span>
                            </div>
                            <el-form-item prop="punctuality" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.punctuality" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">根据迟早、早退、请假具体情况扣分。因病或其它充分理由不扣分。</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">纪律(10%)：</span>
                            </div>
                            <el-form-item prop="discipline" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.discipline" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">课堂上嬉戏打闹、大声喧哗、吃零食等扰乱课堂秩序的行为均为扣分项。</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">专注(15%)：</span>
                            </div>
                            <el-form-item prop="concentration" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.concentration" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">反应课堂参与度和积极性。是否紧跟老师节奏？是否积极思考？是否踊跃回答问题？</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">创新(15%)：</span>
                            </div>
                            <el-form-item prop="innovation" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.innovation" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">对关键知识点提出质疑；对目标任务有不一样的实现方法；对作品内容提出创新性的改进。</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                </el-row>
                <el-row class="row-card" type="flex" justify="space-between">
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">逻辑(20%)：</span>
                            </div>
                            <el-form-item prop="logic" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.logic" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">是否明确课堂目标？对围绕目标开展的各项子任务以及子任务之间的关系是否清晰？执行是否有条理？对程序的理解程度如何？</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">数学(20%)：</span>
                            </div>
                            <el-form-item prop="mathematics" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.mathematics" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">对作品涉及数学知识的掌握程度，对程序算法的理解程度</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">独立(15%)：</span>
                            </div>
                            <el-form-item prop="complete" style="margin-bottom: 0" required>
                                <el-rate v-model.number="formData.complete" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">反应独立完成作品能力和动手能力，自己探索的部分有多少？是不是跟着老师做下来的？</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-card shadow="hover" :body-style="{padding: '10px'}">
                            <div slot="header">
                                <span style="font-size: 16px">备用：</span>
                            </div>
                            <el-form-item prop="total" style="margin-bottom: 0">
                                <el-rate v-model.number="formData.total" :colors="colors" :allow-half="true" :show-text="true"/>
                                <span class="card-description">...</span>
                            </el-form-item>
                        </el-card>
                    </el-col>
                </el-row>
                <el-form-item prop="review">
                    <el-input style="margin-top: 10px" type="textarea" :rows="5" v-model="formData.review" clearable :placeholder=remark />
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="mini" @click="closeDialog">取 消</el-button>
                <el-button type="primary" size="mini" @click="beforeEnterDialog('rateForm')">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {
        createTrialCourseRecord,
        deleteTrialCourseRecord,
        deleteTrialCourseRecordByIds,
        updateTrialCourseRecord,
        findTrialCourseRecord,
        getTrialCourseRecordList,
        getAllTrialStudents,
        feedbackTrialCourseRecord,
    } from '@/api/z_trial' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'FeedbackTrialCourseRecord',
        mixins: [infoList],
        data() {
            return {
                listApi: getTrialCourseRecordList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                colors: ['#99A9BF', '#F7BA2A', '#FF9900'],  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
                remark: "结合以上要素对学生的课堂表现给出客观文字点评：\n",

                formData: {
                    studentName: null,
                    teacherName: null,
                    courseType: null,
                    courseName: null,
                    time: new Date(),
                    room: null,
                    needFeedback: null,
                    total: null,
                    punctuality: null,
                    discipline: null,
                    concentration: null,
                    innovation: null,
                    logic: null,
                    mathematics: null,
                    complete: null,
                    review: null,
                    comment: null,
                },
                rules: {
                    punctuality: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    discipline: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    concentration: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    innovation: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    logic: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    mathematics: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    complete: [{type: "number", min: 0.1, max: 5, message: '请为该项评分', trigger: 'blur' }],
                    review: [{ required: true, message: '请为学生课堂表现做出点评', trigger: 'blur' }],
                },
            }
        },
        filters: {
            formatDate: function(time) {
                if (time !== null && time !== '') {
                    let date = new Date(time);
                    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
                } else {
                    return ''
                }
            },
            formatBoolean: function(bool) {
                if (bool != null) {
                    return bool ? '是' : '否'
                } else {
                    return ''
                }
            }
        },
        async created() {
            await this.getTableData()
        },
        methods: {
            beforeEnterDialog(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.enterDialog()
                    }
                })
            },
            formatDate(row) {
                return formatTimeToStr(row.CreatedAt, 'yyyy-MM-dd');
            },
            // 条件搜索前端看此方法
            onSubmit() {
                this.page = 1
                this.pageSize = 10
                this.getTableData()
            },
            handleSelectionChange(val) {
                this.multipleSelection = val
            },
            deleteRow(row) {
                this.$confirm('确定要删除吗?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.deleteTrialCourseRecord(row)
                })
            },
            async onDelete() {
                const ids = []
                if (this.multipleSelection.length === 0) {
                    this.$message({
                        type: 'warning',
                        message: '请选择要删除的数据'
                    })
                    return
                }
                this.multipleSelection &&
                this.multipleSelection.map(item => {
                    ids.push(item.ID)
                })
                const res = await deleteTrialCourseRecordByIds({ ids })
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功'
                    })
                    if (this.tableData.length === ids.length && this.page > 1) {
                        this.page--
                    }
                    this.deleteVisible = false
                    this.getTableData()
                }
            },
            async updateTrialCourseRecord(row) {
                const res = await findTrialCourseRecord({ ID: row.ID })
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.retrialCourseRecord
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    studentName: null,
                    teacherName: null,
                    courseType: null,
                    courseName: null,
                    time: new Date(),
                    room: null,
                    needFeedback: null,
                    total: null,
                    punctuality: null,
                    discipline: null,
                    concentration: null,
                    innovation: null,
                    logic: null,
                    mathematics: null,
                    complete: null,
                    review: null,
                    comment: null,
                }
            },
            async deleteTrialCourseRecord(row) {
                const res = await deleteTrialCourseRecord({ ID: row.ID })
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功'
                    })
                    if (this.tableData.length === 1 && this.page > 1 ) {
                        this.page--
                    }
                    this.getTableData()
                }
            },
            async enterDialog() {
                let res
                switch (this.type) {
                    case "create":
                        res = await createTrialCourseRecord(this.formData)
                        break
                    case "update":
                        res = await feedbackTrialCourseRecord(this.formData)
                        break
                    default:
                        res = await createTrialCourseRecord(this.formData)
                        break
                }
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '创建/更改成功'
                    })
                    this.closeDialog()
                    this.getTableData()
                }
            },
            openDialog() {
                this.type = 'create'
                this.dialogFormVisible = true
            }
        },
    }
</script>

<style>
    .card-description {
        display: block;
        height: 50px;
        font-size: 10px;
        color: #5e6d82;
        margin-top: 10px;
        line-height: 15px;
    }
    .row-card {
        width: 100%;
        padding: 0 0
    }
</style>


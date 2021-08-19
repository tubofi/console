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
            <el-table-column label="日期" prop="CreatedAt" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="学生姓名" prop="studentName" align="center"/>
            <el-table-column label="班级ID" prop="classId" align="center"/>
            <el-table-column label="课程ID" prop="courseId" align="center"/>
            <el-table-column label="课程阶段" prop="courseContent" align="center"/>
            <el-table-column label="课程名称" prop="courseName" align="center"/>
            <el-table-column label="授课教师" prop="teacherName" align="center"/>
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <div v-if="scope.row.needFeedback === 1">
                        <el-button plain type="danger" icon="el-icon-star-off" size="mini" @click="updateCourseRecord(scope.row)">反馈</el-button>
                    </div>
                    <div v-else-if="scope.row.needFeedback === 0 && scope.row.total === 0">
                        <el-tag type="info" size="medium">已忽略</el-tag>
                    </div>
                    <div v-else>
                        <el-button plain size="mini" type="primary" icon="el-icon-edit" @click="updateCourseRecord(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="反馈表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="学生姓名:" prop="studentName" >
                    <el-input :disabled="true" v-model="formData.studentName" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="课程名称:" prop="courseName">
                    <el-input :disabled="true" v-model="formData.courseName" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="授课教师:" prop="teacherName">
                    <el-input :disabled="true" v-model="formData.teacherName" clearable placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="守时:" prop="punctuality">
                    <el-rate v-model.number="formData.punctuality" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <!---<el-form-item label="纪律:" prop="discipline">
                    <el-rate v-model.number="formData.discipline" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="专注:" prop="concentration">
                    <el-rate v-model.number="formData.concentration" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="创新:" prop="innovation">
                    <el-rate v-model.number="formData.innovation" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="逻辑:" prop="logic">
                    <el-rate v-model.number="formData.logic" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="数学:" prop="mathematics">
                    <el-rate v-model.number="formData.mathematics" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="完成:" prop="complete">
                    <el-rate v-model.number="formData.complete" :colors="colors" :allow-half="true" :show-text="true"/>
                </el-form-item>
                <el-form-item label="教师点评:" prop="review">
                    <el-input type="textarea" :rows="5" v-model="formData.review" clearable placeholder="课堂点评"/>
                </el-form-item>--->
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button type="primary" @click="beforeEnterDialog('formData')">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    const absentOptions = [
        {label: '正课', value: 0},
        {label: '补课', value: 1},
    ];
    import {
        createCourseRecord,
        deleteCourseRecord,
        deleteCourseRecordByIds,
        updateCourseRecord,
        findCourseRecord,
        getCourseRecordList,
        getCramCourseRecordList,
        cramCourseRecord,
        getFeedbackCourseRecordList,
        feedbackCourseRecord
    } from '@/api/z_course_record' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'CourseRecord',
        mixins: [infoList],
        data() {
            return {
                listApi: getFeedbackCourseRecordList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                absentOptions: absentOptions,
                colors: ['#99A9BF', '#F7BA2A', '#FF9900'],  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }

                formData: {
                    CreatedAt: new Date(),
                    classId: null,
                    courseId: null,
                    studentName: null,
                    teacherName: null,
                    courseContent: null,
                    courseName: null,
                    isAbsentee: null,
                    attendTime: new Date(),
                    needCram: null,
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
                    discipline: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
                    concentration: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
                    innovation: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
                    logic: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
                    mathematics: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
                    complete: [{ required: true, message: '请为该项评分', trigger: 'blur' }],
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
            fmtBody(value) {
                try {
                    return JSON.parse(value)
                } catch (err) {
                    return value
                }
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
                    this.deleteCourseRecord(row)
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
                const res = await deleteCourseRecordByIds({ ids })
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
            async updateCourseRecord(row) {
                const res = await findCourseRecord({ ID: row.ID })
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.recourseRecord
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    CreatedAt: new Date(),
                    classId: null,
                    courseId: null,
                    studentName: null,
                    teacherName: null,
                    courseContent: null,
                    courseName: null,
                    isAbsentee: null,
                    attendTime: new Date(),
                    needCram: null,
                    needFeedback: null,
                    total: null,
                    punctuality: null,
                    discipline: null,
                    concentration: null,
                    innovation: null,
                    logic: null,
                    mathematics: null,
                    complete: null,
                    comment: null,
                }
            },
            async deleteCourseRecord(row) {
                const res = await deleteCourseRecord({ ID: row.ID })
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
                        res = await createCourseRecord(this.formData)
                        break
                    case "update":
                        res = await feedbackCourseRecord(this.formData)
                        break
                    default:
                        res = await createCourseRecord(this.formData)
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
</style>


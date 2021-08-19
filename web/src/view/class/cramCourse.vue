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
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <div v-if="scope.row.needCram === 1">
                        <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateCourseRecord(scope.row)">补课</el-button>
                        <el-button type="warning" icon="el-icon-delete" size="mini" @click="ignoreRow(scope.row)">忽略</el-button>
                    </div>
                    <div v-else-if="scope.row.needFeedback === 0 && scope.row.total === 0">
                        <el-tag type="info" icon="el-icon-close" size="mini">已忽略</el-tag>
                    </div>
                    <div v-else>
                        <el-tag type="success" icon="el-icon-check" size="mini">已完成</el-tag>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="补课表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="学生姓名:">
                    <el-input :disabled="true" v-model="formData.studentName" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="课程名称:">
                    <el-input :disabled="true" v-model="formData.courseName" placeholder="请输入" />
                </el-form-item>
                <el-form-item label="授课教师:">
                    <el-input v-model="formData.teacherName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="上课时间:">
                    <el-date-picker
                            v-model="formData.attendTime"
                            type="datetime"
                            placeholder="选择日期时间">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
                    <el-input type="textarea" :rows="3" v-model="formData.comment" clearable placeholder="备注信息"/>
                </el-form-item>
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
        ignoreCramCourseRecord,
    } from '@/api/z_course_record' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'CourseRecord',
        mixins: [infoList],
        data() {
            return {
                listApi: getCramCourseRecordList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                absentOptions: absentOptions,

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
                    comment: null,
                },
                rules: {
                    teacherName: [{ required: true, message: '教师姓名不能为空', trigger: 'blur' }],
                    attendTime: [{ required: true, message: '请选择补课时间', trigger: 'blur' }],
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
            ignoreRow(row) {
                this.$confirm('确定要不需要补课了吗?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    const res = ignoreCramCourseRecord({ID: row.ID})
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '忽略成功'
                        })
                        this.getTableData()
                    }
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
                        res = await cramCourseRecord(this.formData)
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


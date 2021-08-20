<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="学生姓名">
                    <el-input v-model="searchInfo.studentName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="课程名称">
                    <el-input v-model="searchInfo.courseName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="参课形式">
                    <el-select v-model.number="searchInfo.isAbsentee" placeholder="请选择正课或补课">
                        <el-option
                                v-for="item in absentOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
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
            <el-table-column label="是否参加" prop="isAbsentee" align="center">
                <template slot-scope="scope">
                    <el-tag size="medium" type="warning" v-if="scope.row.isAbsentee === 1">未参加</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="是否补课" prop="needAttend" align="center">
                <template slot-scope="scope">
                    <div v-if="scope.row.isAbsentee === 1">
                        <el-tag size="medium" v-if="scope.row.needCram === 1" type="danger">未补课</el-tag>
                        <el-tag v-else size="medium" type="success">已补课</el-tag>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="补课时间" prop="attendTime" align="center">
                <template slot-scope="scope">
                    <div v-if="scope.row.isAbsentee === 1 && scope.row.needCram === 0">{{formatTime(scope.row.attendTime)}}</div>
                </template>
            </el-table-column>
            <el-table-column label="是否反馈" prop="needFeedback" align="center">
                <template slot-scope="scope">
                    <el-tag size="medium" v-if="scope.row.needFeedback === 1 && scope.row.needCram === 0" type="warning">待反馈</el-tag>
                    <el-tag v-else-if="scope.row.needFeedback === 0 && scope.row.total > 0" size="medium" type="success">已反馈</el-tag>
                    <el-tag v-else-if="scope.row.needFeedback === 0 && scope.row.total === 0" size="medium" type="info">已忽略</el-tag>
                    <el-tag v-else size="medium" type="danger">异常</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="授课教师" prop="teacherName" align="center"/>
            <el-table-column label="反馈详情" align="center" width="100">
                <template slot-scope="scope">
                    <div>
                        <el-popover v-if="scope.row" placement="top-start" trigger="hover">
                            <div class="popover-box">
                                <pre>{{ fmtBody(scope.row) }}</pre>
                            </div>
                            <i slot="reference" class="el-icon-view" />
                        </el-popover>
                        <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <!---<el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateCourseRecord(scope.row)">变更</el-button>
                    <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
                </template>
            </el-table-column>--->
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新班级表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="班级ID:">
                    <el-input v-model.number="formData.classId" clearable placeholder="请输入班级ID" />
                </el-form-item>
                <el-form-item label="学生姓名:">
                    <el-input v-model="formData.studentName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="教师:">
                    <el-input v-model="formData.teacherName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="课程阶段:">
                    <el-input v-model="formData.courseContent" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="课程名称:">
                    <el-input v-model="formData.courseName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="参课类型:">
                    <el-input v-model.number="formData.isAbsentee" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="需要补课:">
                    <el-input v-model.number="formData.needCram" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="是否反馈:">
                    <el-input v-model.number="formData.needFeedback" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
                    <el-input type="textarea" :rows="3" v-model.number="formData.comment" clearable placeholder="备注信息"/>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="mini" @click="closeDialog">取 消</el-button>
                <el-button type="primary" size="mini" @click="enterDialog">确 定</el-button>
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
        getCourseRecordList
    } from '@/api/z_course_record' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'CourseRecord',
        mixins: [infoList],
        data() {
            return {
                listApi: getCourseRecordList,
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
                    studentName: [{ required: true, message: '学生姓名不能为空', trigger: 'blur' }],
                    teacherName: [{ required: true, message: '教师姓名不能为空', trigger: 'blur' }],
                    courseName: [{ required: true, message: '课程名称不能为空', trigger: 'blur' }],
                },
            }
        },
        filters: {
            formatDate: function(time) {
                if (time !== null && time !== '') {
                    var date = new Date(time);
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
            formatTime(time) {
                return formatTimeToStr(time, 'yyyy-MM-dd hh:mm:ss');
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
                        res = await updateCourseRecord(this.formData)
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


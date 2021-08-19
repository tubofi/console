
<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="班号">
                    <el-input v-model.number="searchInfo.classId" placeholder="请输入班级编号" />
                </el-form-item>
                <el-form-item label="时间">
                    <el-input v-model="searchInfo.timeString" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="课名">
                    <el-input v-model="searchInfo.name" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="教室">
                    <el-input v-model.number="searchInfo.room" placeholder="请输入教室编号" />
                </el-form-item>
                <el-form-item label="教师">
                    <el-input v-model="searchInfo.teacherName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="参课">
                    <el-input v-model="searchInfo.students" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="请假">
                    <el-input v-model="searchInfo.absentStudents" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item>
                    <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
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
            <el-table-column label="班级ID" prop="classId" align="center" width="100"/>
            <el-table-column label="课程阶段" prop="courseContent" align="center"/>
            <el-table-column label="上课时间" prop="timeString" align="center"/>
            <el-table-column label="课程名称" prop="name" align="center"/>
            <el-table-column label="占用教室" prop="room" align="center"/>
            <el-table-column label="授课教师" prop="teacherName" align="center"/>
            <el-table-column label="参课学员" prop="students" align="center">
                <template slot-scope="scope">
                    <span v-for="item in scope.row.students">
                        <el-tag size="medium">{{item}}</el-tag>
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="请假学员" prop="absentStudents" align="center">
                <template slot-scope="scope" v-if="scope.row.absentStudents">
                    <span v-for="item in scope.row.absentStudents">
                        <el-tag type="warning" size="medium">{{item}}</el-tag>
                    </span>
                </template>
            </el-table-column>

            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateCourse(scope.row)">编辑</el-button>
                    <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="课程登记表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="登记班级:" prop="classId">
                    <el-select v-model.number="formData.classId" clearable placeholder="请选择班级" @change="autoFill">
                        <el-option
                                v-for="item in classOptions"
                                :key="item.ID"
                                :label="item.courseContent + ' ' + item.timeString"
                                :value="item.ID"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="课程阶段:" prop="courseContent">
                    <el-input :disabled="true" v-model="formData.courseContent" placeholder="根据班级ID自动生成"/>
                </el-form-item>
                <el-form-item label="上课时间:" prop="timeString">
                    <el-input :disabled="true" v-model="formData.timeString" placeholder="根据班级ID自动生成"/>
                </el-form-item>
                <el-form-item label="占用教室:" prop="room">
                    <el-input :disabled="true" v-model="formData.room" placeholder="根据班级ID自动生成"/>
                </el-form-item>
                <el-form-item label="授课教师:" prop="teacherName">
                    <el-input v-model="formData.teacherName" clearable placeholder="根据班级ID自动生成"/>
                </el-form-item>
                <el-form-item label="课程名称:" prop="name">
                    <el-input v-model="formData.name" clearable placeholder="请填写本节课的名称"/>
                </el-form-item>
                <el-form-item label="参课学生:" prop="students">
                    <el-select
                            v-model="formData.students"
                            multiple
                            @change="autoFillAbsent"
                            placeholder="请选择参课学生">
                        <el-option
                                v-for="item in studentOptions"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="请假学生:" prop="absentStudents">
                    <el-select
                            v-model="formData.absentStudents"
                            multiple
                            :disabled="true"
                            placeholder="自动填充请假学生">
                        <el-option
                                v-for="item in studentOptions"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
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
    import {
        createCourse,
        deleteCourse,
        deleteCourseByIds,
        updateCourse,
        findCourse,
        getCourseList,
        getClassListFromCourse,
    } from '@/api/z_course' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Course',
        mixins: [infoList],
        data() {
            return {
                listApi: getCourseList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                classOptions: [],
                studentOptions: [],

                formData: {
                    classId: null,
                    name: null,
                    courseContent: null,
                    timeString: null,
                    room: null,
                    teacherName: null,
                    comment: null,
                    students: [],
                    absentStudents: [],
                },
                rules: {
                    classId: [{ required: true, message: '请选择班级', trigger: 'blur' }],
                    name: [{ required: true, message: '课程名称不能为空', trigger: 'blur' }],
                    teacherName: [{ required: true, message: '教师名称不能为空', trigger: 'blur' }],
                    students: [{ required: true, message: '参课学生不能为空', trigger: 'blur' }],
                }
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
            formatDate(row) {
                return formatTimeToStr(row.CreatedAt, 'yyyy-MM-dd');
            },
            async getClasses() {
                let res = await getClassListFromCourse()
                if (res.code === 0){
                    this.classOptions = res.data
                }
            },
            autoFill() {
                let id = this.formData.classId
                //选择指定id的班级
                let selectedClass = this.classOptions.find((item) => {return item.ID === id})

                this.formData.courseContent = selectedClass.courseContent
                this.formData.timeString = selectedClass.timeString
                this.formData.room = selectedClass.room
                this.formData.teacherName = selectedClass.teacherName
                //使用map方法获取对象列表中的某个属性集合
                this.studentOptions = selectedClass.students.map((item)=>{return item.name})
                this.formData.students = this.studentOptions
            },
            autoFillAbsent() {
                this.formData.absentStudents = []
                for (let i = 0; i < this.studentOptions.length; i++) {
                    if (this.formData.students.includes(this.studentOptions[i])) {
                        //
                    } else {
                        this.formData.absentStudents.push(this.studentOptions[i])
                    }
                }
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
                    this.deleteCourse(row)
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
                const res = await deleteCourseByIds({ ids })
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
            async updateCourse(row) {
                const res = await findCourse({ ID: row.ID })
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.recourse
                    this.studentOptions = this.formData.students
                    for (let i = 0; i < this.formData.absentStudents.length; i++ ) {
                        this.studentOptions.push(this.formData.absentStudents[i])
                    }
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    classId: null,
                    name: null,
                    courseContent: null,
                    timeString: null,
                    room: null,
                    teacherName: null,
                    comment: null,
                    students: [],
                    absentStudents: [],
                }
                this.classOptions = []
                this.studentOptions = []
            },
            async deleteCourse(row) {
                const res = await deleteCourse({ ID: row.ID })
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
            beforeEnterDialog(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.enterDialog()
                    }
                })
            },
            async enterDialog() {
                let res
                switch (this.type) {
                    case "create":
                        res = await createCourse(this.formData)
                        break
                    case "update":
                        res = await updateCourse(this.formData)
                        break
                    default:
                        res = await createCourse(this.formData)
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
            async openDialog() {
                await this.getClasses()
                this.type = 'create'
                this.dialogFormVisible = true
            }
        },
    }
</script>

<style>
</style>


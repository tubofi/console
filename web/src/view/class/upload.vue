
<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="班号">
                    <el-input v-model.number="searchInfo.classId" placeholder="请输入班级编号" />
                </el-form-item>
                <el-form-item label="课名">
                    <el-input v-model="searchInfo.name" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="教师">
                    <el-input v-model="searchInfo.teacherName" placeholder="请输入模糊关键词" />
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

            <el-table-column label="源码上传" align="center" width="300">
                <template slot-scope="scope">
                    <el-upload
                            class="upload-demo"
                            action="https://jsonplaceholder.typicode.com/posts/"
                            :on-preview="handlePreview"
                            :on-remove="handleRemove"
                            :before-remove="beforeRemove"
                            multiple
                            :limit="3"
                            :on-exceed="handleExceed"
                            :file-list="fileList">
                        <el-button size="mini" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">上传本节课sb3格式源码</div>
                    </el-upload>
                </template>
            </el-table-column>
            <el-table-column label="素材上传" align="center" width="300">
                <template>
                    <el-upload
                            class="avatar-uploader"
                            action="https://saisicode-1304003768.cos.ap-chengdu.myqcloud.com"
                            :show-file-list="false"
                            :before-upload="beforeAvatarUpload"
                            :on-progress="getTmp_secret_keys">
                        <el-button size="mini" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">上传本节课素材压缩文件</div>
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
    import {getAllTeachers} from '@/api/z_teacher'
    import upload from '@/utils/z_upload'
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
                teacherOptions: [],
                fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}, {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}],

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
            beforeAvatarUpload (file) {
                console.log(file)
                this.uploadFile = file
            },
            // 上传文件
            getTmp_secret_keys() {
                let files = upload(this.uploadFile,res => {
                    console.log(res,'ddddd')

                })
            },
            handleRemove(file, fileList) {
                console.log(file, fileList);
            },
            handlePreview(file) {
                console.log(file);
            },
            handleExceed(files, fileList) {
                this.$message.warning(`当前限制选择 3 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`);
            },
            beforeRemove(file, fileList) {
                return this.$confirm(`确定移除 ${ file.name }？`);
            },
            async allTeachers(){
                const res = await getAllTeachers()
                if (res.code === 0) {
                    this.teacherOptions = res.data
                }
            },
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
                    this.allTeachers()
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
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
                };
                this.classOptions = [];
                this.studentOptions = [];
                this.teacherOptions = [];
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
                this.allTeachers()
            }
        },
    }
</script>

<style>
</style>


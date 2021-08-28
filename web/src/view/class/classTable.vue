
<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="课程">
                    <el-input v-model="searchInfo.courseContent" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="时间">
                    <el-input v-model="searchInfo.timeString" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item label="教室">
                    <el-select v-model.number="searchInfo.room" placeholder="请选择教室">
                        <el-option
                                v-for="item in roomOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
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
            <el-table-column label="班级ID" prop="ID" align="center" width="80"/>
            <el-table-column label="日期" prop="CreatedAt" align="center" width="130" :formatter="formatDate"/>
            <el-table-column label="课程" prop="courseContent" align="center"/>
            <el-table-column label="时间" prop="timeString" align="center" width="120"/>
            <el-table-column label="教室" prop="room" align="center"/>
            <el-table-column label="教师" prop="teacherName" align="center"/>
            <el-table-column label="学员" prop="students" align="left" header-align="center" width="400">
                <template slot-scope="scope">
                    <span v-for="item in scope.row.students">
                        <el-tag size="medium">{{item.name}}</el-tag>
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="详情" align="center" width="80">
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
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateClass(scope.row)">编辑</el-button>
                    <el-button plain type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新班级表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="课程类型:" prop="courseType">
                    <el-select v-model="formData.courseType" clearable placeholder="选择课程类型">
                        <el-option
                                v-for="item in courseOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="课程阶段:" prop="stage">
                    <el-select v-model.number="formData.stage" clearable placeholder="选择课程阶段">
                        <el-option
                                v-for="item in stageOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="第几期:" prop="phase">
                    <el-select v-model.number="formData.phase" clearable placeholder="选择第几期">
                        <el-option
                                v-for="item in phaseOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="周几上课:" prop="weekday">
                    <el-select v-model.number="formData.weekday" clearable placeholder="选择周几上课">
                        <el-option
                                v-for="item in weekOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="上课时间:" prop="minute">
                    <el-select v-model.number="formData.hour" clearable placeholder="选择时">
                        <el-option
                                v-for="item in hourOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                    <el-select v-model.number="formData.minute" clearable placeholder="选择分">
                        <el-option
                                v-for="item in minuteOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="选择教室:" prop="room">
                    <el-select v-model.number="formData.room" clearable placeholder="选择教室">
                        <el-option
                                v-for="item in roomOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="添加学生:" prop="students">
                    <el-select
                            v-model="formData.students"
                            value-key="ID"
                            multiple
                            filterable
                            placeholder="输入关键字搜索">
                        <el-option
                                v-for="item in studentOptions"
                                :key="item.ID"
                                :label="item.name"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="选择教师:" prop="teacherName">
                    <el-select v-model="formData.teacherName" placeholder="请选择教师">
                        <el-option
                                v-for="item in teacherOptions"
                                :key="item.ID"
                                :label="item.name"
                                :value="item.name"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
                    <el-input type="textarea" :rows="3" v-model.number="formData.comment" clearable placeholder="备注信息"/>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="mini" @click="closeDialog">取 消</el-button>
                <el-button type="primary" size="mini" @click="beforeEnterDialog('formData')">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    const courseOptions = [
        {label: '无', value: 'none'},
        {label: 'osmo', value: 'osmo'},
        {label: 'scratchJR', value: 'scratchJR'},
        {label: 'scratch', value: 'scratch'},
        {label: 'python', value: 'python'},
        {label: 'cpp', value: 'cpp'},
    ];
    const stageOptions = [
        {label: '第一阶段', value: 1},
        {label: '第二阶段', value: 2},
        {label: '第三阶段', value: 3},
        {label: '第四阶段', value: 4},
    ];
    const phaseOptions = [
        {label: '第一期', value: 1},
        {label: '第二期', value: 2},
        {label: '第三期', value: 3},
        {label: '第四期', value: 4},
    ];
    const weekOptions = [
        {label: '星期一', value: 1},
        {label: '星期二', value: 2},
        {label: '星期三', value: 3},
        {label: '星期四', value: 4},
        {label: '星期五', value: 5},
        {label: '星期六', value: 6},
        {label: '星期日', value: 0},
    ];
    const hourOptions = [
        {label: '08:', value: 8},
        {label: '10:', value: 10},
        {label: '13:', value: 13},
        {label: '15:', value: 15},
        {label: '17:', value: 17},
        {label: '18:', value: 18},
    ];
    const minuteOptions = [
        {label: '00', value: 0},
        {label: '10', value: 10},
        {label: '20', value: 20},
        {label: '30', value: 30},
        {label: '40', value: 40},
        {label: '50', value: 50},
    ];
    const roomOptions = [
        {label: '1号教室', value: 1},
        {label: '2号教室', value: 2},
        {label: '3号教室', value: 3},
        {label: '4号教室', value: 4},
    ];
    import {
        createClass,
        deleteClass,
        updateClass,
        findClass,
        getClassList,
        getIdleStudents
    } from '@/api/z_class' //  此处请自行替换地址
    import {getAllTeachers} from '@/api/z_teacher'
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Class',
        mixins: [infoList],
        data() {
            return {
                listApi: getClassList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                courseOptions: courseOptions,
                stageOptions: stageOptions,
                phaseOptions: phaseOptions,
                weekOptions: weekOptions,
                hourOptions: hourOptions,
                minuteOptions: minuteOptions,
                roomOptions: roomOptions,
                studentOptions: [],
                teacherOptions: [],
                students: [],

                formData: {
                    ID: null,
                    teacherName: null,
                    room: null,
                    courseType: null,
                    stage: null,
                    phase: null,
                    courseContent: null,
                    weekday: null,
                    hour: null,
                    minute: null,
                    comment: null,
                    timeString: null,
                    students: [],
                },
                rules: {
                    courseType: [{ required: true, message: '课程类型不能为空', trigger: 'blur' }],
                    teacherName: [{ required: true, message: '教师不能为空', trigger: 'blur' }],
                    room: [{ required: true, message: '教室不能为空', trigger: 'blur' }],
                    stage: [{ required: true, message: '请选择课程阶段', trigger: 'blur' }],
                    phase: [{ required: true, message: '请选择第几期', trigger: 'blur' }],
                    weekday: [{ required: true, message: '请选择星期几', trigger: 'blur' }],
                    hour: [{ required: true, message: '请选择上课时间', trigger: 'blur' }],
                    minute: [{ required: true, message: '请选择上课时间', trigger: 'blur' }],
                    students: [{ required: true, message: '请选择学生', trigger: 'blur' }],
                },
            }
        },
        filters: {
            formatDate: function(time) {
                if (time !== null && time !== '') {
                    let date = new Date(time);
                    return formatTimeToStr(date, 'yyyy-MM-dd');
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
            async idleStudents(){
                const res = await getIdleStudents()
                if (res.code === 0) {
                    this.studentOptions = res.data
                }
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
                    this.deleteClass(row)
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
                const res = await deleteClassByIds({ ids })
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
            async updateClass(row) {
                await this.idleStudents();
                const res = await findClass({ ID: row.ID });
                this.type = 'update';
                if (res.code === 0) {
                    this.formData = res.data.reclass;
                    for (let i = 0; i < this.formData.students.length; i++) {
                        this.studentOptions.push(this.formData.students[i])
                    }
                    this.dialogFormVisible = true;
                    this.allTeachers();
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {
                    ID: null,
                    teacherName: null,
                    room: null,
                    courseType: null,
                    stage: null,
                    phase: null,
                    courseContent: null,
                    weekday: null,
                    hour: null,
                    minute: null,
                    comment: null,
                    timeString: null,
                    students: [],
                };
                this.studentOptions = [];
                this.teacherOptions = [];
            },
            async deleteClass(row) {
                const res = await deleteClass({ ID: row.ID })
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功'
                    });
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
                let res;
                switch (this.type) {
                    case "create":
                        res = await createClass(this.formData)
                        break
                    case "update":
                        res = await updateClass(this.formData)
                        break
                    default:
                        res = await createClass(this.formData)
                        break
                }
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '创建/更改成功'
                    });
                    this.closeDialog();
                    this.getTableData()
                }
            },
            openDialog() {
                this.type = 'create';
                this.dialogFormVisible = true
                this.allTeachers();
                this.idleStudents();
            },
            fmtBody(value) {
                try {
                    return JSON.parse(value)
                } catch (err) {
                    return value
                }
            }
        },
    }
</script>

<style>
</style>


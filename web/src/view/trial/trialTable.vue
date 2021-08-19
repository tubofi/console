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
                <el-form-item label="教师姓名">
                    <el-input v-model="searchInfo.teacherName" placeholder="请输入模糊关键词" />
                </el-form-item>
                <el-form-item>
                    <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
                    <el-popover v-model="deleteVisible" placement="top" width="160">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
                        </div>
                        <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
                    </el-popover>
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
            <el-table-column type="selection" width="60px" />
            <el-table-column label="日期" prop="time" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="学生姓名" prop="studentName" align="center"/>
            <el-table-column label="课程类型" prop="courseType" align="center"/>
            <el-table-column label="课程名称" prop="courseName" align="center"/>
            <el-table-column label="授课教师" prop="teacherName" align="center"/>
            <el-table-column label="教室" prop="name" align="center"/>
            <el-table-column label="详情" align="center" width="100">
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
                    <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateTrialCourseRecord(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新试听表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="选择学生:" prop="studentName">
                    <el-select
                            v-model="formData.studentName"
                            value-key="ID"
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
                <el-form-item label="课程类型:" prop="courseType">
                    <el-select v-model="formData.courseType" clearable placeholder="选择课程类型">
                        <el-option
                                v-for="item in courseOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="课程名称:">
                    <el-input v-model="formData.courseName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="上课时间:">
                    <el-date-picker
                            v-model="formData.attendTime"
                            type="datetime"
                            placeholder="选择日期时间">
                    </el-date-picker>
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
                <el-form-item label="教师:">
                    <el-input v-model="formData.teacherName" clearable placeholder="请输入" />
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
                    <el-input type="textarea" :rows="3" v-model.number="formData.comment" clearable placeholder="备注信息"/>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button type="primary" @click="enterDialog">确 定</el-button>
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
    const roomOptions = [
        {label: '1号教室', value: 1},
        {label: '2号教室', value: 2},
        {label: '3号教室', value: 3},
        {label: '4号教室', value: 4},
    ];
    import {
        createTrialCourseRecord,
        deleteTrialCourseRecord,
        deleteTrialCourseRecordByIds,
        updateTrialCourseRecord,
        findTrialCourseRecord,
        getTrialCourseRecordList,
        getAllTrialStudents,
    } from '@/api/z_trial_course_record' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'TrialCourseRecord',
        mixins: [infoList],
        data() {
            return {
                listApi: getTrialCourseRecordList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                courseOptions: courseOptions,
                roomOptions: roomOptions,
                studentOptions: [],

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
                }
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
            async trialStudents() {
                const res = await getAllTrialStudents()
                if (res.code === 0) {
                    this.studentOptions = res.data
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
                this.dialogFormVisible = false;
                this.formData = {
                    studentName: '',
                    teacherName: '',
                    courseType: '',
                    courseName: '',
                    time: new Date(),
                    room: 0,
                    total: 0,
                    punctuality: 0,
                    discipline: 0,
                    concentration: 0,
                    innovation: 0,
                    logic: 0,
                    mathematics: 0,
                    complete: 0,
                    review: '',
                    comment: '',
                };
                this.studentOptions = [];
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
                        res = await updateTrialCourseRecord(this.formData)
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
                this.trialStudents()
                this.type = 'create'
                this.dialogFormVisible = true
            }
        },
    }
</script>

<style>
</style>

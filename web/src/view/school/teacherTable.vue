
<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="姓名">
                    <el-input v-model="searchInfo.name" placeholder="请输入教师姓名" />
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
            <el-table-column label="日期" prop="entryTime" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="姓名" prop="name" align="center"/>
            <el-table-column label="性别" prop="sex" align="center" />
            <el-table-column label="手机号" prop="phone" align="center" width="180"/>
            <el-table-column label="地址" prop="address" align="center" />
            <el-table-column label="身份证号" prop="identityNumber" align="center"/>
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
                    <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateTeacher(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新教师表单" width="50%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right">
                <el-form-item>
                    <span style="font-size: 18px">基本信息</span>
                    <!---第一行--->
                    <el-row :gutter="10" type="flex" justify="space-between">
                        <el-col :span="8">
                            <el-form-item label="姓名:" prop="name" label-width="100px">
                                <el-input v-model="formData.name" clearable placeholder="请输入教师姓名" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="性别:" prop="sex" label-width="100px">
                                <el-select v-model="formData.sex" placeholder="请选择性别">
                                    <el-option
                                            v-for="item in sexOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="电话:" prop="phone" label-width="100px">
                                <el-input v-model="formData.phone" clearable placeholder="请输入电话号码" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <!---第二行--->
                    <el-row :gutter="10" type="flex" justify="space-between">
                        <el-col :span="8">
                            <el-form-item label="现居地址:" prop="address" label-width="100px">
                                <el-input v-model="formData.address" clearable placeholder="请输入地址" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="身份证号:" prop="identityNumber" label-width="100px">
                                <el-input v-model="formData.identityNumber" clearable placeholder="请输入身份证号" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="加入时间:" prop="time" label-width="100px">
                                <el-date-picker v-model="formData.time" type="date" placeholder="请选择日期"/>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </el-form-item>

                <el-form-item>
                    <!---第一行--->
                    <span style="font-size: 18px">基础工资</span>
                    <el-row :gutter="10" type="flex" justify="space-between">
                        <el-col :span="8">
                            <el-form-item label="基本工资:" prop="baseWage" label-width="100px">
                                <el-input v-model.number="formData.baseWage" clearable placeholder="请输入基本工资" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="岗位工资:" prop="postWage" label-width="100px">
                                <el-input v-model.number="formData.postWage" clearable placeholder="请输入岗位工资" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="岗位补贴:" prop="allowance" label-width="100px">
                                <el-input v-model.number="formData.allowance" clearable placeholder="请输入岗位补贴" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <!---第二行--->
                    <span style="font-size: 18px">教师工资</span>
                    <el-row :gutter="10" type="flex" justify="space-between">
                        <el-col :span="8">
                            <el-form-item label="课时费:" prop="coursePer" label-width="100px">
                                <el-input v-model.number="formData.coursePer" clearable placeholder="？元/次/人" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="转化提成:" prop="turnPercent" label-width="100px">
                                <el-select v-model.number="formData.turnPercent" placeholder="请选择提成比例">
                                    <el-option
                                            v-for="item in percentOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="续费提成:" prop="renewPercent" label-width="100px">
                                <el-select v-model.number="formData.renewPercent" placeholder="请选择提成比例">
                                    <el-option
                                            v-for="item in percentOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <!---第三行--->
                    <span style="font-size: 18px">招生工资</span>
                    <el-row :gutter="10" type="flex" justify="space-between">
                        <el-col :span="8">
                            <el-form-item label="新单提成:" prop="newPercent" label-width="100px">
                                <el-select v-model.number="formData.newPercent" placeholder="请选择提成比例">
                                    <el-option
                                            v-for="item in percentOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="续费提成:" prop="oldPercent" label-width="100px">
                                <el-select v-model.number="formData.oldPercent" placeholder="请选择提成比例">
                                    <el-option
                                            v-for="item in percentOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="团队提成:" prop="newTeamPercent" label-width="100px">
                                <el-select v-model.number="formData.newTeamPercent" placeholder="请选择提成比例">
                                    <el-option
                                            v-for="item in percentOptions"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"/>
                                </el-select>
                            </el-form-item>
                        </el-col>
                    </el-row>
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
    const sexOptions = [
        {label: '男', value: '男'},
        {label: '女', value: '女'},
    ];
    const percentOptions = [
        {label: '0%', value: 0.0},
        {label: '2%', value: 0.02},
        {label: '3%', value: 0.03},
        {label: '5%', value: 0.05},
        {label: '10%', value: 0.1},
        {label: '20%', value: 0.2},
    ];
    import {
        createTeacher,
        deleteTeacher,
        deleteTeacherByIds,
        updateTeacher,
        findTeacher,
        getTeacherList
    } from '@/api/z_teacher' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Teacher',
        mixins: [infoList],
        data() {
            return {
                listApi: getTeacherList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                sexOptions: sexOptions,
                percentOptions: percentOptions,

                formData: {
                    name: null,
                    sex: null,
                    address: null,
                    identityNumber: null,
                    phone: null,
                    entryTime: new Date(),
                    baseWage: null,
                    postWage: null,
                    allowance: null,
                    coursePer: null,
                    turnPercent: null,
                    renewPercent: null,
                    newPercent: null,
                    oldPercent: null,
                    newTeamPercent: null,
                    oldTeamPercent: null,
                    comment: null,
                },
                rules: {
                    name: [{ required: true, message: '姓名不能为空', trigger: 'blur' }],
                    sex: [{ required: true, message: '性别不能为空', trigger: 'blur' }],
                    phone: [{ required: true, message: '联系方式不能为空', trigger: 'blur' }],
                    baseWage: [{ required: true, message: '基本工资不能为空', trigger: 'blur' }],
                    postWage: [{ required: true, message: '岗位工资不能为空', trigger: 'blur' }],
                    allowance: [{ required: true, message: '岗位补贴不能为空', trigger: 'blur' }],
                    coursePer: [{ required: true, message: '课时费不能为空', trigger: 'blur' }],
                    turnPercent: [{ required: true, message: '请选择提成比例', trigger: 'blur' }],
                    renewPercent: [{ required: true, message: '请选择提成比例', trigger: 'blur' }],
                    newPercent: [{ required: true, message: '请选择提成比例', trigger: 'blur' }],
                    oldPercent: [{ required: true, message: '请选择提成比例', trigger: 'blur' }],
                    newTeamPercent: [{ required: true, message: '请选择提成比例', trigger: 'blur' }],
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
            fmtBody(value) {
                try {
                    return JSON.parse(value)
                } catch (err) {
                    return value
                }
            },
            beforeEnterDialog(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.enterDialog()
                    }
                })
            },
            formatDate(row) {
                return formatTimeToStr(row.entryTime, 'yyyy-MM-dd');
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
                    this.deleteTeacher(row)
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
                const res = await deleteTeacherByIds({ ids })
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
            async updateTeacher(row) {
                const res = await findTeacher({ ID: row.ID })
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.reteacher
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    schoolId: 0,
                    name: '',
                    sex: '',
                    address: '',
                    identityNumber: '',
                    phone: '',
                    entryTime: new Date(),
                    baseWage: 0,
                    postWage: 0,
                    allowance: 0,
                    coursePer: 0,
                    turnPercent: 0,
                    renewPercent: 0,
                    newPercent: 0,
                    oldPercent: 0,
                    newTeamPercent: 0,
                    oldTeamPercent: 0,
                    comment: '',

                }
            },
            async deleteTeacher(row) {
                const res = await deleteTeacher({ ID: row.ID })
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
                        res = await createTeacher(this.formData)
                        break
                    case "update":
                        res = await updateTeacher(this.formData)
                        break
                    default:
                        res = await createTeacher(this.formData)
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


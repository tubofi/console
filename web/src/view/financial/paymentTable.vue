<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="姓名">
                    <el-input v-model.nuber="searchInfo.studentName" placeholder="请输入关键字"/>
                </el-form-item>
                <el-form-item label="日期">
                    <el-date-picker
                            v-model="searchInfo.CreatedAt"
                            type="month"
                            placeholder="选择日期">
                    </el-date-picker>
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
            <el-table-column label="缴费日期" prop="CreatedAt" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="学生姓名" prop="studentName" align="center"/>
            <el-table-column label="缴费金额" prop="money" align="center"/>
            <el-table-column label="缴费类型" align="center">
                <template slot-scope="scope">
                    <el-tag size="medium" v-if="scope.row.pattern === 1" type="success">报名</el-tag>
                    <el-tag v-else-if="scope.row.pattern === 2" size="medium">续费</el-tag>
                    <el-tag v-else-if="scope.row.pattern === 3" size="medium" type="warning">体验</el-tag>
                    <el-tag v-else size="medium" type="danger">异常</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="课程次数" prop="times" align="center"/>
            <el-table-column label="说明" prop="comment" align="center"/>
             <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updatePayment(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="缴费表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="学生姓名:" prop="studentName">
                    <el-select
                            v-model="formData.studentName"
                            value-key="ID"
                            filterable
                            placeholder="输入关键字搜索">
                        <el-option
                                v-for="item in studentOptions"
                                :key="item.ID"
                                :label="item.name"
                                :value="item.name">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="缴费金额:" prop="money">
                    <el-input v-model.number="formData.money" clearable placeholder="请输入缴费金额" />
                </el-form-item>
                <el-form-item label="缴费类型:" prop="pattern">
                    <el-select v-model="formData.pattern" clearable placeholder="选择缴费类型">
                        <el-option
                                v-for="item in patternOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="课程次数:" prop="times">
                    <el-input v-model.number="formData.times" clearable placeholder="请输入课程次数" />
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
    const patternOptions = [
        {label: '报名', value: 1},
        {label: '续费', value: 2},
        {label: '体验', value: 3},
    ];
    import {
        createPayment,
        deletePayment,
        deletePaymentByIds,
        updatePayment,
        findPayment,
        getPaymentList,
        getAllStudents,
    } from '@/api/z_payment' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Payment',
        mixins: [infoList],
        data() {
            return {
                listApi: getPaymentList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                patternOptions: patternOptions,
                studentOptions: [],

                formData: {
                    CreatedAt: null,
                    studentName: null,
                    money: null,
                    pattern: null,
                    times: null,
                    comment: null,
                },
                rules: {
                    studentName: [{ required: true, message: '学生姓名不能为空', trigger: 'blur' }],
                    money: [{ required: true, message: '请输入缴费金额', trigger: 'blur' }],
                    pattern: [{ required: true, message: '请选择缴费类型', trigger: 'blur' }],
                    times: [{ required: true, message: '请输入课程次数', trigger: 'blur' }],
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
            async allStudents() {
                const res = await getAllStudents()
                if (res.code === 0) {
                    this.studentOptions = res.data
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
                    this.deletePayment(row)
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
                const res = await deletePaymentByIds({ ids })
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
            async updatePayment(row) {
                const res = await findPayment({ ID: row.ID })
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.repayment
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    CreatedAt: null,
                    studentName: null,
                    money: null,
                    pattern: null,
                    times: null,
                    comment: null,
                };
                this.studentOptions = [];
            },
            async deletePayment(row) {
                const res = await deletePayment({ ID: row.ID })
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
                        res = await createPayment(this.formData)
                        break
                    case "update":
                        res = await updatePayment(this.formData)
                        break
                    default:
                        res = await createPayment(this.formData)
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
                this.allStudents()
            }
        },
    }
</script>

<style>
</style>

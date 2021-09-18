<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="年">
                    <el-input v-model.nuber="searchInfo.year" placeholder="请输入整数"/>
                </el-form-item>
                <el-form-item label="月">
                    <el-input v-model.number="searchInfo.month" placeholder="请输入整数"/>
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
            <el-table-column label="日期" align="center" width="150">
                <template slot-scope="scope">
                    <pre>{{scope.row.year}}年{{scope.row.month}}月</pre>
                </template>
            </el-table-column>
            <el-table-column label="收入情况" align="center">
                <el-table-column label="缴费收入" prop="paymentIncome" align="center"></el-table-column>
                <el-table-column label="其它收入" prop="otherIncome" align="center"></el-table-column>
                <el-table-column label="总收入" prop="allIncome" align="center"></el-table-column>
            </el-table-column>
            <el-table-column label="支出情况" align="center">
                <el-table-column label="工资支出" prop="wageOutcome" align="center"></el-table-column>
                <el-table-column label="其它支出" prop="otherOutcome" align="center"></el-table-column>
                <el-table-column label="总支出" prop="allOutcome" align="center"></el-table-column>
            </el-table-column>
            <el-table-column label="利润" prop="profit" align="center"/>
            <el-table-column label="是否结算" align="center">
                <template slot-scope="scope">
                    <el-tag size="medium" v-if="scope.row.isFinished === 1" type="success">已结算</el-tag>
                    <el-tag size="medium" v-if="scope.row.isFinished === 0" type="warning">未结算</el-tag>
                    <el-tag size="medium" v-else type="danger">异常</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="详情" align="center" width="100">
                <template slot-scope="scope">
                    <div>
                        <el-popover v-if="scope.row" placement="top-start" trigger="hover">
                            <div class="popover-box">
                                <pre>{{ fmtBody(scope.row) }}</pre>
                            </div>
                            <i slot="reference" class="el-icon-view"/>
                        </el-popover>
                        <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button size="mini" type="danger" icon="el-icon-edit" class="table-button" @click="updateBill(scope.row)">更新</el-button>
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
        createBill,
        deleteBill,
        deleteBillByIds,
        updateBill,
        findBill,
        getBillList
    } from '@/api/z_bill' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Bill',
        mixins: [infoList],
        data() {
            return {
                listApi: getBillList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],

                formData: {
                    year: null,
                    month: null,
                    isFinished: null,
                    paymentIncome: null,
                    otherIncome: null,
                    allIncome: null,
                    wageOutcome: null,
                    otherOutcome: null,
                    allOutcome: null,
                    profit: null,
                    comment: null,
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
            formatDate(row) {
                return formatTimeToStr(row.entryTime, 'yyyy-MM');
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
                    this.deleteBill(row)
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
                const res = await deleteBillByIds({ ids })
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
            async updateBill(row) {
                const res = await updateBill({ ID: row.ID })
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '更新成功'
                    })
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    year: null,
                    month: null,
                    isFinished: null,
                    paymentIncome: null,
                    otherIncome: null,
                    allIncome: null,
                    wageOutcome: null,
                    otherOutcome: null,
                    allOutcome: null,
                    profit: null,
                    comment: null,
                }
            },
            async deleteBill(row) {
                const res = await deleteBill({ ID: row.ID })
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
                        res = await createBill(this.formData);
                        break
                    case "update":
                        res = await updateBill(this.formData);
                        break
                    default:
                        res = await createBill(this.formData);
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

<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="年份">
                    <el-select v-model.number="searchInfo.year" placeholder="请选择年份">
                        <el-option
                                v-for="item in yearOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="月份">
                    <el-select v-model.number="searchInfo.month" placeholder="请选择月份">
                        <el-option
                                v-for="item in monthOptions"
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
            <el-table-column label="创建日期" prop="CreatedAt" align="center" width="180" :formatter="formatDate"/>
            <el-table-column label="年份" prop="year" align="center"/>
            <el-table-column label="月份" prop="month" align="center"/>
            <el-table-column label="总额" prop="all" align="center"/>
            <el-table-column label="是否结算" align="center">
                <template slot-scope="scope">
                    <el-tag size="medium" v-if="scope.row.isClose === 1" type="success">已结算</el-tag>
                    <el-tag v-else size="medium" type="danger">未结算</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="消费详情" align="center" width="100">
                <template slot-scope="scope">
                    <div>
                        <el-popover v-if="scope.row.livingCosts" placement="top-start" trigger="hover">
                            <div class="popover-box">
                                <pre>{{ fmtBody(scope.row.livingCosts) }}</pre>
                            </div>
                            <i slot="reference" class="el-icon-view" />
                        </el-popover>
                        <span v-else>无</span>
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
    </div>
</template>

<script>
    const yearOptions = [
        {label: '2021年', value: 2021},
        {label: '2022年', value: 2022},
        {label: '2023年', value: 2023},
        {label: '2024年', value: 2024},
        {label: '2025年', value: 2025},
        {label: '2026年', value: 2026},
        {label: '2027年', value: 2027},
        {label: '2028年', value: 2028},
        {label: '2029年', value: 2029},
    ];
    const monthOptions = [
        {label: '一月', value: 1},
        {label: '二月', value: 2},
        {label: '三月', value: 3},
        {label: '四月', value: 4},
        {label: '五月', value: 5},
        {label: '六月', value: 6},
        {label: '七月', value: 7},
        {label: '八月', value: 8},
        {label: '九月', value: 9},
        {label: '十月', value: 10},
        {label: '十一月', value: 11},
        {label: '十二月', value: 12},
    ];
    import {
        getLivingBillList,
    } from '@/api/z_living' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'LivingBill',
        mixins: [infoList],
        data() {
            return {
                listApi: getLivingBillList,
                yearOptions: yearOptions,
                monthOptions: monthOptions,

                formData: {
                    CreatedAt: null,
                    year: null,
                    month: null,
                    isClose: null,
                    all: null,
                    livingCosts: [],
                    comment: null,
                },
            }
        },
        async created() {
            await this.getTableData();
        },
        methods: {
            formatDate(row) {
                return formatTimeToStr(row.CreatedAt, 'yyyy-MM-dd');
            },
            // 条件搜索前端看此方法
            onSubmit() {
                this.page = 1;
                this.pageSize = 10;
                this.getTableData();
            },
            handleSelectionChange(val) {
                this.multipleSelection = val
            },
        },
    }
</script>

<style>
</style>


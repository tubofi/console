
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
                <el-form-item label="类别">
                    <el-select v-model.number="searchInfo.type" placeholder="请选择类别">
                        <el-option
                                v-for="item in typeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="统计">
                    <el-input v-model.number="total" :disabled="true" placeholder="0"/>
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
            <el-table-column label="日期" prop="time" align="center" width="180" :formatter="formatDate"/>
            <el-table-column label="消费金额" prop="money" align="center"/>
            <el-table-column label="消费类型" prop="type" align="center" :formatter="typeLabel"/>
            <el-table-column label="消费说明" prop="comment" align="center"/>
            <el-table-column label="按钮组" align="center" width="200">
                <template slot-scope="scope">
                    <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateCost(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新消费表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="消费日期:" prop="time">
                    <el-date-picker
                            v-model="formData.time"
                            type="date"
                            placeholder="选择日期">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="消费金额:" prop="money">
                    <el-input v-model.number="formData.money" clearable placeholder="请输入消费金额" />
                </el-form-item>
                <el-form-item label="消费类型:" prop="type">
                    <el-select v-model.number="formData.type" clearable placeholder="请选择消费类型">
                        <el-option
                                v-for="item in typeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="消费说明:" prop="comment">
                    <el-input type="textarea" :rows="3" v-model="formData.comment" clearable placeholder="备注信息"/>
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
    //1衣裤 2饮食 3居住 4交通 5生活用品 6家用电器 7电子设备 8化妆品 9药物 10宠物 11游玩 12书本文具 13其它
    const typeOptions = [
        {label: '衣裤', value: 1},
        {label: '饮食', value: 2},
        {label: '居住', value: 3},
        {label: '交通', value: 4},
        {label: '生活用品', value: 5},
        {label: '家用电器', value: 6},
        {label: '电子设备', value: 7},
        {label: '化妆品', value: 8},
        {label: '药物', value: 9},
        {label: '宠物', value: 10},
        {label: '游玩', value: 11},
        {label: '书本文具', value: 12},
        {label: '其它', value: 13},
    ];
    import {
        createLivingCost,
        deleteLivingCost,
        updateLivingCost,
        findLivingCost,
        getLivingCostList,
    } from '@/api/z_living' //  此处请自行替换地址
    import {getAllTeachers} from '@/api/z_teacher'
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Class',
        mixins: [infoList],
        data() {
            return {
                listApi: getLivingCostList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                yearOptions: yearOptions,
                monthOptions: monthOptions,
                typeOptions: typeOptions,
                all: 0,

                formData: {
                    year: null,
                    month: null,
                    time: null,
                    money: null,
                    type: null,
                    comment: null,
                },
                rules: {
                    time: [{ required: true, message: '请选择消费日期', trigger: 'blur' }],
                    money: [{ required: true, message: '请输入消费金额', trigger: 'blur' }],
                    type: [{ required: true, message: '请选择消费类型', trigger: 'blur' }],
                },
            }
        },
        async created() {
            this.searchInfo.year = new Date().getFullYear();
            this.searchInfo.month = new Date().getMonth();
            await this.getTableData();
            this.all = 0;
            for (let i = 0; i < this.tableData.length; i++) {
                this.all += this.this.tableData[i].money
            }
        },
        methods: {
            typeLabel(row){
                for (let i = 0; i < typeOptions.length; i++) {
                    if (typeOptions[i].value === row.type){
                        return typeOptions[i].label
                    }
                }
            },
            formatDate(row) {
                return formatTimeToStr(row.CreatedAt, 'yyyy-MM-dd');
            },
            // 条件搜索前端看此方法
            onSubmit() {
                this.page = 1;
                this.pageSize = 10;
                this.getTableData();
                this.all = 0;
                for (let i = 0; i < this.tableData.length; i++) {
                    this.all += this.this.tableData[i].money
                }
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
                    this.deleteCost(row)
                })
            },
            async updateCost(row) {
                const res = await findLivingCost({ ID: row.ID });
                this.type = 'update';
                if (res.code === 0) {
                    this.formData = res.data.relivingCost;
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {
                    year: null,
                    month: null,
                    time: null,
                    money: null,
                    type: null,
                    comment: null,
                };
            },
            async deleteCost(row) {
                const res = await deleteLivingCost({ ID: row.ID });
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
                        res = await createLivingCost(this.formData);
                        break;
                    case "update":
                        res = await updateLivingCost(this.formData);
                        break;
                    default:
                        res = await createLivingCost(this.formData);
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
            },
        },
    }
</script>

<style>
</style>


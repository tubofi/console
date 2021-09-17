<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="姓名">
                    <el-input v-model="searchInfo.name" placeholder="请输入学生姓名" />
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
            <el-table-column label="次数" prop="amount" align="center"/>
            <el-table-column label="类型" prop="type" align="center"/>
            <el-table-column label="说明" prop="comment" align="center"/>
            <el-table-column label="comment字段" prop="comment" width="150" />
            <el-table-column label="按钮组" align="center" width="180">
                <template slot-scope="scope">
                    <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateCourseChange(scope.row)">编辑</el-button>
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
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="课次表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="添加学生:" prop="studentName">
                    <el-select
                            v-model="formData.studentName"
                            value-key="ID"
                            multiple
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
                <el-form-item label="选择类型:" prop="type">
                    <el-select v-model="formData.type" clearable placeholder="选择类型">
                        <el-option
                                v-for="item in typeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="输入课次:" prop="amount">
                    <el-input v-model.number="formData.amount" clearable placeholder="请输入次数"/>
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
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
    const typeOptions = [
        {label: '报名', value: 1},
        {label: '续费', value: 2},
        {label: '推荐', value: 3},
        {label: '赠送', value: 4},
        {label: '积分', value: 5},
        {label: '修正', value: 11},
    ];
    import {
        createCourseChange,
        deleteCourseChange,
        deleteCourseChangeByIds,
        updateCourseChange,
        findCourseChange,
        getCourseChangeList
    } from '@/api/z_course_change' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'CourseChange',
        mixins: [infoList],
        data() {
            return {
                listApi: getCourseChangeList,
                dialogFormVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                studentOptions: [],
                typeOptions: typeOptions,

                formData: {
                    studentName: null,
                    amount: null,
                    type: null,
                    comment: null,

                },

                rules: {
                    studentName: [{ required: true, message: '请选择学生', trigger: 'blur' }],
                    amount: [{ required: true, message: '课程次数不能为空', trigger: 'blur' }],
                    type: [{ required: true, message: '请选择课次类型', trigger: 'blur' }],
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
                this.page = 1;
                this.pageSize = 10;
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
                    this.deleteCourseChange(row)
                })
            },
            async onDelete() {
                const ids = [];
                if (this.multipleSelection.length === 0) {
                    this.$message({
                        type: 'warning',
                        message: '请选择要删除的数据'
                    });
                    return
                }
                this.multipleSelection &&
                this.multipleSelection.map(item => {
                    ids.push(item.ID)
                });
                const res = await deleteCourseChangeByIds({ ids });
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功'
                    });
                    if (this.tableData.length === ids.length && this.page > 1) {
                        this.page--
                    }
                    this.deleteVisible = false;
                    this.getTableData()
                }
            },
            async updateCourseChange(row) {
                const res = await findCourseChange({ ID: row.ID });
                this.type = 'update';
                if (res.code === 0) {
                    this.formData = res.data.recourseChange;
                    this.dialogFormVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {
                    studentName: '',
                    amount: 0,
                    type: 0,
                    comment: '',

                }
            },
            async deleteCourseChange(row) {
                const res = await deleteCourseChange({ ID: row.ID })
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
            async enterDialog() {
                let res;
                switch (this.type) {
                    case "create":
                        res = await createCourseChange(this.formData)
                        break;
                    case "update":
                        res = await updateCourseChange(this.formData)
                        break;
                    default:
                        res = await createCourseChange(this.formData)
                        break
                }
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '创建/更改成功'
                    });
                    this.closeDialog();
                    this.getTableData();
                }
            },
            openDialog() {
                this.type = 'create';
                this.dialogFormVisible = true
            }
        },
    }
</script>

<style>
</style>


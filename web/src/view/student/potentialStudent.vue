<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="姓名">
                    <el-input v-model="searchInfo.name" placeholder="请输入学生姓名" />
                </el-form-item>
                <el-form-item label="电话">
                    <el-input v-model="searchInfo.phone" placeholder="请输入联系方式" />
                </el-form-item>
                <el-form-item label="年龄">
                    <el-select v-model.number="searchInfo.age" clearable placeholder="请选择年龄范围">
                        <el-option
                                v-for="item in ageOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('addApi')">新增</el-button>
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
                @selection-change="handleSelectionChange"
        >
            <el-table-column label="日期" prop="CreatedAt" align="center" width="150" :formatter="formatDate"/>
            <el-table-column label="姓名" prop="name" align="center"/>
            <el-table-column label="性别" prop="sex" align="center" />
            <el-table-column label="年龄" prop="age" align="center" />
            <el-table-column label="负责人" prop="managerName" align="center" />
            <el-table-column label="状态" prop="level" align="center">
                <template slot-scope="scope">
                    <el-tag v-if="scope.row.level >= 2" size="medium" type="danger">积极</el-tag>
                    <el-tag v-else-if="scope.row.level === 1" size="medium" type="warning">一般</el-tag>
                    <el-tag v-else size="medium" type="info">忽略</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="监护人" prop="guardian" align="center"/>
            <el-table-column label="联系方式" prop="phone" align="center" width="180"/>
            <el-table-column label="跟单详情" align="center" width="80">
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
            <el-table-column label="按钮组" align="center" width="320">
                <template slot-scope="scope">
                    <el-button size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updatePotential(scope.row.ID)">编辑</el-button>
                    <el-button size="mini" type="success" icon="el-icon-top"  @click="increaseLevel(scope.row.ID)">升序</el-button>
                    <el-button size="mini" type="warning" icon="el-icon-bottom"  @click="reduceLevel(scope.row.ID)">降序</el-button>
                    <el-button size="mini" type="danger" icon="el-icon-finished"  @click="turnTrial(scope.row.ID)">试听</el-button>
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
                @size-change="handleSizeChange"
        />
        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="意向学员表单" width="30%">
            <el-form ref="formData" :model="formData" :rules="rules" label-position="right" label-width="100px">
                <el-form-item label="姓名:" prop="name">
                    <el-input v-model="formData.name" clearable placeholder="请准确输入学生姓名" />
                </el-form-item>
                <el-form-item label="性别:" prop="sex">
                    <el-select v-model="formData.sex" clearable placeholder="请选择性别">
                        <el-option
                                v-for="item in sexOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="年龄:" prop="age">
                    <el-input v-model.number="formData.age" clearable placeholder="请输入学生年龄"/>
                </el-form-item>
                <el-form-item label="负责人:" prop="managerName">
                    <el-input v-model="formData.managerName" clearable placeholder="学生管理、维护负责人" />
                </el-form-item>
                <el-form-item label="监护人:" prop="guardian">
                    <el-input v-model="formData.guardian" clearable placeholder="监护人身份" />
                </el-form-item>
                <el-form-item label="联系方式:" prop="phone">
                    <el-input v-model="formData.phone" clearable placeholder="请输入手机号码" />
                </el-form-item>
                <el-form-item label="推荐人:" prop="referee">
                    <el-input v-model="formData.referee" clearable placeholder="填写已报名学员名称，没有则不填" />
                </el-form-item>
                <el-form-item label="备注:" prop="comment">
                    <el-input  type="textarea" :rows="3" v-model="formData.comment" clearable placeholder="请输入" />
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button plain size="mini" @click="closeDialog">取 消</el-button>
                <el-button plain type="primary" size="mini" @click="beforeEnterDialog('formData')">确 定</el-button>
            </div>
        </el-dialog>
        <el-dialog :before-close="closeTrialDialog" :visible.sync="dialogTrialVisible" title="选择试听课程" width="30%">
            <el-form ref="formTrial" :model="formData" :rules="rulesTrial" label-position="right" label-width="100px">
                <el-form-item label="课程:" prop="courseType">
                    <el-select v-model="formData.courseType" clearable placeholder="请选择课程类别">
                        <el-option
                                v-for="item in courseOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                        />
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeTrialDialog">取 消</el-button>
                <el-button type="primary" @click="beforeEnterTrialDialog('formTrial')">确 定</el-button>
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
    const sexOptions = [
        {label: '男', value: '男'},
        {label: '女', value: '女'},
    ];
    const ageOptions = [
        {label: '无', value: 0},
        {label: '小于等于6岁', value: 6},
        {label: '大于6小于等于10', value: 10},
        {label: '大于10小于等于15', value: 15},
        {label: '大于15岁', value: 16},
    ];
    import {
        createPotential,
        deletePotential,
        updatePotential,
        findPotential,
        getPotentialList,
        upgradePotentialLevel,
        degradePotentialLevel,
        potentialToTrial,
    } from '@/api/z_potential' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    export default {
        name: 'Potential',
        mixins: [infoList],
        data() {
            return {
                listApi: getPotentialList,
                dialogFormVisible: false,
                dialogTrialVisible: false,
                type: '',
                deleteVisible: false,
                multipleSelection: [],
                courseOptions: courseOptions,
                sexOptions: sexOptions,
                ageOptions: ageOptions,
                formData: {
                    CreatedAt: new Date(),
                    address: '',
                    age: 0,
                    comment: '',
                    guardian: '',
                    managerName: '',
                    name: '',
                    phone: '',
                    referee: '',
                    sex: '',
                    level: 0,
                    courseType: '',
                    accesses: [],
                },
                rules: {
                    name: [{ required: true, message: '学生姓名不能为空', trigger: 'blur' }],
                    phone: [{ required: true, message: '联系方式不能为空', trigger: 'blur' }],
                },
                rulesTrial: {
                    courseType: [{ required: true, message: '课程不能为空', trigger: 'blur' }],
                }
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
                    this.deletePotential(row)
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
                const res = await deleteStudentByIds({ ids })
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
            async updatePotential(id) {
                const res = await findPotential({ ID: id})
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.repotential
                    this.dialogFormVisible = true
                }
            },
            async increaseLevel(id) {
                await upgradePotentialLevel({ ID: id})
                this.getTableData()
            },
            async reduceLevel(id) {
                await degradePotentialLevel({ ID: id})
                this.getTableData()
            },
            async turnTrial(id) {
                const res = await findPotential({ ID: id})
                this.type = 'update'
                if (res.code === 0) {
                    this.formData = res.data.repotential
                    this.dialogTrialVisible = true
                }
            },
            closeDialog() {
                this.dialogFormVisible = false
                this.formData = {
                    address: '',
                    age: 0,
                    comment: '',
                    guardian: '',
                    managerName: '',
                    name: '',
                    phone: '',
                    referee: '',
                    sex: '',
                    courseType: '',
                }
            },
            closeTrialDialog() {
                this.dialogTrialVisible = false
                this.formData = {
                    address: '',
                    age: 0,
                    comment: '',
                    guardian: '',
                    managerName: '',
                    name: '',
                    phone: '',
                    referee: '',
                    sex: '',
                    courseType: '',
                }
            },
            async deletePotential(row) {
                const res = await deletePotential({ ID: row.ID })
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
            beforeEnterTrialDialog(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.enterTrialDialog()
                    }
                })
            },
            async enterDialog() {
                let res
                switch (this.type) {
                    case "create":
                        res = await createPotential(this.formData)
                        break
                    case "update":
                        res = await updatePotential(this.formData)
                        break
                    default:
                        res = await createPotential(this.formData)
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
            async enterTrialDialog() {
                let res
                switch (this.type) {
                    case "create":
                        res = await createPotential(this.formData)
                        break
                    case "update":
                        res = await potentialToTrial(this.formData)
                        break
                    default:
                        res = await createPotential(this.formData)
                        break
                }
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '创建/更改成功'
                    })
                    this.closeTrialDialog()
                    this.getTableData()
                }
            },
            openDialog() {
                this.type = 'create'
                this.dialogFormVisible = true
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
    .popover-box {
        background: #feeeed;
        color: #444693;
        height: 600px;
        width: 420px;
        overflow: auto;
    }
</style>

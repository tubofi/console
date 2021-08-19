<template>
    <div class="shell">
        <el-card class="top-card" shadow="hover" :body-style="{padding: '0px'}">
            <el-row :gutter="20">
                <el-col :xs="4" :sm="3" :md="3" :lg="3" :xl="2">
                    <span class="card-img">
                        <img :src="userInfo.headerImg" alt="">
                    </span>
                </el-col>
                <el-col :xs="18" :sm="9" :md="9" :lg="9" :xl="14">
                    <div class="header-left">
                        <h4>欢迎登录，{{ userInfo.nickName }}， 请开始您一天的工作吧！</h4>
                        <p class="text-left">
                            <i class="el-icon-sunny el-icon" />
                            <span>今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</span>
                        </p>
                    </div>
                </el-col>
                <el-col :xs="0" :sm="12" :md="12" :lg="12" :xl="8">
                    <el-col :span="8">
                        <div class="car-item">
                            <span class="flow"><i class="el-icon-s-grid" /></span>
                            <span>今日流量 </span>
                            <b>13260</b>
                        </div>
                    </el-col>
                    <el-col :span="8">
                        <div class="car-item">
                            <span class="user-number">
                                <i class="el-icon-s-custom" />
                            </span>
                            <span>总用户 </span>
                            <b>48286</b>
                        </div>
                    </el-col>
                    <el-col :span="8">
                        <div class="car-item">
                            <span class="feedback">
                                <i class="el-icon-star-on" />
                            </span>
                            <span>好评率 </span>
                            <b>98%</b>
                        </div>
                    </el-col>
                </el-col>
            </el-row>
        </el-card>
        <!---下部分--->
        <el-row style="width: 100%; margin-top: 5px" type="flex" justify="space-between">
            <el-col :span="16">
                <!---班级分布表格--->
                <el-table
                        :data="remoteClassData"
                        style="width: 100%"
                        :header-cell-style="{background:'rgba(19,127,254,0.08)',color:'#3e4145',fontFamily:'MicrosoftYaHeiUI',fontSize:'14px',fontWeight:900}"
                        :row-style="{height: '100px'}"
                        :cell-style="{padding: '0'}"
                >
                    <!---第一列--->
                    <el-table-column prop="time" label="时间" align="center"></el-table-column>
                    <!---周课程分组总表头--->
                    <el-table-column label="周课程分布" align="center">
                        <!---周课程分组分表头--->
                        <el-table-column  prop="monday" label="星期一" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.monday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期二" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.tuesday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期三" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.wednesday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期四" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.thursday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期五" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.friday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期六" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.saturday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column  prop="monday" label="星期日" align="center">
                            <template slot-scope="scope">
                                <div v-for="(item, index) in scope.row.sunday" :key="index">
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratch-style" v-if="item.type === 'scratch'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="python-style" v-else-if="item.type === 'python'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="scratchJR-style" v-else-if="item.type === 'scratchJR'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="osmo-style" v-else-if="item.type === 'osmo'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="cpp-style" v-else-if="item.type === 'cpp'">
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                    <el-card shadow="hover"  :body-style="{padding: '0px'}" class="other-style" v-else>
                                        <p class="card-text">{{item.describe}}</p>
                                    </el-card>
                                </div>
                            </template>
                        </el-table-column>
                    </el-table-column>
                </el-table>
            </el-col>
            <!---课程分布表右侧部分--->
            <el-col :span="8" style="margin-left: 15px">
                <el-calendar v-model="value"></el-calendar>
                <div class="todo">
                    <todo-list />
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {mapGetters} from "vuex";
import TodoList from './component/todoList'
import {getProfileClass} from '@/api/z_profile'
export default {
    name: "index",
    components: {
        TodoList // TodoList
    },
    data() {
        return {
            remoteClassData: [],
            tableData: [
                {
                    time: "8:50 ~ 10:20",
                    monday: [],
                    saturday: [{type: "scratch", describe:"S-2-1 教室1 童志鹏"}],
                    sunday: [{type: "scratch", describe:"S-2-1 教室1 童志鹏"}, {type: "scratch", describe:"S-1-2 教室2 刘蒙"}]
                },
                {
                    time: "10:30 ~ 12:00",
                    monday: [],
                    saturday: [{type: "scratch", describe:"S-2-2 教室1 童志鹏"}, {type: "scratch", describe:"S-3-2 教室2 周鹏太"}],
                    sunday: [{type: "osmo", describe:"O-1-2 教室4 刘蒙"}, {type: "scratch", describe:"S-2-1 教室1 童志鹏"}, {type: "cpp", describe:"C-1-2 教室3 周鹏太"}]
                },
                {
                    time: "13:40 ~ 15:10",
                    saturday: [{type: "scratchJR", describe:"J-1-2 教室4 刘蒙"}, {type: "scratch", describe:"S-2-2 教室1 童志鹏"},{type: "python", describe:"P-2-2 教室3 周鹏太"}],
                    sunday: [{type: "scratchJR", describe:"J-1-2 教室4 刘蒙"}, {type: "scratch", describe:"S-2-3 教室2 童志鹏"}, {type: "python", describe:"P-2-2 教室3 周鹏太"}]
                },
                {
                    time: "15:20 ~ 16:50",
                    monday: [],
                    saturday: [{type: "osmo", describe:"O-1-2 教室4 刘蒙"}, {type: "scratch", describe:"S-2-2 教室1 童志鹏"}],
                    sunday: [{type: "osmo", describe:"O-1-2 教室4 刘蒙"}, {type: "scratch", describe:"S-2-2 教室2 童志鹏"}]
                },
                {
                    time: "17:00 ~ 18:30",
                    monday: [],
                    saturday: [{type: "scratch", describe:"S-2-2 教室1 童志鹏"},{type: "python", describe:"P-2-2 教室3 周鹏太"}],
                    sunday: [{type: "scratch", describe:"S-3-2 教室1 周鹏太"}, {type: "scratch", describe:"S-2-2 教室2 童志鹏"}]
                },
                {
                    time: "18:30 ~ 20:00",
                    monday: [{type: "scratch", describe:"S-1-2 教室2 周鹏太"}],
                    friday: [{type: "scratch", describe:"S-3-2 教室1 童志鹏"}, {type: "scratch", describe:"S-1-2 教室2 刘蒙"}, {type: "python", describe:"P-2-2 教室3 周鹏太"}]
                },
            ],
            value: new Date()
        }
    },
    computed: {
        ...mapGetters('user', ['userInfo'])
    },
    created() {
        this.onCreated()
    },
    methods: {
        async onCreated() {
            let res = await getProfileClass()
            if (res.code === 0) {
                this.remoteClassData = res.data
            }
        },
    },
}
</script>

<style lang="scss" scoped>
    .shell {
        background-color: rgb(243, 243, 243);
        padding: 0 0 0 0 ;
    }

    .top-card {
        height: 100px;
    }
    .card-img {
        width: 68px;
        height: 68px;
        display: inline-block;
        margin-left: 50px;
        margin-top: 5px;
        overflow: hidden;
    }
    img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
    }
    .header-left {
        margin-left: 20px;
        float: left;
        margin-top: 20px;
    }
    .text-left {
        color: #8c8c8c;
        margin-top: 0px;
    }
    .el-icon{
        margin-right: 8px;
        display:inline-block;
    }
    h4 {
        font-size: 20px;
        color: #262626;
        font-weight: 500;
        white-space: nowrap;
        word-break: break-all;
        text-overflow: ellipsis;
    }
    .car-item {
        text-align: right;
        margin-right: 50px;
        margin-top: 25px
    }
    b {
        display: block;
    }
    .flow {
        background-color: #fff7e8;
        border-color: #feefd0;
        color: #faad14;
    }
    .user-number {
        background-color: #ecf5ff;
        border-color: #d9ecff;
        color: #409eff;
    }

    .feedback {
        background-color: #eef9e8;
        border-color: #dcf3d1;
        color: #52c41a;
    }
    .card-text {
        color: #fffffb;
    }
    .scratch-style {
        background-color: #f47920;
    }
    .python-style {
        background-color: #009ad6;
    }
    .cpp-style {
        background-color: #1d953f;
    }
    .osmo-style {
        background-color: #ea66a6;
    }
    .scratchJR-style {
        background-color: #fcaf17;
    }
    .other-style {
        background-color: #a1a3a6;
    }
    .todo {
        width: 100%;
        height: 270px;
        padding: 0px;
        background-color: #fff;
    }
</style>
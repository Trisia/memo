<template>
    <div class="wrapper">
        <div class="container" v-loading="loadding">
            <div>
                <h1>MEMO</h1>
                <div class="title2">创建您的MEMO账号</div>
            </div>
            <div style="margin-top: 30px;">
                <g-input :err="pErr.username" @blur="handleUsernameBlur" v-model="param.username" placeholder="用户名" />
                <div class="hit" style="color: #d93025;" v-show="pErr.username">
                    <el-icon>
                        <Warning />
                    </el-icon>用户名为空
                </div>
            </div>
            <div style="margin-top: 15px;">
                <g-input :err="pErr.email" @blur="handleEmailBlur" v-model="param.email" placeholder="邮箱" />
                <div class="hit" style="color: #d93025;" v-show="pErr.email">
                    <el-icon>
                        <Warning />
                    </el-icon>邮箱格式错误
                </div>
            </div>
            <div style="margin-top: 15px; display: flex;justify-content: space-between;">
                <div style="width: 49%;">
                    <g-input type="password" :err="pErr.password" @blur="handlePasswordBlur" v-model="param.password"
                        placeholder="登录口令" />
                    <div class="hit" style="color: #d93025;" v-show="pErr.password">
                        <el-icon>
                            <Warning />
                        </el-icon>口令长度不足8位
                    </div>
                </div>
                <div style="width: 49%;">
                    <g-input type="password" :err="pErr.repeat" @blur="handleRepatBlur" v-model="repeat"
                        placeholder="确认" />
                    <div class="hit" style="color: #d93025;" v-show="pErr.repeat">
                        <el-icon>
                            <Warning />
                        </el-icon>两次口令输入不一致
                    </div>
                </div>
            </div>
            <div class="hit">请保证您的口令长度至少8位</div>
            <div class="ensure-container">
                <div>
                    <el-link style="font-size: 16px;" :underline="false" type="primary" @click="router.push('/')">
                        现有账号登录</el-link>
                </div>
                <div>
                    <el-button @click="handleRegister" size="large" style="font-size: 16px; width: 95px;"
                        type="primary">
                        注 册</el-button>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import GInput from '../components/GInput.vue';
import { Warning } from '@element-plus/icons-vue';
import axios from 'axios';

const router = useRouter();

const param = reactive({
    username: '',
    email: '',
    password: '',
});
const pErr = reactive({
    username: false,
    email: false,
    password: false,
    repeat: false,
});

const repeat = ref('');
const loadding = ref(false);


const handleRegister = () => {
    for (const key in pErr) {
        if (pErr[key]) {
            return;
        }
    }
    loadding.value = true;
    axios.post("./api/user/register", param).then(() => {
        ElMessageBox.alert('账号注册成功', '提示', {
            confirmButtonText: '前往登录',
            callback: () => {
                router.replace("/");
            },
        })
    }).catch((err) => {
        ElMessage.error(err.response.data)
    }).finally(() => {
        loadding.value = false;
    })
}

const handleUsernameBlur = (v) => {
    if (v != "") {
        pErr.username = false;
    } else {
        pErr.username = true;
    }
}
const handleEmailBlur = (v) => {
    let res = String(v)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
    if (res) {
        pErr.email = false;
    } else {
        pErr.email = true;
    }
}
const handlePasswordBlur = (v) => {
    if (v.length >= 8) {
        pErr.password = false;
    } else {
        pErr.password = true;
    }
}
const handleRepatBlur = (v) => {
    if (v == param.password) {
        pErr.repeat = false;
    } else {
        pErr.repeat = true;
    }
}
</script>


<style scoped>
.wrapper {
    min-width: 1040px;
    min-height: 99.9vh;
    position: relative;
}

.container {
    transform: translate(-50%, -50%);
    top: 50%;
    left: 50%;
    padding: 0;
    margin: 0;
    position: absolute;
    min-width: 550px;
    min-height: 500px;
    border-radius: 4px;
    box-shadow: 0px 0px 6px rgba(0, 0, 0, .12);
    padding: 50px 45px;
}

.title2 {
    font-family: 'Noto Sans Myanmar UI', arial, sans-serif;
    font-size: 24px;
    font-weight: 400;
    line-height: 1.3333;
}

.ensure-container {
    display: flex;
    justify-content: space-between;
    padding: 0 20px;
    margin-top: 50px;
    margin-bottom: 30px;
    align-items: center;
}

.hit {
    margin-left: 15px;
    margin-top: 5px;
}
</style>

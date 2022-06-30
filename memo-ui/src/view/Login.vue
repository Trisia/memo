<template>
    <div class="wrapper">
        <div class="container" v-loading="loadding">
            <div>
                <h1 style="font-size: 50px; text-align: center;">MEMO</h1>
            </div>
            <g-input :err="pErr.username" @blur="handleUsernameBlur" v-model="username" placeholder="用户名或邮箱" />
            <div class="hit" style="color: #d93025;" v-show="pErr.username">
                <el-icon>
                    <Warning />
                </el-icon>用户名或邮箱为空
            </div>
            <g-input :err="pErr.password" @blur="handlePasswordBlur" style="margin-top: 15px;" v-model="password"
                type="password" placeholder="登录口令" />
            <div class="hit" style="color: #d93025;" v-show="pErr.password">
                <el-icon>
                    <Warning />
                </el-icon>口令不能为空
            </div>
            <div class="ensure-container">
                <div>
                    <el-link @click="router.push('/user/register')" style="font-size: 16px;" :underline="false"
                        type="primary">创建账号</el-link>
                </div>
                <div>
                    <el-button @click="handleLogin" size="large" style="font-size: 16px; width: 95px;" type="primary">登
                        录</el-button>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup>

import { ref, reactive } from "vue";
import { ElButton, ElLink, ElMessage } from 'element-plus';
import { Warning } from '@element-plus/icons-vue';
import axios from "axios";
import { useRouter } from 'vue-router';
import GInput from "../components/GInput.vue";


const loadding = ref(false)
const username = ref('');
const password = ref('');
const router = useRouter();

const pErr = reactive({
    username: false,
    password: false,
});


const handleLogin = () => {
    if (username.value == '') {
        ElMessage.error("用户名为空");
        return
    }
    if (password.value == '') {
        ElMessage.error("登录口令为空");
        return
    }

    loadding.value = true;
    axios.post("./api/auth", {
        username: username.value,
        password: password.value,
    }).then(({ data }) => {
        localStorage.setItem('token', data);
        router.replace("/main")
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

const handlePasswordBlur = (v) => {
    if (v != '') {
        pErr.password = false;
    } else {
        pErr.password = true;
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
    min-width: 400px;
    min-height: 500px;
    border-radius: 4px;
    box-shadow: 0px 0px 6px rgba(0, 0, 0, .12);
    padding: 30px 45px;
}

.ensure-container {
    display: flex;
    justify-content: space-between;
    padding: 0 20px;
    margin-top: 50px;
    align-items: center;
}

.hit {
    margin-left: 15px;
    margin-top: 5px;
}
</style>

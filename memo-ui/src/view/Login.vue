<template>
    <div class="wrapper">
        <div class="container" v-loading="loadding">
            <div>
                <h1 style="font-size: 50px; text-align: center;">MEMO</h1>
            </div>
            <g-input v-model="username" placeholder="用户名或邮箱"/>
            <g-input style="margin-top: 15px;" v-model="password" type="password" placeholder="登录口令"/>
            <div class="ensure-container">
                <div>
                    <el-link style="font-size: 16px;" :underline="false" type="primary">创建账号</el-link>
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

import { ref } from "vue"
import { ElButton, ElLink } from 'element-plus'
import axios from "axios";
import { useRouter } from 'vue-router'
import GInput from "../components/GInput.vue";

const loadding = ref(false)
const username = ref('');
const password = ref('');
const router = useRouter();


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
        router.replace("/home")
    }).catch((err) => {
        ElMessage.error(err.response.data)
    }).finally(() => {
        loadding.value = false;
    })
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
    padding: 30px 25px;
}

.ensure-container {
    display: flex;
    justify-content: space-between;
    padding: 0 20px;
    margin-top: 50px;
    align-items: center;
}
</style>

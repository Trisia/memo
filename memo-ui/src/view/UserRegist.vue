<template>
    <div class="wrapper">
        <div class="container" v-loading="loadding">
            <div>
                <h1 style="font-size: 50px; text-align: center;">MEMO</h1>
            </div>

            <div class="imformation">
                <input class="gg-input" type="text" v-model="username" autocomplete="off" @focus="focusUsername = true"
                    @blur="focusUsername = false">
                <div :class="usernameHitCss">邮箱或用户名</div>
            </div>

            <div class="imformation" style="margin-top: 15px;">
                <input @keyup.enter="handleLogin" class="gg-input" v-model="password" type="password" autocomplete="off"
                    @focus="focusPassword = true" @blur="focusPassword = false">
                <div :class="passwordHitCss">登录口令</div>
            </div>
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

import { ref, computed } from "vue"
import { ElButton, ElLink } from 'element-plus'
import axios from "axios";
import { useRouter } from 'vue-router'

const loadding = ref(false)
const focusUsername = ref(false);
const focusPassword = ref(false);
const username = ref('');
const password = ref('');
const router = useRouter();

// 根据状态改变用户名的提示消息位置和颜色
const usernameHitCss = computed(() => {
    let res = { 'gg-div': true }
    if (focusUsername.value) {
        res['gg-hit-focus'] = true
    }
    if (username.value != '' && !focusUsername.value) {
        res['gg-hit-focus-content'] = true
    }
    return res;
});

// 根据状态改变口令的提示消息位置和颜色
const passwordHitCss = computed(() => {
    let res = { 'gg-div': true }
    if (focusPassword.value) {
        res['gg-hit-focus'] = true
    }
    if (password.value != '' && !focusPassword.value) {
        res['gg-hit-focus-content'] = true
    }
    return res;
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

.imformation {
    width: 100%;
    position: relative;
}

.gg-input {
    width: 100%;
    height: 56px;
    font-size: 16px;
    padding: 13px 15px 13px 15px;
    box-sizing: border-box;
    outline: none;
    border-radius: 8px;
    border: 1px solid #dadce0;
}

.gg-input:focus {
    border: 2px solid #1a73e8;
}

.gg-div {
    position: absolute;
    /*设置绝对定位， 那么给父级元素添加相对定位*/
    top: 17px;
    left: 12px;
    padding: 0 5px;
    font-size: 16px;
    font-weight: 400;
    font-family: roboto, 'Noto Sans Myanmar UI', arial, sans-serif;
    letter-spacing: 0.1px;
    /*字体间的空隙*/
    color: #5f6368;
    pointer-events: none;
    transition: all 0.15s ease;
}

.gg-hit-focus {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #1a73e8;
}

.gg-hit-focus-content {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #5f6368;
}

.ensure-container {
    display: flex;
    justify-content: space-between;
    padding: 0 20px;
    margin-top: 50px;
    align-items: center;
}
</style>

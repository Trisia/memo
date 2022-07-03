import { createRouter, createWebHashHistory } from 'vue-router';

import Login from "../view/Login.vue"
import Main from "../view/Main/Main.vue"
import UserRegister from "../view/UserRegister.vue"
import Test from "../view/Test.vue"

const routes = [
  { path: "/", component: Login },
  { path: "/Main", component: Main },
  { path: "/test", component: Test },
  { path: "/user/register", component: UserRegister },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


export default router

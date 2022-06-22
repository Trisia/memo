import { createRouter, createWebHashHistory } from 'vue-router';

import Login from "../view/Login.vue"
import Home from "../view/Home.vue"
import UserRegist from "../view/UserRegist.vue"

const routes = [
  { path: "/", component: Login },
  { path: "/home", component: Home },
  { path: "/user/register", component: UserRegist },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


export default router

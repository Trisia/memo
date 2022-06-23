import { createRouter, createWebHashHistory } from 'vue-router';

import Login from "../view/Login.vue"
import Home from "../view/Home.vue"
import UserRegister from "../view/UserRegister.vue"

const routes = [
  { path: "/", component: Login },
  { path: "/home", component: Home },
  { path: "/user/register", component: UserRegister },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


export default router

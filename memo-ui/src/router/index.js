import { createRouter, createWebHashHistory } from 'vue-router';

import LoginPage from "../view/LoginPage.vue"

const routes = [
  { path: "/", component: LoginPage }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


export default router

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import useInterceptor from './useInterceptor.js'

useInterceptor();

createApp(App).use(router).mount('#app')

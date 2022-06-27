import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import useInterceptor from './useInterceptor.js'


useInterceptor();

const app = createApp(App);
app.use(router).mount('#app')

import axios from 'axios'
// import router from './router';


// http request 拦截器
axios.interceptors.request.use(
    config => {
        if (localStorage.token) { //判断token是否存在
            config.headers.Token = localStorage.token;  //将token设置成请求头
        }
        return config;
    },
    err => {
        return Promise.reject(err);
    }
);

// http response 拦截器
axios.interceptors.response.use(
    response => {
        if (response.data.errno === 999) {
            router.replace('/');
            console.log("token过期");
        }
        return response;
    },
    error => {
        return Promise.reject(error);
    }
);


export default axios;
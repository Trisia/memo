import axios from 'axios'

export default function () {
    // http request 拦截器
    axios.interceptors.request.use(
        config => {
            // 判断token是否存在
            if (localStorage.getItem('token')) { 
                // 将token设置成请求头
                config.headers.Token = localStorage.getItem('token');  
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
            if (response.status == 401) {
                router.replace('/');
                console.log("token过期");
            }
            return response;
        },
        error => {
            return Promise.reject(error);
        }
    );
}

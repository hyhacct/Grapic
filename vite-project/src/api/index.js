// 二次封装配置

import axios from "axios";


// // 二次封装axios;
// const service = axios.create({
//     baseURL: "http://localhost:9099/api/v50", // 基础 URL
//     timeout: 10000, // 请求超时时间
// });


const service = axios.create({
    baseURL: "/api/v50", // 基础 URL
    timeout: 5000, // 请求超时时间
});


// 添加响应拦截器
service.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
}, function (error) {

    console.log(error);

    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    // if (error.response.status === 401) {
    //     // 401 未授权
    //     router.push('/');
    //     message.error(error.response.data.msg);
    // }

    return Promise.reject(error);
});



export default service;

import axios from 'axios'
// axios.defaults.headers.post['Content-Type'] = 'text/plain'
// 创建拦截器
// axios.defaults.baseURL = 'http://121.89.200.190:18081/'
axios.defaults.baseURL = 'http://127.0.0.1:18081/'
const service = axios.create({
    // baseURL: 'http://121.89.200.190:18081/'
    baseURL: 'http://127.0.0.1:18081/'
})

// 添加请求拦截器
service.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error)
})

// 添加响应拦截器
service.interceptors.response.use(function (response) {
    // 对响应数据做些什么
    return response
}, function (error) {
    // 对响应错误做些什么
    return error
})

export default service

import axios from 'axios';
let newAxios = axios.create({
    headers: {'Content-Type': 'multipart/form-data;charset=utf-8'}
});
newAxios.defaults.transformRequest = [function (data, config) {
    switch (config['Content-Type'].toLowerCase()) {
        case 'multipart/form-data;charset=utf-8':{
            return data;
        }
    }
}]

function upload(url, formdata) {
    let promise;
    return new Promise((resolve, reject) => {
        promise = newAxios.post(url, formdata);
        promise.then((response) => {
            resolve(response.data);
        }).catch(error => {
            reject(error);
        })
    })
}

// let globalUrl = "http://127.0.0.1:18081/" // 地址可以改成自己项目中的地址
let globalUrl = "http://121.89.200.190:18081/" // 地址可以改成自己项目中的地址

// 接口列表，供其它页面调用使用，如还有其它上传接口，可复制一份改下接口地址根函数名就行
export const uploadFile = (formdata) => upload(globalUrl + 'files/images', formdata);

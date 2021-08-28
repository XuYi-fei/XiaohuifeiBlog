import service from '../utils/request'

const baseUrl = '#'
export function postLogin (ruleForm) {
    return service.request({
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        method: 'post',
        url: '/users/login',
        data: {
            userId: ruleForm.userId,
            password: ruleForm.pass
        }
    }).then(function (response) {
        if (response.data.status === 0) {
            alert(response.data.data.Name + ' ,你好')
            window.localStorage.setItem('token', response.data.data.Token)
            window.location.reload()
            return response.data
        }else{
            return response.data
        }
    }).catch(function (response) {
        console.log(response.data.msg)
    })
}

export function postRegister (ruleForm) {
    return service.request({
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        method: 'post',
        url: '/users/register',
        data: {
            userName: ruleForm.userName,
            passWord: ruleForm.pass,
            userId: ruleForm.userId,
            question: ruleForm.question,
            answer: ruleForm.answer
        }
    }).then(function (response) {
        return response
    })
}

export function loginCheck () {
    const token = window.localStorage.getItem('token')
    return service.request({
        method: 'post',
        url: '/users/index',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            Token: token
        }
    }).then(function (response) {
        try {
            if (response.data.status === 0) {
                return response.data
            } else {
                window.localStorage.removeItem('token')
                return response.data
            }
        } catch (e) {
            window.localStorage.removeItem('token')
        }
    }).catch(function (response) {
        alert("未知错误")
    })
}

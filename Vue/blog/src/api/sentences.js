import service from "../utils/request";

export function submitSentence (body) {
    return service.request({
        method: 'post',
        url: '/sentences/upload',
        data: {
            userId: body.userId,
            content: body.content,
            author: body.author
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        return response.data
    }).catch(function (response) {
        alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
    })
}

export function getAllSentences(pageNo, mode, userId, order="time"){
    return service.request({
        method: 'post',
        url: '/sentences/all',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
            orderType: order,
            pageNo: pageNo,
            mode: mode,
            userId: userId
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("暂无数据")
    })
}

export function getSentencesToShow(){
    return service.request({
        method: 'get',
        url: '/sentences/show',
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getSentencesPageNum(mode, userId){
    return service.request({
        method: 'post',
        url: '/sentences/getPageNum',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
            mode: mode,
            userId: userId
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getSentencesPageNumPersonal(userId){
    return service.request({
        method: 'post',
        url: '/sentences/getPageNumPersonal',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
            userId: userId
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getAllSentencesPersonal(userId, pageNo){
    return service.request({
        method: 'post',
        url: '/sentences/getSentencesPersonal',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: {
            pageNo: pageNo,
            userId: userId
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("暂无数据")
    })
}

// 用户删除句子
export function DeleteSentencesPersonal(sentence){
    return service.request({
        method: 'post',
        url: '/sentences/delete/Personal',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        data: {
            id: sentence.id,
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

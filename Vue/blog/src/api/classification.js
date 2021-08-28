import service from "../utils/request";

export function getField (mode, userId) {
    return service.request({
        method: 'post',
        url: '/classification/getField',
        data: {
            mode: mode,
            userId: userId
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

export function getClassification (mode, userId, field) {
    return service.request({
        method: 'post',
        url: '/classification/getClassification',
        data: {
            mode: mode,
            userId: userId,
            field: field
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

export function getTagBlogs (mode, userId, field, classification) {
    return service.request({
        method: 'post',
        url: '/classification/getTagBlogs',
        data: {
            mode: mode,
            userId: userId,
            field: field,
            classification: classification
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

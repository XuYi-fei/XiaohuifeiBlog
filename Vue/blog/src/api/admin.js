import service from "../utils/request";

export function getAllSentencesTmp(){
    return service.request({
        method: 'get',
        url: '/sentences/sentencesTmp',
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("暂无数据")
    })
}

export function PassSentences(sentence){
    return service.request({
        method: 'post',
        url: '/admin/sentencesTmp/pass',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            userName: sentence.userName,
            userId: sentence.userId,
            content: sentence.content,
            id: sentence.id,
            author: sentence.author

        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function DeleteSentences(sentence){
    return service.request({
        method: 'post',
        url: '/admin/sentences/delete',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            userName: sentence.userName,
            userId: sentence.userId,
            content: sentence.content,
            id: sentence.id,
            author: sentence.author
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function DeleteSentencesTmp(sentence){
    return service.request({
        method: 'post',
        url: '/admin/sentencesTmp/delete',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            userName: sentence.userName,
            userId: sentence.userId,
            content: sentence.content,
            id: sentence.id,
            author: sentence.author
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getAllBlogsTmp(){
    return service.request({
        method: 'get',
        url: '/blogs/blogsTmp',
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("暂无数据")
    })
}

export function PassBlogs(blog){
    console.log(blog)
    return service.request({
        method: 'post',
        url: '/admin/blogsTmp/pass',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            author: blog.author,
            blogId: blog.blogId,
            classification: blog.classification,
            content: blog.filePath,
            digest: blog.digest,
            field: blog.field,
            id: blog.id,
            tag: blog.tag,
            title: blog.title,
            userId: blog.userId

        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function DeleteBlogs(blog){
    return service.request({
        method: 'post',
        url: '/admin/blogsTmp/delete',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            author: blog.author,
            blogId: blog.blogId,
            classification: blog.classification,
            content: blog.content,
            digest: blog.digest,
            field: blog.field,
            id: blog.id,
            tag: blog.tag,
            title: blog.title,
            userId: blog.userId
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getBlogTmp(blogId){
    return service.request({
        method: 'post',
        url: '/blogs/getBlogTmpById',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
            // 'Content-Type': 'application/json'
        },
        data: {
            blogId: blogId,
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

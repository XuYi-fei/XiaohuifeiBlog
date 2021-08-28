import service from "../utils/request";

export function submitBlog (body) {
    return service.request({
        method: 'post',
        url: '/blogs/upload',
        data: {
            body: body
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        try {
            if (response.data.status === 0) {
                alert("发表成功!!")
                window.location.reload()
            } else {
                alert(response.data.msg)
            }
        } catch (e) {
            alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
        }
    }).catch(function (response) {
        alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
    })
}


export function updateBlog (body) {
    return service.request({
        method: 'post',
        url: '/blogs/update',
        data: {
            blogId: body.blogId,
            title: body.title,
            digest: body.digest,
            content: body.content,
            userId: body.userId,
            tag: body.tag,
            classification: body.classification,
            field: body.field
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        try {
            if (response.data.status === 0) {
                alert("更新成功!!")
                window.location.reload()
            } else {
                alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
            }
        } catch (e) {
            alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
        }
    }).catch(function (response) {
        alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
    })
}

export function deleteBlog (body) {
    return service.request({
        method: 'post',
        url: '/blogs/deleteBlog',
        data: {
            blogId: body.blogId,
            title: body.title,
            digest: body.digest,
            content: body.content,
            userId: body.userId,
            tag: body.tag,
            classification: body.classification,
            field: body.field
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


export function getLatestBlogs(pageNo, mode, userId){
    return service.request({
        method: 'post',
        url: '/blogs/latest/all',
        data: {
            userId: userId,
            mode: mode,
            pageNo: pageNo
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getBlogById(blogId){
    return service.request({
        method: 'post',
        url: '/blogs/blogId',
        data: {
            blogId: blogId
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("未知错误")
    })
}

export function getBlogByUser(UserId){
    return service.request({
        method: 'post',
        url: '/blogs/getByUserId',
        data: {
            userId: UserId
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("文章不存在!")
    })
}

export function getBlogPageNum(userId, mode){
    return service.request({
        method: 'post',
        url: '/blogs/getPageNum',
        data: {
            userId: userId,
            mode: mode
        },
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }).then(function (response) {
        return response.data
    }).catch(function () {
        alert("文章不存在!")
    })
}

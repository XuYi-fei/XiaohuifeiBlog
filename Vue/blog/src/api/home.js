import service from "../utils/request";

export function getBlogsInfo (body) {
    return service.request({
        method: 'get',
        url: '/blogs/getBlogsInfo',
    }).then(function (response) {
        return response.data
    }).catch(function (response) {
        alert("网络可能出了点小故障- -|| 请休息一会再试下吧!")
    })
}

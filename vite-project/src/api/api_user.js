import request from './index'


// 发送新数据到服务器
export const ApiUserSubmit = (data) => {
    return request.post('/user/submit', data)
}



// 查询是否需要密码
export const ApiUserNeedPassword = (guid) => {
    return request.get(`/user/needPassword?guid=${guid}`)
}



// 获取文章的详细信息
export const ApiUserGetComments = (guid, password) => {
    return request.get(`/user/get_comments?guid=${guid}&password=${password}`)
}


// 获取文章的主体内容
export const ApiUserGetContent = (guid, password) => {
    return request.get(`/user/get_content?guid=${guid}&password=${password}`)
}



// 提交评论
export const ApiUserIssue = (data) => {
    return request.post("/user/issue", data)
}


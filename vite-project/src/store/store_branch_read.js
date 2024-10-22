import { defineStore } from "pinia";

export const useStoreBranchRead = defineStore("useStoreBranchRead", {
    state: () => {
        return {
            showDialogInputPassword: false, // 展开密码输入框
            showMessageError: false, // 显示错误信息
            showMessageText: "", // 显示提示信息

            password: "", // 密码输入框内容

            respond_comments: {
                allow_comment: false, // 是否允许评论
                comments: [], // 评论列表
                expire_time: "", // 过期时间
                ip: "", // IP地址
                type: "", // 文章类型
                url: {
                    api: "", // 接口
                    web: "" // 网页
                }
            }, // 评论列表
            respond_content: "", // 评论内容

            loading_comments: false, // 评论列表加载中状态
            loading_content: false, // 评论内容加载中状态

            respond_msg_comments: "", // 评论列表加载失败信息
            respond_msg_content: "", // 评论内容加载失败信息

            showMessageComments: false, // 显示评论列表加载失败信息
            showMessageContent: false, // 显示评论内容加载失败信息

            loader_comments: false, // 评论列表加载动画
            loader_content: false, // 评论内容加载动画

            item_comments: [], // 文章评论列表
        };
    },
});



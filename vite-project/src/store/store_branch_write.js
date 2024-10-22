import { defineStore } from "pinia";

export const useStoreBranchWrite = defineStore("useStoreBranchWrite", {
    state: () => {
        return {
            form: {
                type: "text", // 文本类型
                once: false, // 只允许访问一次
                exceed_time: -1, // 超过时间限制
                password: "", // 访问密码
                ips_input: "", // 限制访问IP(白名单)
                context: "", // 内容
                allow_comment: true, // 允许评论
            },


            showDialog: false, // 是否显示消息提示
            showDialogLoading: false, // 是否显示加载中提示

            showMessageSuccess: false, // 消息提示内容
            showMessageError: false, // 消息提示内容
            showMessageContent: "", // 消息提示内容
        };
    },
});



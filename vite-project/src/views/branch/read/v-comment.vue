<template>
    <div>

        <v-card>
            <div v-if="store.loader_content">
                <v-skeleton-loader :elevation="2" type="list-item-avatar@4"></v-skeleton-loader>
            </div>

            <div v-else>
                <v-list :items="store.item_comments" lines="three" item-props :max-height="400">
                    <template v-slot:subtitle="{ subtitle }">
                        <div v-html="subtitle"></div>
                    </template>
                </v-list>
            </div>


            <v-divider class="border-opacity-30"></v-divider>

            <div style="padding: 30px;">
                <v-text-field v-model="form.email" label="邮箱" prepend-icon="mdi-email-outline" variant="outlined"
                    density="compact"></v-text-field>

                <v-textarea v-model="form.text" prepend-icon="mdi-comment-minus-outline" label="评论" row-height="25"
                    rows="3" variant="outlined" auto-grow shaped :counter="300" density="compact"></v-textarea>

                <v-btn prepend-icon="mdi-send-check-outline" block variant="tonal" @click="submitComment"
                    :loading="loading">
                    发送
                </v-btn>
            </div>
        </v-card>
    </div>


    <div>
        <v-snackbar v-model="show_message_success" variant="text">
            <v-alert icon="mdi-check" :text="message_success_msg" type="success"></v-alert>
        </v-snackbar>

        <v-snackbar v-model="show_message_err" variant="text">
            <v-alert icon="mdi-close" :text="message_err_msg" type="error"></v-alert>
        </v-snackbar>
    </div>

</template>

<script setup>
import { ref } from 'vue'
import { useStoreBranchRead } from "@/store/store_branch_read"
import { Base64 } from 'js-base64';
import { useRoute } from 'vue-router'
import { ApiUserIssue } from "@/api/api_user";
import user from "@/assets/60ec05e4370111626080740292.png";


// 获取当前页面的路由参数
const guid = useRoute().params.guid;
const store = useStoreBranchRead();
const form = ref({
    text: "",
    email: "",
});


const message_err_msg = ref("");
const message_success_msg = ref("");

const show_message_success = ref(false);
const show_message_err = ref(false);

const loading = ref(false);


const showComment = () => {
    store.item_comments = []; // 清空之前的评论数据

    // 如果评论功能被关闭，则没必要处理评论数据了
    if (!store.respond_comments.allow_comment) {
        store.item_comments = [
            { type: 'subheader', title: '历史评论' },
            {
                subtitle: `<span class="d-flex justify-center text-error">评论功能已被关闭</span>`,
            },
        ]
        return;
    }


    // 如果没有评论数据，则提示用户
    if (store.respond_comments.comments?.length == 0 || !store.respond_comments.comments) {
        store.item_comments = [
            { type: 'subheader', title: '历史评论' },
            {
                subtitle: `<span class="d-flex justify-center">当前暂无评论哦~</span>`,
            },
        ]
        return;
    }

    // 从 store 中获取评论数据
    for (let i = 0; i < store.respond_comments.comments?.length; i++) {

        const comment = store.respond_comments.comments[i]

        store.item_comments.push({
            // prependAvatar: "https://cdn.vuetifyjs.com/images/lists/2.jpg",
            prependAvatar: user,
            title: Base64.decode(comment.name),
            subtitle: `<span class="text-primary">${comment.time}</span> &mdash; ${Base64.decode(comment.text)}`,
        })
        store.item_comments.push({ type: 'divider', inset: true },)

    }
}


defineExpose({
    showComment
})




// 提交评论
const submitComment = () => {

    if (!form.value.text || !form.value.email) {
        message_err_msg.value = "评论内容和邮箱都要提供哦~";
        show_message_err.value = true;
        return;
    }

    // 评论内容需要进行base64加密一下
    var data = {
        text: Base64.encode(form.value.text),
        email: Base64.encode(form.value.email),
        guid: guid, // 当前文档的guid标识，从路由获取
    }

    // 开始加载
    loading.value = true;

    setTimeout(() => {
        ApiUserIssue(data)
            .then(res => {
                message_success_msg.value = res?.data?.msg || "评论成功";
                show_message_success.value = true;
            })
            .catch(err => {
                message_err_msg.value = err?.response?.data?.msg || err.message;
                show_message_err.value = true;
            })
            .finally(() => {
                loading.value = false;
            });
    }, 2000);

}

</script>

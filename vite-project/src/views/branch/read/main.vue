<template>
    <div>

        <v-layout class="rounded rounded-md">
            <v-app-bar title="App Read">
                <template v-slot:append>
                    <v-btn variant="plain" prepend-icon="mdi-keyboard-return" density="compact"
                        @click="router.push({ name: 'branch_write' })">
                        返回
                    </v-btn>
                </template>
            </v-app-bar>

            <v-main>
                <div style="padding: 20px;">
                    <v-row>
                        <v-col cols="8">
                            <ViewInput />
                        </v-col>

                        <v-col cols="4">
                            <ViewV />

                            <ViewComment ref="viewComment" />
                        </v-col>
                    </v-row>
                </div>
            </v-main>
        </v-layout>


        <!-- 模态框-输入密码 -->
        <div>
            <v-dialog activator="parent" max-width="500" persistent v-model="store.showDialogInputPassword">
                <template v-slot:default="{ isActive }">
                    <v-card rounded="lg">
                        <v-card-title class="d-flex justify-space-between align-center">
                            <div class="text-h5 text-medium-emphasis ps-2">
                                需要提供访问密码
                            </div>

                            <v-btn icon="mdi-close" variant="text" @click="closeInputPassword"></v-btn>
                        </v-card-title>

                        <v-divider class="mb-4"></v-divider>

                        <v-card-text>

                            <div class="mb-2">请在此输入访问密码，以查看该页面内容。</div>

                            <v-text-field :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                                :type="show ? 'text' : 'password'" class="input-group--focused" label="密码"
                                name="input-10-2" @click:append="show = !show" v-model="store.password"></v-text-field>
                        </v-card-text>

                        <v-divider class="mt-2"></v-divider>

                        <v-card-actions class="my-2 d-flex justify-end">
                            <v-btn class="text-none" rounded="xl" text="返回" @click="closeInputPassword"></v-btn>

                            <v-btn class="text-none" color="primary" rounded="xl" text="提交" variant="flat"
                                @click="getComments" :loading="store.loading_comments"></v-btn>
                        </v-card-actions>
                    </v-card>
                </template>
            </v-dialog>
        </div>




        <!-- 消息条-错误 -->
        <div>
            <v-snackbar v-model="store.showMessageError" variant="text">
                <v-alert icon="mdi-close" :text="store.showMessageText" type="error"></v-alert>
            </v-snackbar>

            <v-snackbar v-model="store.showMessageComments" variant="text">
                <v-alert icon="mdi-close" :text="store.respond_msg_comments" type="error"></v-alert>
            </v-snackbar>

            <v-snackbar v-model="store.showMessageContent" variant="text">
                <v-alert icon="mdi-close" :text="store.respond_msg_content" type="error"></v-alert>
            </v-snackbar>
        </div>

    </div>
</template>

<script setup>
import ViewV from './v-view.vue'
import ViewInput from "./v-input.vue"
import ViewComment from "./v-comment.vue"
import { useRoute, useRouter } from 'vue-router'
import { ApiUserNeedPassword, ApiUserGetComments, ApiUserGetContent } from "@/api/api_user"
import { useStoreBranchRead } from "@/store/store_branch_read"
import { ref } from 'vue'
import CryptoJS from 'crypto-js';


// 获取当前页面的路由参数
const guid = useRoute().params.guid;
const store = useStoreBranchRead();
const show = ref(false);
const router = useRouter();
const viewComment = ref(null);


// 生命周期内判断阅读是否需要提供密码
const handleNeedPassword = () => {
    store.loader_content = true;

    ApiUserNeedPassword(guid)
        .then(res => {
            store.showDialogInputPassword = res.data.data == true;
            if (!res.data.data) {
                getComments();
            }

            console.log(res.data);
        })
        .catch(err => {
            store.showMessageText = err?.response?.data?.msg || err.message;
            store.showMessageError = true;
            console.log("error: ", err);
        })
}
handleNeedPassword();



// 用户取消输入密码，将其导航到首页去
const closeInputPassword = () => {
    router.push("/view");
}



// 获取评论数据,需要提供密码
const getComments = () => {
    store.loading_comments = true;

    // 密码加密
    if (store.showDialogInputPassword) var password = CryptoJS.MD5(store.password).toString();

    // 延迟2秒开始请求
    setTimeout(() => {
        ApiUserGetComments(guid, password)
            .then(res => {
                console.log("success comments: ", res);
                store.respond_comments = res.data.data;
                store.showDialogInputPassword = false; // 关闭输入密码的弹窗

                viewComment.value?.showComment(); // 显示评论组件

                getContent(password);
            })
            .catch(err => {
                store.respond_msg_comments = err?.response?.data?.msg || err.message;
                store.showMessageComments = true;
                console.log("error comments: ", err);
            })
            .finally(() => {
                store.loading_comments = false;
            });
    }, 2000);
}



// 获取消息数据,需要提供密码
const getContent = (password) => {
    store.loading_content = true;

    ApiUserGetContent(guid, password)
        .then(res => {
            console.log("success content data: ", res.data);

            store.respond_content = res.data.toString();
            console.log("success content: ", res);
        })
        .catch(err => {
            store.respond_msg_content = err?.response?.data?.msg || err.message;
            store.showMessageContent = true;
            console.log("error content: ", err);
        })
        .finally(() => {
            store.loading_content = false;
            store.loader_content = false;
        });
}




// 输出当前页面的路由参数
console.log("router params: ", guid)

</script>
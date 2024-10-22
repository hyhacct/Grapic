<template>
    <div style="padding: 15px;">


        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">类型</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        选择类型, 默认为文本, 如果为代码, 则在展示时会有语法高亮辅助阅读
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>

            <v-btn-toggle v-model="store.form.type" color="deep-purple-accent-3" group size="small" base-color="#424242"
                density="compact" mandatory>
                <v-btn value="text">
                    文本
                </v-btn>

                <v-btn value="code">
                    代码
                </v-btn>

                <v-btn value="logs">
                    日志
                </v-btn>
            </v-btn-toggle>
        </div>


        <br>
        <v-divider></v-divider>


        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">访问一次</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        若是启用此项, 则只能访问一次, 阅读后文章将立即被删除
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>

            <v-btn-toggle v-model="store.form.once" color="deep-purple-accent-3" group size="small" base-color="#424242"
                density="compact" mandatory>
                <v-btn :value="true">
                    是
                </v-btn>

                <v-btn :value="false">
                    否
                </v-btn>
            </v-btn-toggle>
        </div>


        <br>
        <v-divider></v-divider>


        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">开放评论</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        是否允许游客留下评论, 默认开启, 关闭后任何人无法评论
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>

            <v-btn-toggle v-model="store.form.allow_comment" color="deep-purple-accent-3" group size="small"
                base-color="#424242" density="compact" mandatory>
                <v-btn :value="true">
                    是
                </v-btn>

                <v-btn :value="false">
                    否
                </v-btn>
            </v-btn-toggle>
        </div>


        <br>
        <v-divider></v-divider>

        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">生效时长</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        设置访问生效时长, 超时后此内容将不可访问, 若是留空或者选择[永久]则永远不失效
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>


            <v-select v-model="store.form.exceed_time" label="请选择" :items="options_time" variant="solo-filled"
                density="compact" />
        </div>

        <v-divider></v-divider>

        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">访问密码</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        访问密码, 留空则无密码, 否则只有输入密码后方可访问
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>

            <v-text-field v-model="store.form.password" type="password" label="可空" variant="solo-filled"
                density="compact" />
        </div>

        <v-divider></v-divider>


        <div>
            <div class="d-flex justify-start">
                <div :class="['text-button', 'pa-2']">白名单IP</div>

                <v-btn density="compact" icon="mdi-help-circle-outline" style="margin-top: 12px;">
                    <v-tooltip activator="parent" location="end">
                        只希望某些IP可以访问, 则在此填写, 多个IP使用换行分隔, 留空则不限制IP访问
                    </v-tooltip>
                    <v-icon size="small">
                        mdi-help-circle-outline
                    </v-icon>
                </v-btn>
            </div>


            <v-textarea v-model="store.form.ips_input" label="注意换行哦" variant="solo-filled" density="compact" />
        </div>

        <div>
            <v-btn color="deep-purple-accent-3" block @click="submitForm()">
                发送
            </v-btn>
        </div>

        <div>
            <v-dialog v-model="store.showDialog" max-width="500" theme="error" transition="dialog-top-transition">
                <template v-slot:default="{ isActive }">
                    <v-card title="Error" prepend-icon="mdi-alert-circle-outline">
                        <v-divider class="mt-3"></v-divider>
                        <v-card-text>
                            你还没填写任何内容呢，请先填写内容再提交！
                        </v-card-text>

                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="#EF9A9A" text="关闭" @click="isActive.value = false"></v-btn>
                        </v-card-actions>
                    </v-card>
                </template>
            </v-dialog>
        </div>

        <div>
            <v-snackbar v-model="store.showMessageError" variant="text">
                <v-alert icon="mdi-close" :text="store.showMessageContent" type="error"></v-alert>
            </v-snackbar>
        </div>

        <div>
            <v-dialog v-model="store.showDialogLoading" max-width="320" persistent>
                <v-list class="py-2" color="primary" elevation="12" rounded="lg">
                    <v-list-item prepend-icon="$vuetify-outline" title="正在创建中, 请稍后...">
                        <template v-slot:prepend>
                            <div class="pe-4">
                                <v-icon color="primary" size="x-large"></v-icon>
                            </div>
                        </template>

                        <template v-slot:append>
                            <v-progress-circular color="primary" indeterminate="disable-shrink" size="16"
                                width="2"></v-progress-circular>
                        </template>
                    </v-list-item>
                </v-list>
            </v-dialog>
        </div>

    </div>
</template>

<script setup>
import { useStoreBranchWrite } from "@/store/store_branch_write";
import { Base64 } from 'js-base64';
import CryptoJS from 'crypto-js';
import { ApiUserSubmit } from "@/api/api_user"
import { useRouter } from "vue-router";

const router = useRouter();
const store = useStoreBranchWrite();


const options_time = [
    {
        "title": "永久",
        "value": -1,
    },
    {
        "title": "1分钟",
        "value": 60
    },
    {
        "title": "10分钟",
        "value": 600,
    },
    {
        "title": "30分钟",
        "value": 1800,
    },
    {
        "title": "1小时",
        "value": 3600,
    },
    {
        "title": "12小时",
        "value": 43200,
    },
    {
        "title": "1天",
        "value": 86400,
    },
    {
        "title": "3天",
        "value": 259200,
    },
    {
        "title": "7天",
        "value": 604800,
    },
    {
        "title": "30天",
        "value": 2592000,
    }
]


// 表单提交
const submitForm = () => {
    // 内容不能为空
    if (store.form.context.replace(/\s+/g, "").length === 0) {
        store.showDialog = true;
        return;
    }

    // 显示加载中弹窗
    store.showDialogLoading = true;

    const jsonData = {
        type: store.form.type, // 文本类型
        once: store.form.once, // 只允许访问一次
        exceed_time: store.form.exceed_time, // 超过时间限制
        password: "", // 访问密码
        ips: [], // 限制访问IP(白名单)
        context: "", // 内容
        allow_comment: store.form.allow_comment // 是否允许评论
    }


    // 处理白名单IP
    const ips = store.form.ips_input.split("\n");
    for (let i = 0; i < ips.length; i++) {
        if (ips[i].replace(/\s+/g, "").length !== 0) {
            jsonData.ips.push(ips[i]);
        }
    }

    // 加密context上下文，做个base64编码，防止特殊字符影响
    jsonData.context = Base64.encode(store.form.context);


    // 设置访问密码的hash值（如果不为空）
    if (store.form.password.replace(/\s+/g, "").length !== 0) {
        jsonData.password = CryptoJS.MD5(store.form.password).toString();
    }

    // 延迟2秒再发送请求，防止太频繁
    setTimeout(() => {
        // 发送请求新增条目
        ApiUserSubmit(jsonData)
            .then(res => {
                console.log("success:", res);
                router.push(`/view/read/${res.data.data}`);
            })
            .catch(err => {
                store.showMessageContent = err?.response?.data?.msg || err.message;
                store.showMessageError = true;
                console.log("error:", err);
            })
            .finally(() => {
                store.showDialogLoading = false;
            });
    }, 2000);

}

</script>

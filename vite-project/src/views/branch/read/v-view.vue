<template>

    <div v-if="store.loader_content">
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
    </div>

    <div v-else>
        <v-banner color="#26A69A" icon="mdi-web" lines="one">
            <v-banner-text>
                访问地址
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyText(store.respond_comments.url.web)">复制</v-btn>
            </template>
        </v-banner>

        <v-banner color="error" icon="mdi-bash" lines="one">
            <v-banner-text>
                接口地址
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyText(store.respond_comments.url.api)">复制</v-btn>
            </template>
        </v-banner>

        <v-banner color="#388E3C" icon="mdi-content-copy" lines="one">
            <v-banner-text>
                直接复制
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyText(store.respond_content)">复制</v-btn>
            </template>
        </v-banner>
    </div>

    <div>
        <v-snackbar v-model="message_success" variant="text">
            <v-alert icon="mdi-check" text="复制成功！" type="success"></v-alert>
        </v-snackbar>

        <v-snackbar v-model="message_error" variant="text">
            <v-alert icon="mdi-check" text="复制失败" type="error"></v-alert>
        </v-snackbar>
    </div>

</template>

<script setup>
import { ref } from 'vue'
import { useStoreBranchRead } from "@/store/store_branch_read"
import useClipboard from "vue-clipboard3";


const { toClipboard } = useClipboard();
const store = useStoreBranchRead()
const message_success = ref(false)
const message_error = ref(false);

const copyText = async (text) => {
    if (navigator?.clipboard?.writeText) {
        try {
            await navigator.clipboard.writeText(text)
            message_success.value = true;
        } catch (error) {
            console.error("复制失败", error);
            message_error = true;
        }
    } else {
        try {
            toClipboard(text);
            message_success.value = true;
        } catch (error) {
            console.error("复制失败", error);
            message_error = true;
        }
    }
}


</script>

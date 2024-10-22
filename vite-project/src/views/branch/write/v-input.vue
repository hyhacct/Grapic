<template>
    <div style="padding: 20px;">

        <v-alert color="#2A3B4D" density="compact" icon="mdi-alert" theme="dark">
            1. 请勿输入敏感信息, 比如你的手机号, 身份证, 以及密码等, 如果一定需要, 建议开启访问密码。
            <br>
            2. 请勿发布违反相关规定或各种违法犯罪行为的言论, 否则引起的任何后果均由发布者承担。
        </v-alert>

        <!-- <v-alert type="info" variant="tonal">
            1. 请勿输入敏感信息, 比如你的手机号, 身份证, 以及密码等, 如果一定需要, 建议开启访问密码。
            <br>
            2. 请勿发布违反相关规定或各种违法犯罪行为的言论, 否则引起的任何后果均由发布者承担。
        </v-alert> -->
        <br>

        <v-textarea v-model="store.form.context" auto-grow label="内容" variant='solo-filled'></v-textarea>
    </div>
</template>

<script setup>
import { useStoreBranchWrite } from "@/store/store_branch_write";


const store = useStoreBranchWrite();
</script>
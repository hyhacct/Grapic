import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],


    server: {
        // port: 33505, // 端口设置为33505
        // open: true, // 自动打开浏览器
        cors: true, // 允许跨域请求
    },

    // 基础路径，用于部署时的路径
    base: '/r',

    // 静态资源目录
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src')
        }
    }
});

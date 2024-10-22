import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'; // 确保你的项目能够处理 CSS 文件


// Highlight.js
import VueHighlightJS from '@highlightjs/vue-plugin';
import 'highlight.js/lib/common';
import 'highlight.js/styles/dark.css';


const app = createApp(App)
const pinia = createPinia()
const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'dark'
    },
    icons: {
        defaultSet: 'mdi', // 这是默认值，仅作展示用途
    },
})


app.use(pinia)
app.use(router)
app.use(vuetify)
app.use(VueHighlightJS); // 注册 Highlight.js 插件


app.mount('#app')

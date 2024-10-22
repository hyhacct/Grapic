import { createRouter, createWebHistory } from 'vue-router'


// 文档路由，只有两个页面，只读和编写推送
import ViewBranchWrite from "@/views/branch/write/main.vue"
import ViewBranchRead from "@/views/branch/read/main.vue"


const routes = [
    {
        path: '/',
        redirect: '/view',
        children: [
            {
                path: '/view',
                name: 'view',
                redirect: '/view/write',
                children: [
                    {
                        path: 'write',
                        name: 'branch_write',
                        component: ViewBranchWrite,
                        props: true
                    },
                    {
                        path: 'read/:guid',
                        name: 'branch_read',
                        component: ViewBranchRead,
                        props: true
                    }
                ]
            },
        ]
    },
]


const router = createRouter({
    history: createWebHistory(),
    routes,
})


// router.beforeEach((to, from, next) => {
//     loadingBar.start();
//     next(); // 继续导航
// });

// router.afterEach(() => {
//     loadingBar.finish(); // 导航完成后停止loading
// });



export default router
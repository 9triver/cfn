import {createRouter, createWebHashHistory} from "vue-router";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "home",
            component: () => import("/@/views/home.vue"),
            meta: {
                title: "home",
            }
        }
    ],
});

router.beforeEach((to, from, next) => {
    if (to.path.startsWith('/api')) {
        return;
    }
    next();
});

export default router;

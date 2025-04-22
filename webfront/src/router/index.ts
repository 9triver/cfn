import {createRouter, createWebHashHistory} from "vue-router";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      component: () => import("@/views/HomeView.vue"),
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

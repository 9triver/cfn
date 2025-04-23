const homeRoutes = [
  {
    name: "overview",
    path: "/overview",
    component: () => import("@/views/OverView.vue"),
    meta: {title: "overview",}
  }
  // 其他子路由...
]

export default homeRoutes
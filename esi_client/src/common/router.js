import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    { path: "/", component: () => import("../views/Login.vue") },
    { 
        path: "/dashboard", component: () => import("../views/dashboard/Dashboard.vue"), children: [
            {path: "/dashboard/schrank", component: () => import("../views/dashboard/Schrank.vue")},
            {path: "/dashboard/subrank", component: () => import("../views/dashboard/Subrank.vue")},
            {path: "/dashboard/school", component: () => import("../views/dashboard/School.vue")},
            {path: "/dashboard/subject", component: () => import("../views/dashboard/Subject.vue")},
            {path: "/dashboard/paper", component: () => import("../views/dashboard/Paper.vue")},
            {path: "/dashboard/user", component: () => import("../views/dashboard/User.vue")},
        ]     
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      component: () => import("./views/Home.vue"),
      meta: {
        title: "Home",
      }
    },
    {
      name: "login",
      path: "/login",
      component: () => import("./views/Login.vue"),
      meta: {
        title: "Login",
      }
    },
    {
      name: "room",
      path: "/room/:name",
      component: () => import("./views/Chatroom.vue"),
      meta: {
        title: "Room",
      }
    },
  ],
});

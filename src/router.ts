import { createRouter, createWebHistory } from "vue-router";
import { isloggedin } from "./auth";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      component: () => import("./views/Home.vue"),
      meta: {
        title: "Home",
        auth: true,
      }
    },
    {
      name: "auth",
      path: "/auth",
      component: () => import("./views/Auth.vue"),
      meta: {
        title: "Auth",
      }
    },
    {
      name: "room",
      path: "/room/:name",
      component: () => import("./views/Chatroom.vue"),
      meta: {
        title: "Room",
        auth: true,
      }
    },
  ],
});


router.beforeEach(async (to, from) => {
  if (to.meta.auth && !isloggedin()) {
    // TODO: redirect to login with redirect back to current route
    await router.push({ name: 'auth' });
    return;
  }
});
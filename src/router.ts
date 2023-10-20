import { createRouter, createWebHistory } from "vue-router";
import { isloggedin } from "./api";

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
      name: "new-room",
      path: "/new-room",
      component: () => import("./views/NewRoom.vue"),
      meta: {
        title: "New Room",
        auth: true,
      }
    },
    {
      name: "room",
      path: "/room/:id",
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
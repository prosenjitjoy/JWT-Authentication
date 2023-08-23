import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import RegisterView from "../views/RegisterView.vue";
import LoginView from "../views/LoginView.vue";

import { useStore } from "../stores";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/signin",
      name: "login",
      component: LoginView,
    },
    {
      path: "/signup",
      name: "register",
      component: RegisterView,
    },
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
  ],
});

router.beforeEach(async (to, from) => {
  const auth = useStore();
  if (!auth.isAuthenticated && to.name !== "login" && to.name !== "register") {
    return { name: "login" };
  }
});

export default router;

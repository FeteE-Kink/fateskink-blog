import { createRouter, createWebHistory } from "vue-router";
import AuthLayout from "@/layouts/AuthLayout.vue";
import Default from "@/layouts/Default.vue";
import SignIn from "@/pages/auths/SignIn.vue";
import Home from "@/pages/Home.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Home,
      meta: {
        layout: Default,
        publics: true,
      },
    },
    {
      path: "/signin",
      component: SignIn,
      meta: {
        layout: AuthLayout,
        publics: true,
      },
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  // const requiresAuth = !to.meta.publics;

  // const authStore = useAuthStore();

  // const isAuthenticated = authStore.token;

  // if (requiresAuth && !isAuthenticated) {
  //   authStore.token = null;
  //   next("/signIn");
  // } else
  next();
});

export default router;

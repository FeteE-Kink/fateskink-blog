import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { useLocalStorage } from "@vueuse/core";

import { useRouter } from "vue-router";

import AuthRepository from "@/apis/repositories/authRepository";

export const useAuthStore = defineStore("auths", () => {
  const router = useRouter();

  // ================ STATES ==========================
  const token = ref(useLocalStorage("authToken", null));

  // ================ GETTERS ========================
  const layout = computed(() => !!token.value || "Default");

  // ================ ACTIONS ========================
  function setToken(tokenValue) {
    token.value = tokenValue;
  }

  async function signInAction({ email, password }) {
    const data = await AuthRepository.signIn({ email, password });
    token.value = data.signIn?.token;

    router.push("/");
  }

  return {
    // data
    token,
    layout,

    // function
    setToken,
    signInAction,
  };
});

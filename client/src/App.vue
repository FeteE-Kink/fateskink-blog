<template>
  <component :is="layout">
    <router-view></router-view>
  </component>
</template>

<script>
import { defineAsyncComponent, computed, onMounted } from "vue";
import { useRoute } from "vue-router";

// import HomePage from "./components/HomePage.vue";

export default {
  name: "App",
  components: {
    // HomePage,
    AuthLayout: defineAsyncComponent(() => import("@/layouts/AuthLayout.vue")),
    Default: defineAsyncComponent(() => import("@/layouts/Default.vue")),
  },

  setup() {
    const route = useRoute();

    console.log("meta", route.meta);
    const layout = computed(() => route.meta.layout);

    return { layout };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.centered-image {
  display: block;
  margin: 0 auto;
}
</style>

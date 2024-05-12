<template>
  <img
    alt="Fateskink logo"
    src="./assets/logo.png"
    class="centered-image"
    style="width: 150px; height: 150px"
  />
  <HomePage msg="Welcome to Your Vue.js App" />
</template>

<script>
import HomePage from "./components/HomePage.vue";

export default {
  name: "App",
  components: {
    HomePage,
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

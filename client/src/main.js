import { createApp, h } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import axios from "@/apis/axios";

import "./style.css";

const pinia = createPinia();

const app = createApp({
  render: () => h(App),
});

app.use(router);
app.use(pinia);
app.provide("api", axios);

app.mount("#app");

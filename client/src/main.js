import { createApp, h } from "vue";
import App from "./App.vue";
import "./style.css";
import router from "./router";

const app = createApp({
  render: () => h(App),
});

app.use(router);

app.mount("#app");

import { createApp } from "vue";
import App from "./App.vue";
import axios from "axios";

import VueGoogleMaps from "@fawmi/vue-google-maps";

const app = createApp(App);

app.config.globalProperties.$axios = axios.create({
  baseURL: "localhost:3003/distance",
});

app.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyD2OrZQyMj5LZe_69ZeNQq0uKIp2UQpO-w",
    libraries: "places",
  },
});

app.mount("#app");

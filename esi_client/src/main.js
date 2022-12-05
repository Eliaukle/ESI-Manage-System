import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import naive from "naive-ui"; // 引入ui框架
import { createDiscreteApi } from "naive-ui"; // 引入createDiscreteApi
import { createPinia } from "pinia"; // 引入pinia
import { router } from "./common/router"; // 引入路由
import axios from "axios"; // 引入axios
import { UserStore } from "./stores/UserStore" // 引入UserStore
import fileDownload from 'js-file-download';    // 引入js-file-download

axios.defaults.baseURL = "http://127.0.0.1:8080"; // 服务端地址全局配置
const { message, notification, dialog } = createDiscreteApi(["message", "notification", "dialog"])

const app = createApp(App);

app.provide("axios", axios); // 将axios全局放入
app.provide("message", message)
app.provide("notification", notification)
app.provide("dialog", dialog)
app.provide("serverUrl", axios.defaults.baseURL)
app.provide("fileDownload", fileDownload)

app.use(naive); // 引入ui框架
app.use(createPinia()); // 引入pinia

const userStore = UserStore()
// 请求拦截器
axios.interceptors.request.use((config) => {
    config.headers.authorization = `Bearer ${userStore.token}`
    return config
})

app.use(router); // 引入路由
app.mount("#app");

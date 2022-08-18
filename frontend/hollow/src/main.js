import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import pinia from './store/store'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { globalStore } from './store/global'
import { dialogStore } from './store/dialog'
import { snackbarStore } from './store/snackbar'
import { httpStore } from './store/http'

loadFonts()

const app = createApp(App)
const global = globalStore()
const dialog = dialogStore()
const snackbar = snackbarStore()
const http = httpStore()

axios.defaults.withCredentials = true
axios.defaults.maxRedirects = 0
axios.defaults.headers['Content-Type'] = 'application/json'
axios.defaults.headers['Accept'] = '*/*'

app.use(VueAxios,axios)
app.use(vuetify)
app.use(pinia)
app.use(router)

app.config.globalProperties.global = global;        //全局变量
app.config.globalProperties.dialog = dialog;        //对话框
app.config.globalProperties.snackbar = snackbar;    //底部短消息
app.config.globalProperties.http = http;            //HTTP请求封装

app.mount('#app')

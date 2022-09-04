import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import { useVuelidate } from '@vuelidate/core'
import pinia from './store/store'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { globalStore } from './store/global'
import { dialogStore } from './store/dialog'
import { snackbarStore } from './store/snackbar'
import { httpStore } from './store/http'
import { utilsStore } from './store/utils'

loadFonts()

const app = createApp(App)
const global = globalStore()
const dialog = dialogStore()
const snackbar = snackbarStore()
const http = httpStore()
const utils = utilsStore()


// Axios全局设置
axios.defaults.withCredentials = true
axios.defaults.maxRedirects = 0
axios.defaults.headers['Content-Type'] = 'application/json'
axios.defaults.headers['Accept'] = '*/*'

app.use(VueAxios,axios)
app.use(vuetify)
app.use(pinia)
// app.use(Vuelidate)
app.use(router)

app.config.globalProperties.$v = useVuelidate()     // 启用vuelidate
app.config.globalProperties.global = global;        // 全局变量
app.config.globalProperties.dialog = dialog;        // 对话框
app.config.globalProperties.snackbar = snackbar;    // 底部短消息
app.config.globalProperties.HTTP = http;            // HTTP请求封装
app.config.globalProperties.utils = utils;          // 工具函数类

// 全局函数
app.config.globalProperties.goback = () => {
    if(window.history.length >1){
        router.go(-1)
    }else{
        router.push({ path: '/' })
    }
}
app.config.globalProperties.scollTop = () => {
    const upRoll = setInterval(() => {
        const top = document.documentElement.scrollTop // 每次获取页面被卷去的部分
        const speed = Math.ceil(top / 10) // 每次滚动多少 （步长值）
        if (document.documentElement.scrollTop > 0) {
          document.documentElement.scrollTop -= speed // 不在顶部 每次滚动到的位置
        } else {
          clearInterval(upRoll) // 回到顶部清除定时器
        }
    }, 20)
}

app.mount('#app')

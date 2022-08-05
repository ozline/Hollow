import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import pinia from './store/store'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'

loadFonts()

const app = createApp(App)

axios.defaults.withCredentials = true
axios.defaults.maxRedirects = 0
axios.defaults.headers['Content-Type'] = 'application/json'
axios.defaults.headers['Accept'] = '*/*'

app.use(VueAxios,axios)
app.use(vuetify)
app.use(pinia)
app.use(router)

app.mount('#app')

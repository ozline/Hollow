import { createWebHistory,createRouter } from 'vue-router'
import Main from '../views/main.vue'
import Login from '../views/user/login.vue'
import Register from '../views/user/register.vue'
import Detail from '../views/user/detail.vue'
import NotFound from '../views/404.vue'
import pinia from '../store/store'
import { globalStore } from '../store/global';


const routes=[
    { path:'/', component:Main, meta:{ title: '首页', auth: false }},
    { path:'/user/login', component:Login, meta:{ title: '账号登录', auth: false }},
    { path:'/user/register', component:Register, meta:{ title: '账号注册', auth: false }},
    { path:'/user/detail', component:Detail, meta:{ title: '账号信息', auth: true }},
    { path:'/404', component:NotFound, meta:{ title: '404', auth: false }},
    { path:'/:pathMatch(.*)', redirect:'/404'},
]

const router = createRouter({
    base: '/',
    history: createWebHistory(),
    mode: 'history',
    routes
})

const storeGlobal = globalStore(pinia);

router.beforeEach((to,from,next)=>{
    // console.log('组件跳转 从 '+from.meta.title+" 至 "+to.meta.title)
    const authCheck = to.meta.auth;
    if(authCheck == true){
        if(storeGlobal.isLogin == false){
            console.log("Auth Failed")
            next('/user/login');
            return;
        }
    }
    if(to.meta.title){
        document.title = to.meta.title
    }
    next()
})

export default router
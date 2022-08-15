import { createWebHistory,createRouter } from 'vue-router'
import Main from '../views/main.vue'
import Login from '../views/user/login.vue'
import Register from '../views/user/register.vue'
import Detail from '../views/user/space.vue'
import NotFound from '../views/404.vue'
import Upload from '../views/forest/upload.vue'
import pinia from '../store/store'
import { globalStore } from '../store/global';


/*
template:
    {
        path: [路径],
        component: [组件],
        meta:{ title: [标题], auth: [鉴权开关] }
    },
*/

const routes=[
    {
        path:'/',
        component:Main,
        meta:{ title: '首页', auth: false }
    },
    {
        path:'/user/login',
        component:Login,
        meta:{ title: '账号登录', auth: false }
    },
    {
        path:'/user/register',
        component:Register,
        meta:{ title: '注册账号', auth: false }
    },
    {
        path:'/user/detail',
        component:Detail,
        meta:{ title: '账号信息', auth: true }
    },
    {
        path:'/forest/upload',
        component:Upload,
        meta:{ title: '发表想法', auth: true }
    },
    {
        path:'/404',
        component:NotFound,
        meta:{ title: '404', auth: false }
    },
    {
        path:'/:pathMatch(.*)',
        redirect:'/404'
    },
]

const router = createRouter({
    base: '/',
    history: createWebHistory(),
    mode: 'history',
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    }
})

const storeGlobal = globalStore(pinia);

router.beforeEach((to,from,next)=>{
    // console.log('组件跳转 从 '+from.meta.title+" 至 "+to.meta.title)
    const authCheck = to.meta.auth;

    //鉴权验证
    if(authCheck == true){
        if(storeGlobal.isLogin == false){
            console.log("Auth Failed")
            next('/user/login');
            return;
        }
    }
    if(to.meta.title){
        document.title = 'Hollow - ' + to.meta.title
    }
    next()
})

export default router
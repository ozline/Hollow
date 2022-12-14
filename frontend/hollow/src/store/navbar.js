/**
 * Navbar 导航条与侧边栏变化
 */
import { defineStore } from "pinia";

export const navbarStore = defineStore('navbar', {
    state: () => {
        return {
            title: "Hollow 抒发树洞",
            items: [
                {
                    title: '首页',
                    value: '/',
                    props: {
                        prependIcon: 'mdi-home',
                    },
                    show: true,
                },
                {
                    title: '我的',
                    value: 'myAccount',
                    props: {
                        prependIcon: 'mdi-account',
                    },
                    show: false,
                },
                {
                    title: '投诉',
                    value: '/user/report',
                    props: {
                        prependIcon: 'mdi-pencil',
                    }
                },
                {
                    title: '设置',
                    value: '/user/settings',
                    props: {
                        prependIcon: 'mdi-wrench',
                    }
                },
            ]
        }
    },
})
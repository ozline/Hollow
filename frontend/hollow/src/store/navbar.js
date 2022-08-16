import { defineStore } from "pinia";

export const navbarStore = defineStore('navbarStore', {
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
                    value: '/user/detail',
                    props: {
                        prependIcon: 'mdi-account',
                    },
                    show: false,
                },
                {
                    title: '设置',
                    value: '/settings',
                    props: {
                        prependIcon: 'mdi-account',
                    }
                },
            ]
        }
    },
})
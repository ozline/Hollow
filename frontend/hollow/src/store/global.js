/**
 * Global 全局变量与全局函数
 */
import { defineStore } from "pinia";

export const globalStore = defineStore('global', {
    state: () => {
        return {
            isLogin: false,
            user: [],
            token: '',
            anonymous: ['Alice','Bob','Carol','Dave','Eve','Francis','Grace','Hans','Isabella','Jason','Kate','Louis','Margaret','Olivia','Paul','Queen','Richard','Susan','Thomas','Uma','Vivian','Winner','Xander','Yasmine','Zach']
        }
    },

    actions: {
        logUser(){
            console.log(this.user)
        },
        setUser(isLogin,user,token){
            this.isLogin = isLogin
            this.user = user
            this.token = token
        },
        logout(){
            this.isLogin = false
            this.user = []
            this.token = ''
        }
    }
})
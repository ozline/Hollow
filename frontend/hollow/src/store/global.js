import { defineStore } from "pinia";

export const globalStore = defineStore('globalStore', {
    state: () => {
        return {
            isLogin: false,
            user: [],
            token: '',
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
        },
        isEmptyStr(str){
            return (str == undefined || str == null || str == '')
        },
    }
})
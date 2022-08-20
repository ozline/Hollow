/**
 * Snackbar 底部消息框提示
 */

import { defineStore } from "pinia";

export const snackbarStore = defineStore('snackbar', {
    state: () => {
        return {
            data:{
                show: false,
                content: "null",
                color: "blue",
                timeout: "2000",
            }
        }
    },

    actions: {
        close(){
            this.data.content = "null"
            this.data.show = false
        },
        show(content,timeout){
            if(typeof timeout == "undefined"){
                this.data.timeout = "2000"
            }
            this.data.content = content
            this.data.timeout = timeout
            this.data.show = true
        },
        update(data){
            this.data = data
        }
    }
})
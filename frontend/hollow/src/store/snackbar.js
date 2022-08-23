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
                btnshow: false,
            }
        }
    },

    actions: {
        close(){
            this.data.content = "null"
            this.data.show = false
        },
        show(content,timeout){
            this.data.timeout = (timeout == undefined) ? 1000 : timeout,
            this.data.content = content
            this.data.show = true
        },
        update(data){
            this.data = data
        }
    }
})
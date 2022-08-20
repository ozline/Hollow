/**
 * Dialog 对话框封装
 */

import { defineStore } from "pinia";

export const dialogStore = defineStore('dialog', {
    state: () => {
        return {
            data:{
                show: false,
                title: "null",
                content: "null",
            }
        }
    },

    actions: {
        close(){
            this.data.title = "null"
            this.data.content = "null"
            this.data.show = false
        },
        show(title,content){
            this.data.title = title
            this.data.content = content
            this.data.show = true
        },
        update(data){
            this.data = data
        },
        handleError(error){
            var status = error.status
            if(status == 500){
                this.show("出现错误","服务器请求失败,请检查网络连接")
                return
            }
            var reason = (error.data.reason == undefined ? 'NULL' : error.data.reason)
            var code = (error.data.code == undefined ? 'NULL' : error.data.code)
            var message = (error.data.message == undefined ? 'NULL' : error.data.message)
            this.show("出现错误","代码:"+code+"\n原因:"+reason+"\n信息:"+message)
        }
    }
})

export default global
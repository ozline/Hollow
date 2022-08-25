/**
 * HTTP请求封装，在Axios的基础上二度封装请求
 */

import axios from 'axios'
import { defineStore } from "pinia";
import { dialogStore } from "./dialog";
import { globalStore } from "./global";

const dialog = dialogStore()
const global = globalStore()

export const httpStore = defineStore('http', {
    state: () => {
        return {
            prefix: "/apis"
        }
    },
    actions: {
        // 错误处理
        handleError(error){
            var status = error.status
            if(status == 500) dialog.show("出现错误","服务器请求失败,请检查网络连接(500)")
            else if(status == 404) dialog.show("出现错误","服务器异常,请等待管理员检查(404)")
            else{
                var reason = (error.data.reason == undefined ? 'NULL' : error.data.reason)
                var code = (error.data.code == undefined ? 'NULL' : error.data.code)
                var message = (error.data.message == undefined ? 'NULL' : error.data.message)
                dialog.show("出现错误","["+code+" "+reason+"] "+message)
            }
        },
        // 统一请求结构
        unifiedQuery(method,url,auth,struct){
            struct.method = method                      // 增加方法
            struct.url = this.prefix + url              // 增加路径
            if(auth == true){
                struct.headers = {
                    'Authorization': global.token       // 增加鉴权
                }
            }
            return new Promise((resolve) => {
                axios(struct)
                .then(res => {
                    var result = JSON.parse(JSON.stringify(res.data))
                    if(result.code == 200){
                        if(auth == true){
                            global.updateToken(res.headers.authorization)
                        }
                        resolve(result)
                    }else{
                        this.handleError(result)
                    }
                })
                .catch(error => {
                    // console.log("status",error.response.status)
                    if(error.response.status == 417){
                        resolve("NEEDMFA")
                    }else{
                        this.handleError(error.response)
                    }
                })
            })
        },
        // GET
        get(url,params,auth){
            return this.unifiedQuery("get",url,auth,{
                params: params
            })
        },
        // POST
        post(url,body,auth){
            return this.unifiedQuery("post",url,auth,{
                data: JSON.stringify(body)
            })
        },
        // DELETE
        delete(url,params,auth){
            return this.unifiedQuery("delete",url,auth,{
                params: params
            })
        },
        // PUT
        put(url,body,auth){
            return this.unifiedQuery("put",url,auth,{
                data: JSON.stringify(body)
            })
        },
    }
})
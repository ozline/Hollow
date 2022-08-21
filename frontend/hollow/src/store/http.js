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
            prefix: "/apis/"
        }
    },
    actions: {
        // 错误处理
        handleError(error){
            console.log(error)
            var status = error.status
            if(status == 500) dialog.show("出现错误","服务器请求失败,请检查网络连接(500)")
            else if(status == 404) dialog.show("出现错误","服务器异常,请等待管理员检查(404)")
            else{
                var reason = (error.data.reason == undefined ? 'NULL' : error.data.reason)
                var code = (error.data.code == undefined ? 'NULL' : error.data.code)
                var message = (error.data.message == undefined ? 'NULL' : error.data.message)
                dialog.show("出现错误","代码:"+code+"\n原因:"+reason+"\n信息:"+message)
            }
        },
        // 数据处理
        handleResponse(response){
            var result = JSON.parse(JSON.stringify(response.data))
            if(result.code == 200){
                return {
                    success: true,
                    data: result,
                }
            }
            this.handleError(result)
            return {
                success: false
            }

        },
        // 生成请求头
        generateHeaders(auth){
            if(auth == true){
                return {
                    'Authorization': global.token
                }
            }
            return { }
        },
        // GET
        get(url,params,auth){
            return new Promise((resolve) => {
                axios({
                    url: this.prefix + url,
                    method: "get",
                    params: params,
                    headers: this.generateHeaders(auth)
                })
                .then(res => {
                    var result = this.handleResponse(res)
                    if(result.success == true) resolve(result.data)
                })
                .catch(error => {
                    this.handleError(error.response)
                })
            })
        },
        // POST
        post(url,params,auth){
            return new Promise((resolve) => {
                var data = JSON.stringify(params)
                axios({
                    url: this.prefix + url,
                    method: "post",
                    data: data,
                    headers: this.generateHeaders(auth)
                })
                .then(res => {
                    var result = this.handleResponse(res)
                    if(result.success == true) resolve(result.data)
                })
                .catch(error => {
                    this.handleError(error.response)
                })
            })
        }
    }
})
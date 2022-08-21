/**
 * Utils 工具函数类
 */
 import { defineStore } from "pinia";

 export const utilsStore = defineStore('utils', {
     actions: {
         isEmptyStr(str){
            return (str == undefined || str == null || str == '')
         },
         timestampConvert(timestamp){
            let date = new Date(Number(timestamp))
            var Y = date.getFullYear() + '-'
            var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-'
            const D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' '
            const h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':'
            const m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes())
            // const s = (date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds())
            return Y + M + D + h + m
         },
     }
 })
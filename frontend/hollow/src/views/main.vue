<template>

<v-container>
    <h1> 最新树洞 <v-btn @click="upload" color="success" style="float:right;margin-top:6px;"> 发表想法 </v-btn> </h1>

    <div style="height:26px"></div>

    <v-card
        v-for="(item, index) in datas"
        v-bind:key="index"
        :title="'作者:' + item.owner"
        :subtitle=" '发表时间:' + item.createAt"
        :text=" item.message"
        style="margin-bottom: 26px;"
        @click="cardClickEvent(item.id)"
    ></v-card>

    <!-- 通知框 -->
    <Component-Dialog :data="dialog" @update="update"></Component-Dialog>
</v-container>

</template>

<script>
import ComponentDialog from '../components/dialog.vue'
import { globalStore } from '../store/global';

export default {
    name: 'MainPage',

    components:{
        'Component-Dialog' : ComponentDialog,
    },

    setup(){
        const global = globalStore();
        return{
            global,
        }
    },

    created(){
        this.refresh()
    },

    data(){
        return{
            datas: [],
            total: 0,
            pagesize: 10,
            current: 1,

            dialog: {
                show: false,
                title: 'Title',
                content: 'Content',
                confirm: null,
            },
        }
    },

    methods:{
        upload(){
            this.$router.push('/forest/upload');
        },
        update(data) {
            this.dialog = data;
        },
        refresh(){
            var data = JSON.stringify({
                page: this.current,
                pagesize: this.pagesize,
            })

            this.axios.post("/apis/forest/all", data).then(res => {
                var result = JSON.parse(JSON.stringify(res.data))
                if(result.code == 200){
                    console.log(result)
                    this.datas = result.data.list
                    this.total = result.data.total
                }
                else{
                    this.handleError(result)
                }
            }).catch(err => {
                this.handleError(err.response.data)
            })
        },
        handleError(error){
            var reason = (error.reason == undefined ? 'NULL' : error.reason)
            var code = (error.code == undefined ? 'NULL' : error.code)
            var message = (error.message == undefined ? 'NULL' : error.message)
            this.showDialog("出现错误",'代码:'+code+"\n原因:"+reason+"\n信息:"+message)
        },
        showDialog(title, content,confirm){
            this.dialog = {
                show: true,
                title: title,
                content: content,
                confirm: (confirm != undefined && confirm != null) ? confirm : null,
            }
        },
        cardClickEvent(id){
            this.dialog = {
                show: true,
                title: '提示',
                content: 'ID: '+ id,
            }
        }
    }
}
</script>

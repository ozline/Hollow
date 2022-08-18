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

</v-container>

</template>

<script>

export default {
    name: 'MainPage',

    created(){
        this.refresh()
    },

    data(){
        return{
            datas: [],
            total: 0,
            pagesize: 10,
            current: 1,
        }
    },

    methods:{
        upload(){
            this.$router.push('/forest/upload');
        },
        refresh(){
            var data = {
                page: this.current,
                pagesize: this.pagesize,
            }

            this.http.get('/forest/all', data).then(res => {
                this.datas = res.data.list
                this.total = res.data.total
            })
        },
        cardClickEvent(id){
            this.dialog.show("提示", 'ID:' + id)
        }
    }
}
</script>

<template>

<v-container>
    <h1> 最新树洞 <v-btn @click="upload" color="success" style="float:right;margin-top:6px;"> 发表想法 </v-btn> </h1>

    <div style="height:26px"></div>

    <v-card
        v-for="(item, index) in datas"
        v-bind:key="index"
        :title="'# ID - ' + item.id"
        :subtitle="utils.timestampConvert(item.createdAt)"
        :text="item.message"
        style="margin-bottom: 26px;"
        @click="cardClickEvent(item.id)"
    ></v-card>

    <v-pagination
      v-model="page.index"
      :length="page.length"
      :total-visible="5"
    ></v-pagination>

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
            page:{
                index: 1,
                size: 10,
                length: 0,
            }
        }
    },
    watch:{
        page:{
            deep: true,
            handler(){
                this.refresh()
                this.snackbar.show("加载第" + this.page.index + "页成功")
            }
        }
    },
    methods:{
        upload(){
            this.$router.push('/forest/upload');
        },
        refresh(){
            var data = {
                page: this.page.index,
                pagesize: this.page.size,
            }

            this.HTTP.get('/forest/all', data).then(res => {
                this.datas = res.data.list
                this.total = res.data.total
                this.page.length = Math.ceil(this.total/this.page.size)
            })
        },
        cardClickEvent(id){
            this.$router.push("/forest/" + id)
        }
    }
}
</script>

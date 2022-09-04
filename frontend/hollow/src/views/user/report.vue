<template>

<v-container>
    <h1> 我的投诉 </h1>

    <div style="height:26px"></div>

    <v-card
        v-for="(item, index) in datas"
        v-bind:key="index"
        :title="'# ID - ' + item.id"
        :subtitle="'最后更新于:' + utils.timestampConvert(item.updatedAt)"
        style="margin-bottom: 26px;"
        variant="outlined"
    >
    <v-card-text class="text--primary">
        <v-card
            class="mx-auto"
            tile
        >
            <v-list-item two-line>
                <v-list-item-title>原文</v-list-item-title>
                <v-list-item-subtitle>{{ item.message }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item two-line>
                <v-list-item-title>投诉原因</v-list-item-title>
                <v-list-item-subtitle>{{ item.reason }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item two-line>
                <v-list-item-title>当前状态</v-list-item-title>
                <v-list-item-subtitle>{{ convertStatus(item.status) }}</v-list-item-subtitle>
            </v-list-item>

            <v-list-item two-line>
                <v-list-item-title>反馈</v-list-item-title>
                <v-list-item-subtitle>{{ item.reply }}</v-list-item-subtitle>
            </v-list-item>

        </v-card>
    </v-card-text>
    </v-card>

    <v-pagination
        v-model="page.index"
        :length="page.length"
        :total-visible="5"
    ></v-pagination>

</v-container>

</template>

<script>
export default {
    name: 'ReportPage',

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
        refresh(){
            var data = {
                page: this.page.index,
                pagesize: this.page.size,
            }

            this.HTTP.get('/report', data,true).then(res => {
                console.log(res.data)
                this.datas = res.data.list
                this.total = res.data.total
                this.page.length = Math.ceil(this.total/this.page.size)
            })
            this.scollTop()
        },
        convertStatus(status){
            return (status == "0") ? "待处理" : "处理结束"
        }
    }
}
</script>

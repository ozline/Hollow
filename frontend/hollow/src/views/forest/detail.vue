<template>
    <v-container>
        <div style="margin-bottom: 26px;">
            <h1> 详情
            <v-btn color="primary" style="float:right;margin-top:6px" @click="submit" prepend-icon="mdi-cloud-upload">发表评论</v-btn>
            <v-btn color="warning" style="float:right;margin-top:6px;" class="mr-4" @click="goback">
                <v-icon
                start
                icon="mdi-arrow-left"
                ></v-icon>返回
            </v-btn>
            </h1>
        </div>

        <v-card
                :title="'楼主 - ' + id"
                :subtitle="utils.timestampConvert(data.createdAt)"
                :text="data.message"
                v-bind:key="0"
                style="margin-bottom: 26px;"
        ></v-card>

        <v-card
            v-for="(item,index) in comments.datas"
            v-bind:key="item.id"
            :title="'#' + (index + 1) + ' - ' + this.global.anonymous[index % 26]"
            :subtitle="utils.timestampConvert(item.createdAt)"
            :text="item.message"
            v-show="item.status != 2"
            style="margin-bottom: 26px;"
            variant="outlined"
        >
            <v-card-actions>
                <v-btn variant="outlined" color="error" size="x-small">
                    违规投诉
                </v-btn>
                <v-btn variant="outlined" v-show="item.status == 1" size="x-small">
                    查看层主
                </v-btn>
            </v-card-actions>

            <template v-slot:append>
                <v-btn
                    color="error"
                    @click="delComment(item.id)"
                    v-show="item.owner == global.user.userid"
                    icon
                    size="x-small"
                    class="mr-1"
                >
                    <v-icon>mdi-minus-circle</v-icon>
                    <v-tooltip activator="parent" location="bottom">删除评论</v-tooltip>
                </v-btn>
                <v-btn
                    icon
                    color="success"
                    size="x-small"
                    @click="like(item.id,index)"
                >
                    <v-icon>mdi-thumb-up</v-icon>
                    <v-tooltip activator="parent" location="bottom">点赞</v-tooltip>
                </v-btn>
                <v-badge
                color="info"
                :content="'喜欢: ' + item.liked"
                inline
                ></v-badge>
            </template>
        </v-card>
        <v-pagination
            v-model="page.index"
            :length="page.length"
            :total-visible="5"
        ></v-pagination>
    </v-container>
</template>


<script>
import { mdiDelete } from '@mdi/js'
export default {
    name: 'forestDetail',
    data() {
        return {
            data: [],
            id: 0,
            comments:{
                datas: [],
                total: 0,
            },
            page:{
                index: 1,
                size: 10,
            },
            icons: { mdiDelete },
        }
    },
    created(){
        this.refresh()
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
            this.id = this.$route.params.id

            this.HTTP.get("/forest/"+this.id,{ },true).then(res =>{
                this.data = res.data
            })

            var data = {
                page: this.page.index,
                pagesize: this.page.size,
            }

            this.HTTP.get("/forest/comments/"+ this.id,data,true).then(res => {
                this.comments.datas = res.data.list
                this.comments.total = res.data.total
            })
        },
        submit(){
            this.$router.push("/forest/upload/" + this.id)
        },
        like(id,index){
            this.HTTP.put("/forest/comments/" + id,{},true).then(() => {
                this.comments.datas[index].liked = Number(this.comments.datas[index].liked) + 1
                this.snackbar.show("点赞成功")
            })
        },
        report(){
            //举报
        },
        delComment(id){
            this.HTTP.delete("/forest/comments/" + id,true).then(() => {
                this.snackbar.show("删除成功!")
                this.refresh()
            })
        }
    }
}
</script>
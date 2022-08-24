<template>
    <v-container>
        <h1 style="margin-bottom: 26px;"> 详情 - {{ id }}
        <v-btn color="primary" style="float:right;margin-top:6px" @click="submit" prepend-icon="mdi-cloud-upload">发表评论</v-btn>
        <v-btn color="warning" style="float:right;margin-top:6px" class="mr-4" @click="goback">
            <v-icon
            start
            icon="mdi-arrow-left"
            ></v-icon>返回
        </v-btn>
        </h1>
        <v-card
            title="楼主"
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
                <v-btn variant="outlined" color="error">
                    违规投诉
                </v-btn>
                <v-btn variant="outlined" v-show="item.status == 1">
                    查看层主(ID: {{ item.owner }})
                </v-btn>
            </v-card-actions>

            <template v-slot:append>
                <v-btn color="error" @click="delComment(item.id)" v-show="item.owner == global.user.userid" icon="mdi-minus-circle" size="x-small"></v-btn>
                <v-btn
                    icon="mdi-thumb-up"
                    color="success"
                    size="x-small"
                    @click="like"
                ></v-btn>
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
        console.log(this.global.user.userid)
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
        like(){
            //喜欢
        },
        report(){
            //举报
        },
        delComment(id){
            console.log(id)
        }
    }
}
</script>
<template>
<v-container>
    <div style="margin-bottom: 26px;">
        <h1 style="margin-bottom: 26px;"> 账号信息
        <v-btn color="warning" style="float:right;margin-top:6px;" @click="goback">
            <v-icon
            start
            icon="mdi-arrow-left"
            ></v-icon>返回
        </v-btn>
        </h1>
    </div>

    <v-card
        class="mx-auto"
        tile
    >
        <v-list-item two-line>
            <v-list-item-title>用户ID</v-list-item-title>
            <v-list-item-subtitle># {{ id }}</v-list-item-subtitle>
        </v-list-item>

        <v-list-item two-line>
            <v-list-item-title>用户昵称</v-list-item-title>
            <v-list-item-subtitle>{{ data.nickname }}</v-list-item-subtitle>
        </v-list-item>

        <v-list-item two-line>
            <v-list-item-title>手机号</v-list-item-title>
            <v-list-item-subtitle>{{ data.phone }}</v-list-item-subtitle>
        </v-list-item>

        <v-list-item two-line>
            <v-list-item-title>邮箱</v-list-item-title>
            <v-list-item-subtitle>{{ data.email }}</v-list-item-subtitle>
        </v-list-item>

        <v-list-item two-line>
            <v-list-item-title>注册时间</v-list-item-title>
            <v-list-item-subtitle>{{ utils.timestampConvert(data.createdAt) }}</v-list-item-subtitle>
        </v-list-item>

    </v-card>
</v-container>
</template>

<script>

export default {
    name: 'userDetail',
    data(){
        return {
            id: 0,
            data: [],
        }
    },
    created(){
        this.id = this.$route.params.id
        this.refresh()
    },
    methods:{
        refresh(){
            this.HTTP.get("/user/info/" + this.id,{},true).then( res => {
                this.data = res.data //更新用户信息
            })
        }
    },
}

</script>

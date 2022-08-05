<template>
<v-container>
    <h1 style="margin-bottom: 39px;"> 404 NOT FOUND</h1>
    <v-btn
        variant="outlined"
        color="primary"
        @click="back"
    >返回上一页
    </v-btn>
</v-container>
</template>

<script>

export default{
    name: 'NotFound',
    methods:{
        back(){
            if(window.history.length >1){
                this.$router.go(-1)
            }else{
                this.$router.push({
                path: '/'
                })
            }
        },
    }
}

</script>
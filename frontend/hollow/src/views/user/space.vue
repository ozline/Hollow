<template>
<v-container>
    <h1> 账号信息 -  #{{ id }}</h1>

</v-container>
</template>

<script>

export default {
    name: 'userDetail',
    data(){
        return {
            id: 0,
        }
    },
    created(){
        this.id = this.$route.params.id
        this.refresh()
    },
    methods:{
        refresh(){

        }
    },
}

</script>

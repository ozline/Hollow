<template>
    <v-container>
        <h1> 详情 </h1>
    </v-container>
</template>


<script>
export default {
    name: 'forestDetail',
    setup(){
        const storeGlobal = globalStore();
        return{
            storeGlobal,
        }
    },
}
</script>
<template>
<v-container>
    <h1> 账号信息 </h1>

</v-container>
</template>

<script>

import { globalStore } from '../../store/global';

export default {
    name: 'userDetail',
    setup(){
        const storeGlobal = globalStore();
        return{
            storeGlobal,
        }
    },
}
</script>

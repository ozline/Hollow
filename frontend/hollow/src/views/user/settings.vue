<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号设置 </h1>
    <mfa></mfa>
</v-container>
</template>

<script>

import ComponentMFA from '../../components/mfa.vue'

export default{
    name: 'UserSettings',
    components:{
        'mfa' : ComponentMFA
    }
}

</script>
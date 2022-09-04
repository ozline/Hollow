<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号设置 </h1>
    <mfa></mfa>
    <bind-phone></bind-phone>
</v-container>
</template>

<script>

import ComponentMFA from '../../components/mfa.vue'
import ComponentBindPhone from '../../components/bindPhone.vue'

export default{
    name: 'UserSettings',
    components:{
        'mfa' : ComponentMFA,
        'bind-phone' : ComponentBindPhone
    }
}

</script>
<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号登录 </h1>
    <v-form
        ref="form"
        lazy-validation
    >
        <v-text-field
        v-model="username"
        :counter="10"
        :rules="usernameRules"
        label="用户名"
        variant="outlined"
        required
        ></v-text-field>

        <v-text-field
        v-model="password"
        :rules="passwordRules"
        :counter="16"
        label="密码"
        type="password"
        variant="outlined"
        required
        ></v-text-field>

        <v-btn

        color="error"
        class="mr-4"
        @click="reset"
        >
        重置
        </v-btn>

        <v-btn

        color="success"
        class="mr-4"
        @click="login"
        >
        登录账号
        </v-btn>

        <v-btn

        color="secondary"
        @click="register"
        style="float:right;"
        >
        注册账号
        </v-btn>
    </v-form>

    <!-- 通知框 -->
    <Component-Dialog :data="dialog" @update="update"></Component-Dialog>
</v-container>
</template>


<script>
import ComponentDialog from '../../components/dialog.vue'
export default {
    name: "userLogin",

    components:{
        'Component-Dialog' : ComponentDialog,
    },

    data: () => ({
        username: '',
        usernameRules: [
            v => !!v || '请输入用户名',
            v => (v && v.length <= 10) || '用户名长度超限',
        ],
        password: '',
        passwordRules: [
            v => !!v || '请输入密码',
            v => (v && v.length <= 16) || '密码长度超限',
        ],
        dialog: {
            show: false,
            title: 'Title',
            content: 'Content',
        },
    }),

    methods: {
        isEmptyStr(str){
            return (str == undefined || str == null || str == '')
        },
        reset() {
            this.$refs.form.reset()
            this.$refs.form.resetValidation()
        },
        login() {
            if(this.isEmptyStr(this.username) || this.isEmptyStr(this.password)){
                this.$refs.form.validate()
                return
            }
            this.dialog = {
                show: true,
                title: '账号信息',
                content: '用户名：'+this.username+'\n密码:'+this.password,
            }
        },
        register() {
            this.$router.push('/user/register');
        },
        update(data) {
            this.dialog = data;
        },
    },
}
</script>
<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号注册 </h1>
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
        variant="outlined"
        type="password"
        required
        ></v-text-field>

        <v-text-field
        v-model="phone"
        :rules="phoneRules"
        :counter="11"
        label="手机号"
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
        color="secondary"
        @click="register"
        >
        注册账号
        </v-btn>

        <v-btn
        style="float:right;"
        color="success"
        class="mr-4"
        @click="login"
        >
        登录账号
        </v-btn>
    </v-form>

    <!-- 通知框 -->
    <Component-Dialog :data="dialog" @update="update"></Component-Dialog>
</v-container>
</template>


<script>
import ComponentDialog from '../../components/dialog.vue'
export default {
    name: "userRegister",

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
            v => (v && v.length >=6) || '最少需要6位密码',
            v => (v && v.length <= 16) || '密码长度超限',
        ],
        phone: '',
        phoneRules: [
            v => !!v || '请输入手机号',
            v => (v && v.length === 11) || '手机号格式不正确',
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
            this.$router.push('/user/login');
        },
        register() {
            this.dialog.show = true;
        },
        update(data) {
            this.dialog = data;
        },
    },
}
</script>
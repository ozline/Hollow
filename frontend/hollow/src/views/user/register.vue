<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号注册 </h1>
    <v-form
        ref="form"
        v-model="valid"
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


        <v-row>
            <v-col cols="12" md="6">
                <v-text-field
            v-model="phone"
            :rules="phoneRules"
            :counter="11"
            label="手机号"
            variant="outlined"
            required
        ></v-text-field>
            </v-col>
            <v-col cols="12" md="6">
                <v-text-field
                    v-model="code"
                    :counter="6"
                    label="验证码"
                    variant="outlined"
                    required
                ></v-text-field>
            </v-col>
        </v-row>

        <v-text-field
            v-model="email"
            :rules="emailRules"
            label="邮箱"
            variant="outlined"
            required
        ></v-text-field>

        <v-btn
            color="error"
            class="mr-4"
            @click="reset"
        >重置</v-btn>

        <v-btn
            color="primary"
            @click="postShortMsg"
            class="mr-4"
            :loading="shortMsgLoading"
        >
        获取短信验证码</v-btn>

        <v-btn
            color="secondary"
            @click="register"
        >
        注册账号</v-btn>

        <v-btn
            style="float:right;"
            color="success"
            class="mr-4"
            @click="login"
        >登录账号</v-btn>
    </v-form>
</v-container>
</template>


<script>

export default {
    name: "userRegister",

    data: () => ({
        valid: false,
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
            //TODO:再加个密码复杂度呗
        ],

        phone: '',
        phoneRules: [
            v => !!v || '请输入手机号',
            v => (v && v.length === 11) || '手机号格式不正确',
            v => (v && /^1[3456789]\d{9}$/.test(v)) || '手机号格式不正确',
        ],

        email: '',
        emailRules: [
            v => !!v || '请输入邮箱',
            v => /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(v) || '邮箱格式不正确',
        ],

        code: '',

        shortMsgLoading: false,

    }),

    methods: {
        reset() {
            this.$refs.form.reset()
            this.$refs.form.resetValidation()
        },
        login() {
            this.$router.push('/user/login');
        },
        postShortMsg(){
            var data = {
                phone : this.phone
            }
            this.HTTP.post('/user/register/shortmsg',data).then( () => {
                this.snackbar.show("获取短信验证码成功,1分钟后可再次获取")
                this.shortMsgLoading = true
                setTimeout(() => {
                    this.shortMsgLoading = false
                }, 60000);
            })
        },
        register() {
            this.$refs.form.validate()

            if(this.valid == false){
                this.dialog.show('提示','请检查输入的信息是否正确')
                return
            }

            var data = {
                username: this.username,
                password: this.password,
                phone: this.phone,
                email: this.email,
                code: this.code,
            }

            this.HTTP.post('/user/register', data).then( () => {
                this.snackbar.show('注册成功','注册成功,请登录账号')
                this.$router.push('/user/login');
            })
        },
    },
}
</script>
<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号登录 </h1>
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
</v-container>
</template>


<script>

export default {
    name: "userLogin",

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
            v => (v && v.length <= 16) || '密码长度超限',
        ],
    }),

    methods: {
        reset() {
            this.$refs.form.reset()
            this.$refs.form.resetValidation()
        },
        login() {
            this.$refs.form.validate()

            if(this.valid==false){
                this.dialog.show('提示','请检查输入的信息是否正确')
                return
            }

            var data = {
                username: this.username,
                password: this.password,
            }

            this.http.post('/user/login', data).then(res => {
                this.global.setUser(true, res.data,res.token)
                this.snackbar.show("登录成功")
                this.$router.push('/')
            })
        },
        register() {
            this.$router.push('/user/register');
        },
    },
}
</script>
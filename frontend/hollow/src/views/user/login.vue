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

    <v-dialog v-model="mfa" persistent>
        <v-card title="双因素认证" >
            <v-card-text>
                <v-text-field
                        v-model="code"
                        :rules="codeRules"
                        :counter="6"
                        label="6位数字代码"
                        required
                ></v-text-field>
                <small>*若要继续登录,请在手机中打开Google Authenticator(身份验证器),输入6位动态码</small>
            </v-card-text>
            <v-card-actions>
                <v-btn color="blue-darken-1" text @click="mfa = false">取消登录</v-btn>
                <v-spacer></v-spacer>
                <v-btn color="error" text @click="loginMFA">登录</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</v-container>
</template>


<script>

export default {
    name: "userLogin",

    data: () => ({
        valid: false,
        mfa: false,
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
        code: '',
        codeRules: [
            v => !!v || '请输入6位代码',
            v => (v && v.length == 6) || '请输入合适的6位代码',
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

            this.HTTP.post('/user/login', data).then(res => {
                if(res == "NEEDMFA"){
                    this.mfa = true
                }else{
                    this.global.setUser(true, res.data,res.token)
                    this.snackbar.show("登录成功")
                    this.$router.push('/')
                }
            })
        },
        loginMFA(){
            var data = {
                username: this.username,
                password: this.password,
                code: this.code
            }

            this.HTTP.post('/user/login', data).then(res => {
                this.mfa = false
                console.log(res)
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
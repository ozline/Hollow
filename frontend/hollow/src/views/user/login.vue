<template>
<v-container>
    <h1 style="margin-bottom: 26px;"> 账号登录 </h1>
    <v-tabs
        fixed-tabs
        grow
    >
        <v-tab>
            用户名登录
        </v-tab>
        <v-tab>
            手机号登录
        </v-tab>
    </v-tabs>
    <v-form
            ref="form"
            v-model="valid"
            lazy-validation
    >
    <!-- :rules="usernameRules" -->
        <v-text-field
            v-model="username"
            :counter="10"
            :error-messages="usernameErrors"
            @input="$v.username.$touch()"
            @blur="$v.username.$touch()"
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
        <v-card title="多因素认证" >
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
                <v-btn color="blue-darken-1" text @click="cancelMFA">取消登录</v-btn>
                <v-spacer></v-spacer>
                <v-btn color="error" text @click="loginMFA">登录</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</v-container>
</template>


<script>

import { required, maxLength, minLength } from '@vuelidate/validators'

export default {
    name: "userLogin",

    // setup(){
    //     return {
    //        $v: useVuelidate()
    //     }
    // },

    data: () => ({
        valid: false,
        mfa: false,
        username: '',
        // usernameRules: [
        //     v => !!v || '请输入用户名',
        //     v => (v && v.length <= 10) || '用户名长度超限',
        // ],
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

    validations() {
        return {
            username: {
                required,
                minLength: minLength(3),
                maxLength: maxLength(10),
            },
            password: {
                required,
                minLength: minLength(6),
                maxLength: maxLength(16),
            },
            code: {
                // required,
                minLength: minLength(6),
                maxLength: maxLength(6),
            },
        }
    },

    methods: {
        reset() {
            this.$refs.form.reset()
            this.$refs.form.resetValidation()
        },
        usernameErrors() {
            const errors = []
            if (!this.$v.username.$dirty) return errors
            !this.$v.username.required && errors.push('用户名不能为空')
            return errors
        },
        login() {
            // this.$v.$touch()
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
                this.global.setUser(true, res.data,res.token)
                this.snackbar.show("登录成功")
                this.$router.push('/')
            })
        },
        cancelMFA(){
            this.mfa = false
            this.snackbar.show("已取消登录:未输入动态码")
        },
        register() {
            this.$router.push('/user/register');
        },
    },
}
</script>
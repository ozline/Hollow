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
import { globalStore } from '../../store/global';
import { dialogStore } from '../../store/dialog';


export default {
    name: "userLogin",

    setup(){
        const global = globalStore();
        const dialog = dialogStore();
        return{
            global,
            dialog
        }
    },

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
                this.showDialog('提示','请检查输入的信息是否正确')
                return
            }

            var data = JSON.stringify({
                username: this.username,
                password: this.password,
            })

            this.axios.post("/apis/user/login", data).then(res => {
                var result = JSON.parse(JSON.stringify(res.data))
                if(result.code == 200){
                    this.global.setUser(true, result.data,result.token)
                    this.$router.push('/')
                }
                else{
                    this.dialog.handleError(result)
                }
            }).catch(err => {
                this.dialog.handleError(err.response)
            })
        },
        register() {
            this.$router.push('/user/register');
        },
    },
}
</script>
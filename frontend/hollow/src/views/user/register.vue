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

        <v-text-field
        v-model="phone"
        :rules="phoneRules"
        :counter="11"
        label="手机号"
        variant="outlined"
        required
        ></v-text-field>

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

        dialog: {
            show: false,
            title: '',
            content: '',
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
            this.$refs.form.validate()

            if(this.valid==false){
                this.showDialog('提示','请检查输入的信息是否正确')
                return
            }

            var data = JSON.stringify({
                username: this.username,
                password: this.password,
                phone: this.phone,
                email: this.email,
            })

            this.axios.post("/apis/user/register", data).then(res => {
                var result = JSON.parse(JSON.stringify(res.data))
                if(result.code == 200){
                    this.showDialog('注册成功','注册成功,3秒后跳转到登录页面')
                    setTimeout(() => {
                        this.$router.push('/user/login');
                    }, 3000);
                }
                else{
                    this.handleError(result)
                }
            }).catch(err => {
                this.handleError(err.response.data)
            })
        },
        update(data) {
            this.dialog = data;
        },
        showDialog(title, content,confirm){
            if(confirm != undefined && confirm != null){
                this.dialog.confirm = confirm
            }

            this.dialog = {
                show: true,
                title: title,
                content: content,
            }
        },
        handleError(error){
            var reason = (error.reason == undefined ? 'NULL' : error.reason)
            var code = (error.code == undefined ? 'NULL' : error.code)
            var message = (error.message == undefined ? 'NULL' : error.message)
            console.log(error)
            this.showDialog("出现错误",'代码:'+code+"\n原因:"+reason+"\n信息:"+message)
        }
    },
}
</script>
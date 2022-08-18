<template>

<v-container>
    <h1 style="margin-bottom: 26px;"> 发表想法
    <v-btn color="primary" style="float:right;margin-top:6px" @click="submit">发表</v-btn>
    <v-btn color="warning" style="float:right;margin-top:6px" class="mr-4" @click="back">返回</v-btn>
    </h1>
    <v-form
        ref="form"
        v-model="valid"
        lazy-validation
    >
        <v-textarea
            clearable
            counter="140"
            append-inner-icon="mdi-comment"
            label="Text"
            v-model="message"
            :rules="messageRules"
            rows="8"
            required
        ></v-textarea>
    </v-form>

    <v-alert
        variant="outlined"
        type="warning"
        color="purple"
    >
        <template v-slot:title>
            言论规范
        </template>
        产品没开发完成前，随便说！
    </v-alert>
</v-container>

</template>

<!-- 1213123132312312312312321312311323213123123213123212131231323123123123123213123113232131231232131232121312313231231231231232131231132321312312321312321213123132312312312312321312311323213123123213123212131231323123123123123213123113232131231232131232 -->

<script>

export default {
    name: 'forestUpload',

    data: () => ({
        valid: false,
        message: '',
        messageRules: [
            v => !!v || '请输入内容',
            v => (v && v.length <= 140) || '内容长度不能超过140个字符',
        ],
    }),

    methods:{
        back(){
            if(window.history.length >1){
                this.$router.go(-1)
            }else{
                this.$router.push({
                path: '/'
                })
            }
        },
        submit(){
            this.$refs.form.validate()

            if(this.valid==false){
                this.dialog.show('提示','请检查输入的信息是否正确')
                return
            }

            var data = JSON.stringify({
                userid: Number(this.global.user.userid),
                message: this.message,
            })

            this.axios.post("/apis/forest/push", data).then(res => {
                var result = JSON.parse(JSON.stringify(res.data))
                if(result.code == 200){
                    this.dialog.show('提示','发送成功!!3秒后跳转首页')
                    setTimeout(() => {
                        this.$router.push('/');
                    }, 3000);
                }
                else{
                    this.dialog.handleError(result)
                }
            }).catch(err => {
                this.dialog.handleError(err.response)
            })
        }
    }
}
</script>
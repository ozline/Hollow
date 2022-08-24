<template>

<v-container>
    <h1 style="margin-bottom: 26px;"> {{ title }}
    <v-btn color="primary" style="float:right;margin-top:6px" @click="submit" prepend-icon="mdi-cloud-upload">发表</v-btn>
    <v-btn color="warning" style="float:right;margin-top:6px" class="mr-4" @click="goback">
            <v-icon
            start
            icon="mdi-arrow-left"
            ></v-icon>返回
        </v-btn>
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
        <v-switch
            v-model="isReal"
            label="不匿名"
            color="primary"
            value="true"
        ></v-switch>
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

<script>

export default {
    name: 'forestUpload',

    data: () => ({
        commentID: -1,
        title: '发表想法',
        valid: false,
        isReal: false,
        message: '',
        messageRules: [
            v => !!v || '请输入内容',
            v => (v && v.length <= 140) || '内容长度不能超过140个字符',
        ],
    }),

    created(){
        if(this.$route.params.id != undefined){
            this.commentID = this.$route.params.id;
            // this.title = '发表评论 - ' + this.commentID;
            this.title = '发表评论'
        }
    },

    methods:{
        submit(){
            this.$refs.form.validate()

            if(this.valid == false){
                this.dialog.show('提示','请检查输入的信息是否正确')
                return
            }

            var data = {
                status: Number(Boolean(this.isReal) == true ? 1 : 0),
                message: this.message,
            }

            const url = (this.commentID == -1) ? '/forest' : ('/forest/comments/' + this.commentID)

            this.HTTP.post(url, data, true).then( () => {
                this.snackbar.show("发表成功")
                this.goback()
            })
        }
    }
}
</script>
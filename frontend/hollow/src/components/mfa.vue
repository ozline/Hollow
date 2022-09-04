<template>

<v-card
    title="多因素认证登录(MFA)"
    subtitle="利用Google Authenticator实现二步验证登录"
    style="margin-bottom: 26px;"
    variant="outlined"
>
    <v-card-text class="text--primary">
        启用状态: <b>{{ getMFAStatus() == true ? '已启用' : '未启用' }}</b>
    </v-card-text>
    <v-card-actions>
        <!-- 启用多因素验证 -->
        <v-dialog v-model="activate" persistent>
            <template v-slot:activator="{ props }">
                <v-btn variant="outlined" color="success" v-bind="props" @click="refershQRCode" v-show="getMFAStatus() == false">启用多因素认证</v-btn>
            </template>
            <v-card title="启用MFA" >
                <v-card-text>
                    <v-container>
                        <v-img
                            :src="qrlink"
                            aspect-ratio="1"
                            cover
                            class="bg-grey-lighten-2"
                            style="margin-bottom:13px"
                        >
                            <template v-slot:placeholder>
                            <v-row
                                class="fill-height ma-0"
                                align="center"
                                justify="center"
                            >
                                <v-progress-circular
                                indeterminate
                                color="grey-lighten-5"
                                ></v-progress-circular>
                            </v-row>
                            </template>
                        </v-img>
                        <v-text-field
                            v-model="code"
                            :rules="codeRules"
                            :counter="6"
                            label="数字动态码"
                            required
                        ></v-text-field>
                    </v-container>
                    <small>*手机中打开Google Authenticator(身份验证器) 扫描二维码,输入6位动态码验证</small>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="blue-darken-1" text @click="activate = false">关闭</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn color="blue-darken-1" text @click="refershQRCode">刷新二维码</v-btn>
                    <v-btn color="error" text @click="activateMFA">启用验证</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <!-- 关闭多因素认证 -->
        <v-dialog v-model="cancel" persistent>
            <template v-slot:activator="{ props }">
                <v-btn variant="outlined" color="error" v-bind="props" v-show="getMFAStatus() == true">关闭多因素认证</v-btn>
            </template>
            <v-card title="关闭MFA" >
                <v-card-text>
                    <v-text-field
                            v-model="code"
                            :rules="codeRules"
                            :counter="6"
                            label="6位数字代码"
                            required
                    ></v-text-field>
                    <small>*若要关闭多因素认证,请在手机中打开Google Authenticator(身份验证器),输入6位动态码</small>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="blue-darken-1" text @click="cancel = false">保留多因素认证</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn color="error" text @click="cancelMFA">确认取消</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-card-actions>
</v-card>

</template>

<script>
import { defineComponent } from 'vue-demi';

export default defineComponent({
    name: 'ComponentMFA',
    data(){
        return{
            activate: false,
            cancel: false,
            qrlink: '',
            secret: '',
            code: '',
            codeRules: [
                v => !!v || '请输入6位代码',
                v => (v && v.length == 6) || '请输入合适的6位代码',
            ],
        }
    },
    methods:{
        getMFAStatus(){
            return (this.global.user.mfaEnabled == true) //防止出现undefined
        },
        refershQRCode(){
            this.code = ''
            this.secret = ''
            this.HTTP.get("/user/mfa",{ },true).then(res => {
                this.qrlink = res.data.qrlink
                this.secret = res.data.secret
            })
        },
        activateMFA(){
            var data = {
                secret: this.secret,
                code: this.code,
            }

            this.HTTP.post("/user/mfa",data,true).then( () => {
                this.global.logout()
                this.$router.push("/user/login")
                this.snackbar.show("MFA绑定成功,已自动退出账号,请重新登录",2000)
            })

            this.code = ''
        },
        cancelMFA(){
            this.HTTP.delete("/user/mfa",{ code: this.code },true).then( () => {
                this.global.logout()
                this.$router.push("/user/login")
                this.snackbar.show("MFA解绑成功,已自动退出账号,请重新登录",2000)
            })
        },
    },
})
</script>
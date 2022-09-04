<template>

<v-card
    title="手机号更新"
    subtitle="更换当前手机号"
    style="margin-bottom: 26px;"
    variant="outlined"
>
    <v-card-text class="text--primary">
        当前手机号: <b>{{ getPhoneNumber() }}</b>
    </v-card-text>
    <v-card-actions>
        <!-- 关闭手机号登录 -->
        <v-dialog v-model="activate" persistent>
            <template v-slot:activator="{ props }">
                <v-btn variant="outlined" color="success" v-bind="props" v-show="false"></v-btn>
            </template>
        </v-dialog>
        <!-- 换绑手机号 -->
        <v-dialog v-model="cancel" persistent>
            <template v-slot:activator="{ props }">
                <v-btn variant="outlined" color="success" v-bind="props">更换手机号</v-btn>
            </template>
            <v-card title="换绑手机号" >
                <v-card-text>
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
                                :rules="codeRules"
                                :counter="6"
                                label="验证码"
                                variant="outlined"
                                required
                            ></v-text-field>
                        </v-col>
                    </v-row>
                    <v-text-field
                            v-show="getMFAStatus"
                            v-model="MFAcode"
                            :rules="MFAcodeRules"
                            :counter="6"
                            label="6位数字代码"
                            required
                    ></v-text-field>
                    <small>*请输入新的手机号并进行验证,若已绑定MFA,需要额外输入6位动态码</small>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="blue-darken-1" text @click="cancel = false">关闭</v-btn>
                    <v-btn color="blue-darken-1" text @click="postShortMsg" :loading="shortMsgLoading">发送验证码</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn color="error" text @click="bindNewPhoneNumber">更新手机号</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-card-actions>
</v-card>

</template>

<script>
import { defineComponent } from 'vue-demi';

export default defineComponent({
    name: 'ComponentBindPhone',
    data(){
        return{
            cancel: false,
            code: '',
            codeRules: [
                v => !!v || '请输入验证码',
                v => (v && v.length === 6) || '验证码格式不正确',
                v => (v && /^[0-9]\d{5}$/.test(v)) || '验证码格式不正确',
            ],
            MFAcode: '',
            MFAcodeRules: [
                v => !!v || '请输入6位代码',
                v => (v && v.length == 6) || '请输入合适的6位代码',
            ],
            phone: '',
            phoneRules: [
                v => !!v || '请输入手机号',
                v => (v && v.length === 11) || '手机号格式不正确',
                v => (v && /^1[3456789]\d{9}$/.test(v)) || '手机号格式不正确',
            ],
            shortMsgLoading: false,
            submitLoading: false,
        }
    },
    methods:{
        getMFAStatus(){
            return (this.global.user.mfaEnabled == true) //防止出现undefined
        },
        getPhoneNumber(){
            return this.global.user.phone
        },
        postShortMsg(){
            var data = {
                phone : this.phone
            }
            this.HTTP.post('/user/register/shortmsg',data).then( () => {
                this.snackbar.show("获取短信验证码成功,1分钟后可再次获取",3000)
                this.shortMsgLoading = true
                setTimeout(() => {
                    this.shortMsgLoading = false
                }, 60000);
            })
        },
        bindNewPhoneNumber(){
            if (this.phone == this.global.user.phone){
                this.dialog.show("错误","新手机号与原先手机号相同")
                return
            }
            var data = {
                phone: this.phone,
                code: this.code,
                mfacode: this.MFAcode,
            }

            this.submitLoading = true
            setTimeout(() => {
                    this.submitLoading = false
            }, 5000);
            this.HTTP.post('/user/rebind/phone',data,true).then( () => {
                this.submitLoading = false;
                this.snackbar.show("换绑手机号成功，请重新登录")
                this.global.logout();
                this.$router.push("/user/login")
            })
        }
    },
})
</script>
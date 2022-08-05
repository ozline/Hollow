<template>

<v-dialog v-model="dialog.show">
    <v-card>
        <v-card-title>{{ dialog.title }}</v-card-title>
        <v-card-text>{{ dialog.content }}</v-card-text>
        <v-card-actions>
            <v-btn color="primary" block @click="eventConfirm">确认</v-btn>
        </v-card-actions>
    </v-card>
</v-dialog>

</template>

<script>
import { computed, defineComponent } from 'vue-demi';
export default defineComponent({
    name: 'ComponentDialog',

    setup(props,ctx)
    {
        const dialog = computed({
            get: () => props.data,
            set: val => ctx.emit('update', { dialog: val })
        });
        return{
            dialog
        }
    },

    props: {
        data: {
            type: Object,
            default: () => ({
                show: {
                    type: Boolean,
                    default: false,
                },
                title: {
                    type: String,
                    default: "null",
                },
                content: {
                    type: String,
                    default: "null",
                },
            }),
        },
        confirm: {
            type: Function,
            default: null,
        },
        cancel: {
            type: Function,
            default: null,
        },
    },

    methods: {
        eventConfirm(){
            if(this.confirm != null) this.confirm();
            this.dialog.show = false;
        },
        eventCancel(){
            this.dialog.show = false;
            if(this.cancel != null) this.cancel();
        }
    }
})
</script>
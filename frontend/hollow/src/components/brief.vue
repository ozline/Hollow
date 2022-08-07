<template>

</template>

<script>
import { computed, defineComponent } from 'vue-demi';
export default defineComponent({
    name: 'ComponentBrief',

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
                id: {
                    type: BigInt,
                    default: 0,
                },
                owner: {
                    type: BigInt,
                    default: 0,
                },
                message: {
                    type: String,
                    default: "NULL",
                },
                createAt: {
                    type: String,
                    default: "NULL",
                },
                status: {
                    type: BigInt,
                    default: 0,
                },
            }),
        },
        //TODO: 传入回调函数
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
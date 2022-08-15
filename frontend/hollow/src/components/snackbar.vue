<template>
    <v-snackbar
        v-model="snackbar.show"
        multi-line
    >
        {{ snackbar.context }}

        <template v-slot:actions>
            <v-btn
                color="red"
                variant="text"
                @click="snackbar.show = false"
            >
                关闭
            </v-btn>
        </template>
    </v-snackbar>
</template>

<script>
import { defineComponent } from "vue-demi";

export default defineComponent({
    name: "ComponentSnackbar",
    setup() {
        const snackbar = computed({
            get: () => props.data,
            set: val => ctx.emit('update', { snackbar: val })
        });
        return{
            snackbar
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
                content: {
                    type: String,
                    default: "null",
                },
            }),
        },
    }
});
</script>
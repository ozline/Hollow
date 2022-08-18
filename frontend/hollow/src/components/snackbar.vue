<template>
    <v-snackbar
        v-model="message.show"
        :timeout="message.timeout"
    >
        {{ message.content }}

        <template v-slot:actions>
            <v-btn
                :color="message.color"
                variant="text"
                @click="message.show = false"
            >关闭</v-btn>
        </template>
    </v-snackbar>
</template>

<script>
import { computed, defineComponent } from 'vue-demi';
import { snackbarStore } from '../store/snackbar';

export default defineComponent({
    name: 'ComponentSnackbar',

    setup(){
        const snackbar = snackbarStore();
        const message = computed({
            get: () => snackbar.data,
            set: val => snackbar.update(val)
        })

        return{ message }
    },
})
</script>
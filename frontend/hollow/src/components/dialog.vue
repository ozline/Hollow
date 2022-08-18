<template>

<v-dialog v-model="message.show" multi-line>
    <v-card>
        <v-card-title>{{ message.title }}</v-card-title>
        <v-card-text>{{ message.content }}</v-card-text>
        <v-card-actions>
            <v-btn color="primary" block @click="message.show = false">确认</v-btn>
        </v-card-actions>
    </v-card>
</v-dialog>

</template>

<script>
import { computed, defineComponent } from 'vue-demi';
import { dialogStore } from '../store/dialog';

export default defineComponent({
    name: 'ComponentMessage',

    setup(){
        const dialog = dialogStore();
        const message = computed({
            get: () => dialog.data,
            set: val => dialog.update(val)
        })

        return{ message }
    },
})
</script>
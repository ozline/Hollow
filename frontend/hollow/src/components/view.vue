<template>
    <v-main fill-height>
        <router-view v-slot="{ Component }" style="outline: 0;">
            <transition name="fade" mode="out-in">
                <component :is="Component" />
            </transition>
        </router-view>
    </v-main>
</template>

<style>
/* 渐变设置 */
.fade-enter-from, .fade-leave-to {
    transform: translateX(20px);
    opacity: 0;
}
.fade-enter-to, .fade-leave-from {
    opacity: 1;
}
.fade-enter-active {
    transition: all 0.5s ease;
}
.fade-leave-active {
    transition: all 0.2s cubic-bezier(1, 0.6, 0.6, 1);
}
</style>

<script>
export default {
    name: 'ComponentView',
}
</script>
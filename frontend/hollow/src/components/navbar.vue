<template>

<v-app-bar :color="navColor" prominent>
    <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title>{{ storeNavbar.title }} </v-toolbar-title>
    <v-toolbar-items>
        <v-switch v-model="model" color="dark" inset style="margin-top: 3px;"></v-switch>
    </v-toolbar-items>
</v-app-bar>

<v-navigation-drawer v-model="drawer">
    <v-list>
    <v-list-item
        v-for="(item, index) in storeNavbar.items"
        :key="index"
        :value="item.value"
        :title="item.title"
        :prependIcon="item.props.prependIcon"
        @click="changePage(item.value)"
        rounded="xl"
    ></v-list-item>
    </v-list>
    <template v-slot:append>
        <div class="pa-2" style="text-align : center;">
            {{ new Date().getFullYear() }} — <strong>Hollow</strong>
        </div>
    </template>
</v-navigation-drawer>

</template>

<script>

import { navbarStore } from '../store/navbar';

export default {
    name: 'ComponentNavbar',
    setup(){
        const storeNavbar = navbarStore();
        return{
            storeNavbar,
        }
    },
    data: () => ({
        drawer: true,
        model: false,
        navColor: 'primary',
    }),
    watch: {
        model(){
            this.$vuetify.theme.global.name = (this.model == true) ? "darkTheme" : "lightTheme";
            this.navColor = (this.model == true) ? '' : 'primary';
        }
    },
    methods: {
        changePage(page){
            this.$router.push(page);
        }
    }
}

</script>

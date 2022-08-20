<template>

<v-app>
  <v-layout>
    <navbar></navbar>
    <router-view id="router-view"></router-view>

  </v-layout>

  <dialog-view></dialog-view>
  <snackbar></snackbar>
</v-app>

</template>

<style>

#app {
  overflow-y: hidden;
  height:70%;
}

#router-view:hover {
  overflow-y: overlay;
}

#router-view {
  /* height: 70%; */
  width: 100%;
  overflow-y: hidden;
  overflow-x: hidden;
}

</style>

<script>
import ComponentNavbar from './components/navbar.vue'
import ComponentRouterView from './components/view.vue'
import ComponentMessage from './components/dialog.vue'
import ComponentSnackbar from './components/snackbar.vue'

export default {
  name: 'App',
  setup(){
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

    if(mediaQuery.matches)
    {
      console.log("system dark mode");
      //TODO: 切换主题~
    }
  },
  components: {
    'navbar' : ComponentNavbar,
    'router-view' : ComponentRouterView,
    'dialog-view' : ComponentMessage,
    'snackbar' : ComponentSnackbar
  },
}
</script>

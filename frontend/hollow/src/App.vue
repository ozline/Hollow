
<template>
<v-app>
  <v-layout>
    <navbar></navbar>
    <router-view></router-view>
    <dialog-view></dialog-view>
  </v-layout>
</v-app>
</template>

<script>
import ComponentNavbar from './components/navbar.vue'
import ComponentRouterView from './components/view.vue'
import ComponentMessage from './components/dialog.vue'

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
    'dialog-view' : ComponentMessage
  },
}
</script>

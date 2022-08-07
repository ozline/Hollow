
<template>
<v-app>
  <v-layout>
    <ComponentNavbar></ComponentNavbar>
    <ComponentRouterView></ComponentRouterView>
  </v-layout>
</v-app>
</template>

<script>
import ComponentNavbar from './components/navbar.vue'
import ComponentRouterView from './components/view.vue'

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
    ComponentNavbar,
    ComponentRouterView
  },
}
</script>

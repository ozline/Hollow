
<template>
<v-app>
  <v-layout>
    <navbar></navbar>
    <routerView></routerView>
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
    'navbar' : ComponentNavbar,
    'routerView' : ComponentRouterView
  },
}
</script>

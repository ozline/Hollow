<template>
  <v-card>
    <v-layout>
      <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

      <v-app-bar
        color="primary"
        prominent
      >
        <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

        <v-toolbar-title>Hollow
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn variant="text" icon="mdi-magnify"></v-btn>

        <v-btn variant="text" icon="mdi-filter"></v-btn>

        <v-btn variant="text" icon="mdi-dots-vertical"></v-btn>
      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
        bottom
        temporary
      >
        <v-list
          :items="items"
        ></v-list>
      </v-navigation-drawer>

      <v-main>
        <v-card-text>
          The navigation drawer will appear from the bottom on smaller size screens.
        </v-card-text>
      </v-main>
    </v-layout>
  </v-card>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',

  components: {
    // HelloWorld,
  },

  data: () => ({
    drawer : false,
    group: null,
    items: [
      {
        title: 'Foo',
        value: 'Foo',
        props: {
          prependIcon: 'mdi-account',
        }
      },
      {
        "title": 'Bar',
        "value": 'Bar',
      },
      {
        "title": 'Fizz',
        "value": 'Fizz',
      },
      {
        "title": 'Buzz',
        "value": 'Buzz',
      },
    ]
  }),

  watch: {
    group() {
      this.drawer = false;
    },
  },
}
</script>

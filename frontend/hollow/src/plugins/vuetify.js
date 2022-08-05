// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import colors from 'vuetify/lib/util/colors'

// Vuetify
import { createVuetify } from 'vuetify'

export default createVuetify(
  // https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
  {
    theme: {
      defaultTheme: 'lightTheme',
      themes: {
        lightTheme: {
          dark: false,
          colors: {
            primary: colors.blue.darken1,
          },
        },
        darkTheme: {
          dark: true,
          colors: {
            primary: colors.blue.darken2,
          }
        },
      }
    }
  }
)

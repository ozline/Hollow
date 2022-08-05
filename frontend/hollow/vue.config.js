const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,

  pluginOptions: {
    vuetify: {
			// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vuetify-loader
		}
  },
  devServer:{
    host : 'localhost',
    port : 8080,
    proxy: {
      '/apis': {
        target: 'http://localhost:9001/',
        changeOrigin: true,
        pathRewrite: {
          '^/apis': '/api'
        }
      },
    }
  }
})

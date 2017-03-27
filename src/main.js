import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import App from './App.vue'
import store from './store/store.js'

Vue.use(ElementUI)

new Vue({
  el: '#app',
  store,
  render: h => h(App)
})

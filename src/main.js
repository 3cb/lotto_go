import Vue from 'vue'
import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/en'
import 'element-ui/lib/theme-default/index.css'
import App from './App.vue'
import store from './store/store.js'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(ElementUI, { locale })
Vue.use(VueAxios, axios)

new Vue({
  el: '#app',
  store,
  render: h => h(App)
})

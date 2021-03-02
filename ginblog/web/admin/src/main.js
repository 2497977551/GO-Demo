import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import './plugin/antui' 
import axios from 'axios'
import { Button } from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
Vue.use(Button)
axios.defaults.baseURL = 'http://localhost:9090/Api/V1'
Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

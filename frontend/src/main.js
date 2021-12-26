import Vue from 'vue'
import App from './App.vue'
import router from "@/router";

Vue.config.productionTip = false

// root组件
new Vue({
  el:"#app",
  render: h => h(App),
  router,
}).$mount('#app')

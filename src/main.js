import Vue from 'vue';
import App from '@/App.vue';
import VModal from 'vue-js-modal';
import router from "@/router/index.js";
import BootstrapVue from 'bootstrap-vue';
import store from '@/store/index.js'
import '@/styles/index.scss';

Vue.config.productionTip = false

Vue.use(BootstrapVue);
Vue.use(VModal);

new Vue({
  router,
  render: h => h(App),
  store
}).$mount('#app')
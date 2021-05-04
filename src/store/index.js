import Vue from 'vue';
import Vuex from 'vuex';
import master_list from './master_list.js';
import loading from './loading.js';

Vue.use(Vuex);

const module = {
    modules:{
      master_list,
      loading
    }
}

export default new Vuex.Store(module);

<template>
<div>
  <div class="wrapper centering">
    <div class="main_div">
      <leftMenu>
      </leftMenu> 
      <router-view></router-view>
    </div>

  </div>
  <div v-if="is_loading">
    <Loading></Loading>
  </div>
</div>
</template>
<script>
//import master_list from '@/api/master_list.js';
//import Header from '@/components/Layout/Header.vue';
import leftMenu from '@/components/Layout/leftMenu.vue';
//import MainContents from '@/components/Layout/MainContents.vue';
import Loading from '@/components/Parts/Loading.vue';

export default {
  name: 'Main',
  components: {
    leftMenu,
    //MainContents,
   // Header,
    Loading
  },
  data() {
    return {
      is_loading: null,
      score: []
    }
  },
  methods: {
    bootModal(modal_name) {
      this.$modal.show(modal_name);
    }
  },
  computed: {
    is_login: {
      get() {
        return this.$store.getters["settings/is_login"];
      }
    },
    master_list: {
      get() {
        return this.$store.getters["master_list/get_master_list"];
      }
    }
  },
  mounted() {

  },
  created() {
    //const master = new master_list();
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
h3 {
  margin: 40px 0 0;
}

.centering {
  margin-right: auto;
  margin-left: auto;
}

.inner_centering {
  display: flex;
  justify-content: center;
}

.left{
  flex:0 0 50px;
}

.main{
  flex:1 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}


</style>
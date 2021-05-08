<template>
<div>
  <div class="wrapper centering">
    <div class="main_div">
      <MainContents></MainContents>
    </div>
  </div>
</div>
</template>
<script>
//import master_list from '@/api/master_list.js';
import MainContents from '@/components/Layout/MainContents.vue';

export default {
  name: 'Main',
  components: {
    MainContents,
  },
  data() {
    return {
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
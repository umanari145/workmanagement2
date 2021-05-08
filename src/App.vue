<template>
  <div>
    <div class="main_div wrapper">
      <leftMenu></leftMenu>
      <div>
        <div>
          header(ログインユーザー情報や経過時間など)
        </div>
        <router-view></router-view>
      </div>
    </div>
    <div v-if="is_loading">
      <Loading></Loading>
    </div>
  </div>
</template>

<script>
import leftMenu from '@/components/Layout/leftMenu.vue';
import Loading from '@/components/Parts/Loading.vue';

export default {
  name: 'App',
  components: {
    leftMenu,
    Loading
  },
  data() {
    return {
      is_loading: null
    }
  },
}
</script>

<style>
</style>

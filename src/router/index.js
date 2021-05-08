import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '@/components/Main.vue'
import Reserve from '@/components/Views/Reserve.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/home',
    component: Main
  },
  {
    path: '/reserve',
    component: Reserve
  }
];

const router = new VueRouter({
  mode: 'history',
  hash:false,
  base: process.env.BASE_URL,
  routes
})

export default router

import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'


Vue.use(VueRouter)


  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/all_data',
    name: 'All_Data',   
    component: () => import('../views/All_Data.vue')
  },
  {
    path: '/buyers_list',
    name: 'Buyers_List',
    component: () => import('../views/Buyers_List.vue')
  },

]

const router = new VueRouter({
  routes
})

export default router
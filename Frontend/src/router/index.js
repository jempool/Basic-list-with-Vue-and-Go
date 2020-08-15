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
    path: '/buyerslist',
    name: 'BuyersList',
    component: () => import('../views/BuyersList.vue')
  },
  {
    path: '/buyer',
    name: 'Buyer',   
    component: () => import('../views/Buyer.vue')
  },
  {
    path: '/all_items',
    name: 'All_Items',   
    component: () => import('../views/All_Items.vue')
  }

]

const router = new VueRouter({
  routes
})

export default router
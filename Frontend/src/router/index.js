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
    path: '/alldata',
    name: 'AllData',   
    component: () => import('../views/AllData.vue')
  }

]

const router = new VueRouter({
  routes
})

export default router
import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Regions from '../views/Regions.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/Regions',
    name: 'Regions',
    component: Regions
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

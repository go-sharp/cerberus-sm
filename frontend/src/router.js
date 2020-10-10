import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'

Vue.use(VueRouter)

const routes = [{ component: HelloWorld, name: 'Home', path: '/home' }]

const router = new VueRouter({
  mode: 'abstract', // mode must be set to 'abstract'
  routes,
})

export default router
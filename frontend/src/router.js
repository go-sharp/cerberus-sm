import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import ServiceOverview from './components/ServiceOverview.vue'
import AddService from './components/AddService.vue'

Vue.use(VueRouter)

export const OverviewRoute = "Overview";
export const AddRoute = "Add";

const routes = [
  { component: HelloWorld, name: 'Home', path: '/home' },
  { component: ServiceOverview, name: OverviewRoute, path: '/overview' },
  { component: AddService, name: AddRoute, path: '/add', props: true }
];

const router = new VueRouter({
  mode: 'abstract', // mode must be set to 'abstract'
  routes,
})

export default router
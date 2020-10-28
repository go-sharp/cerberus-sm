import Vue from 'vue'
import VueRouter from 'vue-router'
import ServiceOverview from './components/ServiceOverview.vue'
import AddService from './components/AddService.vue'
import EditService from './components/EditService.vue'
import NamedRoutes from './routes';

Vue.use(VueRouter)

const routes = [
  { component: ServiceOverview, name: NamedRoutes.OverviewRoute, path: '/' },
  { component: AddService, name: NamedRoutes.AddRoute, path: '/add', props: true },
  { component: EditService, name: NamedRoutes.EditRoute, path: '/edit/:svc', props: true }
];

const router = new VueRouter({
  mode: 'abstract', // mode must be set to 'abstract'
  routes,
})

export default router
import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import router from './router';
import Page from "./components/Page.vue";

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

// Styling with buefy
import Buefy from 'buefy';
import '@mdi/font/css/materialdesignicons.min.css';

Vue.use(Buefy);

// Register global components
Vue.component("page", Page);

Wails.Init(() => {
	new Vue({
		render: h => h(App),
		router,
		mounted() {
			this.$router.replace('/');
		},
	}).$mount('#app');
});

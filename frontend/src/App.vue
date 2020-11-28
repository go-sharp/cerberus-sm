<template>
  <div id="app">
    <navbar id="navbar"></navbar>
    <div class="main">
      <router-view></router-view>
    </div>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
  </div>
</template>

<script>
import "./assets/css/main.css";
import Navbar from './components/Navbar.vue';
import { LOADING_EVENT } from "./util";

export default {
  name: "app",
  components: {
    navbar: Navbar,
  },
  data() {
    return {
      isLoading: false,
    };
  },
  // Lifecylce hooks
  created: function () {
    window.wails.Events.On(
      LOADING_EVENT,
      (isLoading) => (this.isLoading = !!isLoading)
    );
  },
};
</script>

<style lang="scss">
// Import custom colors and styles
@import "styles";

@import "~bulma";
@import "~buefy/src/scss/buefy";

.main {
  height: 100%;
  background-color: $background-color;
  overflow: hidden;
}

#navbar {
  margin-bottom: 2rem;
}

.title {
  color: $primary;
}

.label, .help {
  color: $background-contrast;
}

</style>
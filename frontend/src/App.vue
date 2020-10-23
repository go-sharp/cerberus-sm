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
import Navbar from "./components/Navbar.vue";
import { LOADING } from "./util";

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
      LOADING,
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
  background-color: white;
  border-radius: 5px;
  box-shadow: 5px 5px 0px 0px rgba(0, 0, 0, 0.6);
  overflow: hidden;
}

#navbar {
  margin-bottom: 2rem;
}
</style>
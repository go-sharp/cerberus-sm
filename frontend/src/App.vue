<template>
  <div id="app">
    <navbar id="navbar"></navbar>
    <main>
      <router-view></router-view>
    </main>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
  </div>
</template>

<script>
import "./assets/css/main.css";
import Navbar from "./components/Navbar.vue";

export default {
  name: "app",
  components: {
    "navbar": Navbar
  },
  data() {
    return {
      isLoading: false,
    };
  },
};
</script>

<style lang="scss">
// Import Bulma's core
@import "~bulma/sass/utilities/_all";

// Set your colors
$primary: #387891;

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
    "white": ($white, $black),
    "black": ($black, $white),
    "light": ($light, $light-invert),
    "dark": ($dark, $dark-invert),
    "primary": ($primary, $primary-invert),
    "info": ($info, $info-invert),
    "success": ($success, $success-invert),
    "warning": ($warning, $warning-invert),
    "danger": ($danger, $danger-invert),
);

// Links
$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

@import "~bulma";
@import "~buefy/src/scss/buefy";

main {
  height: 90%;
  background-color: white;
  border-radius: 5px;
  box-shadow: 5px 5px 0px 0px rgba(0, 0, 0, 0.6);
  overflow-x: scroll;
}

#navbar {
  margin-bottom: 2rem;
}

</style>
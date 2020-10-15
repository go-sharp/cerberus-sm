<template>
  <b-navbar>
    <template slot="brand">
      <b-navbar-item tag="div">
        <h1 class="title">Cerberus-SM</h1>
      </b-navbar-item>
    </template>
    <!-- Navbar right -->
    <template slot="end">
      <b-navbar-item tag="div">
        <router-link :to="{ name: this.OverviewRoute }">Overview</router-link>
      </b-navbar-item>
      <b-navbar-item tag="div">
        <router-link :to="{ name: this.AddRoute }">Add Service</router-link>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
import { OverviewRoute, AddRoute } from "../router";

export default {
  data() {
    return {
      OverviewRoute,
      AddRoute,
    };
  },
};
</script>

<style scoped>
.navbar {
  border-radius: 5px;
  box-shadow: 5px 5px 0px 0px rgba(0, 0, 0, 0.6);
}

* >>> .navbar-menu.is-active {
  border-radius: 0 0 5px 5px;
}

h1 {
  color: #387891;
}
</style>

<template>
  <div class="level">
    <div class="level-left">
      <div class="level-item">
        <h1 class="title">Cerberus-SM</h1>
      </div>
    </div>
    <div class="level-right">
      <div class="level-item">
        <router-link :to="{ name: this.OverviewRoute }">Overview</router-link>
      </div>
      <div class="level-item">
        <router-link :to="{ name: this.AddRoute }">Add Service</router-link>
      </div>
    </div>
  </div>
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
.level {
  padding: 0.5rem;
  background-color: white;
  border-radius: 5px;
  box-shadow: 5px 5px 0px 0px rgba(0, 0, 0, 0.6);
}

* >>> .navbar-menu.is-active {
  border-radius: 0 0 5px 5px;
}
</style>

<template>
  <div id="app">
    <router-view></router-view>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>    
    <router-link to="/home">Go to HelloWorld</router-link>
  </div>
</template>

<script>
import "./assets/css/main.css";

export default {
  name: "app",
  data() {
    return {
      isLoading: false,
    };
  },
};
</script>

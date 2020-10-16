<template>
  <page title="Service Overview">
    <button @click="isLoading(true)">Call Loading</button>
    <button @click="refreshServices">Refresh services</button>
  </page>
</template>

<script>
import { isLoading } from "../util";
import * as Wails from '@wailsapp/runtime';
 

export default {
  methods: {
    isLoading,
    refreshServices: function() {
      window.backend.Services.ReloadServices().catch(reason => console.log("-->", reason));
      console.log("=>", this.svcStore);
    }
  },
  // Lifecycle hooks
  created: function() {
    this.svcStore = Wails.Store.New("Services");
    console.warn(this.svcStore.subscribe((state) => {
        console.log(">>>", state);
    }));
  }
};
</script>

<style lang="scss" scoped>
</style>
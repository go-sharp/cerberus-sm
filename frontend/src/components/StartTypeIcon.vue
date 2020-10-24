<template>
  <b-tooltip :label="currentStartType.tooltip" position="is-left" style="cursor: help" :delay="1000">
    <b-icon :icon="currentStartType.icon"></b-icon>
  </b-tooltip>
</template>

<script>
export default {
  props: {
    startType: Number,
  },
  computed: {
    currentStartType() {
      switch (this.startType) {
        case 2:
          return { icon: "auto-upload", tooltip: "Automatic" };
        case 3:
          return { icon: "hand-left", tooltip: "Manual" };
        case 4:
          return { icon: "close-circle", tooltip: "Disabled" };
        default:
          return { icon: "", tooltip: "" };
      }
    },
  },
};
</script>


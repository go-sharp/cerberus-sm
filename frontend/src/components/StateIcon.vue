<template>
  <b-tooltip :label="currentState.tooltip" position="is-left" style="cursor: help" :active="currentState.tooltip !== ''" :delay="1000">
    <b-icon :icon="currentState.icon"></b-icon>
  </b-tooltip>
</template>

<script>
export default {
  props: {
    state: Number,
  },
  computed: {
    currentState() {
      switch (this.state) {
        case 1:
          return { icon: "stop-circle", tooltip: "Stopped" };
        case 2:
          return { icon: "arrow-up-circle", tooltip: "Start pending" };
        case 3:
          return { icon: "arrow-down-circle", tooltip: "Stop pending" };
        case 4:
          return { icon: "play-circle", tooltip: "Running" };
        case 5:
          return { icon: "play-speed", tooltip: "Continue pending" };
        case 6:
          return { icon: "pause-motion", tooltip: "Pause pending" };
        case 7:
          return { icon: "pause-circle", tooltip: "Paused" };
        default:
          return { icon: "", tooltip: "" };
      }
    },
  },
};
</script>


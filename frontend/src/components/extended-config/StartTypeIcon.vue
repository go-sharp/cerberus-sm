<template>
  <b-tooltip :label="currentStartType.hint" position="is-left" style="cursor: help" :delay="1000">
    <b-icon :icon="currentStartType.icon"></b-icon>
  </b-tooltip>
</template>

<script>
import { StartTypes } from './StartTypeDropdown';

export default {
  props: {
    startType: Number,
  },
  computed: {
    currentStartType() {
      return StartTypes.find(e => e.value === this.startType) || { hint: '', icon: ''};      
    },
  },
};
</script>


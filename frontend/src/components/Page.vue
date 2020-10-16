<template>
  <div class="page">
    <h1 class="title is-4">{{ title }}</h1>
    <div class="content">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: ["title"],
};
</script>

<style lang="scss" scoped>
@import "../styles";

h1 {
  padding-bottom: 0.5rem;
  border-bottom: 3px solid $primary;  
}

.page {
  padding: 0.75rem;
}

.content {
    overflow-y: auto;
    overflow-y: hidden;
}

</style>
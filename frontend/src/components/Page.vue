<template>
  <div class="page">
    <div class="columns is-gapless page__header">
      <div class="column is-one-third">
        <h1 class="title is-4">{{ title }}</h1>
      </div>
      <div class="column page__menu">
        <slot name="menu"></slot>
      </div>
    </div>
    <div class="columns content">
      <div class="column">
        <slot></slot>
    </div>
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

.page__header {
  padding-bottom: 0.5rem;
  border-bottom: 3px solid $primary;
}

.page__menu {
  display: flex;
  justify-content: flex-end;
}

.page {
  height: 100%;
  padding: 0.75rem;
}

.content {
  max-height: 90%;
  overflow-y: auto;
  overflow-x: hidden;
}
</style>
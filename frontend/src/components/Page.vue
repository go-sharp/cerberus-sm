<template>
    <div class="page">
        <div class="page__header">
            <div class="">
                <h1 class="title is-4">{{ title }}</h1>
            </div>
            <div class="page__menu">
                <slot name="menu"></slot>
            </div>
        </div>
        <div class="content">
            <slot></slot>
        </div>
    </div>
</template>

<script>
export default {
    props: ['title'],
};
</script>

<style lang="scss" scoped>
@import '../styles';

.page__header {
    padding-bottom: 0.5rem;
    border-bottom: 3px outset $background-contrast;
    flex: 0 0 auto;
    display: flex;
    justify-content: flex-start;
    flex-flow: row wrap;
}

.page__menu {
    display: flex;
    justify-content: flex-end;
    flex: 1 0 auto;
}

.page {
    display: flex;
    flex-direction: column;
    height: 100%;
    padding: 0.75rem;
    max-height: 100%;
}

.content {
    flex: 1 1 auto;
    height: 100%;
    max-height: 100%;
    min-height: 50%;
    overflow: hidden;
    overflow-y: auto;
    padding-top: 1rem;
}

.title {
    color: $background-contrast;
}
</style>
<template>
  <section></section>
</template>

<script>

export const AutoStart = 2;
export const AutoStartDelayed = 9999;
export const ManualStart = 3;
export const DisabledStart = 4;

export const StartTypes = [
    {
        text: 'Automatic',
        value: 2,
        hint: 'Service will start on system startup.',
        icon: 'auto-upload',
    },
    {
        text: 'Manual',
        value: 3,
        hint: 'Service must be started manually.',
        icon: 'hand-left',
    },
    {
        text: 'Disabled',
        value: 4,
        hint: 'Service is disabled.',
        icon: 'close-circle',
    },
    {
        text: 'Automatic Delayed',
        value: 9999,
        hint: 'Service will start automatically with a delay.',
        icon: 'history',
    },
];

export default {

}
</script>

<style>

</style>
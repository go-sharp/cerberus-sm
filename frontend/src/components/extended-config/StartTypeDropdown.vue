<template>
    <div class="start-type-dropdown">
        <b-dropdown v-model="selected" aria-role="list">
            <template #trigger>
                <b-button :label="selected.text" type="is-primary" :icon-left="selected.icon" icon-right="menu-down" />
            </template>

            <b-dropdown-item
                v-for="startType in startTypes"
                :key="startType.value"
                :value="startType"
                @click="changeSelection"
                aria-role="listitem"
            >
                <div class="media">
                    <b-icon class="media-left" :icon="startType.icon"></b-icon>
                    <div class="media-content">
                        <h3>{{ startType.text }}</h3>
                        <small>{{ startType.hint }}</small>
                    </div>
                </div>
            </b-dropdown-item>
        </b-dropdown>
    </div>
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
    data: function() {
        return {
            selected: StartTypes.find(s => s.value === this.value) || StartTypes[1],
            startTypes: StartTypes
        }
    },
    props: {
        value: {
            type: Number,
            required: true
        }
    },
    methods: {
        changeSelection: function () {
            this.$emit('input', this.selected.value);
        },
    },
};
</script>

<style>
</style>
<template>
    <b-dropdown v-model="selected" aria-role="list">
        <template #trigger>
            <b-button
                :label="selected.text"
                type="is-primary"
                :icon-left="selected.icon"
                icon-right="menu-down"
            />
        </template>

        <b-dropdown-item
            v-for="action in actions"
            :key="action.value"
            :value="action"
            @click="changeSelection"
            aria-role="listitem"
        >
            <div class="media">
                <b-icon class="media-left" :icon="action.icon"></b-icon>
                <div class="media-content">
                    <h3>{{ action.text }}</h3>
                    <small>{{ action.hint }}</small>
                </div>
            </div>
        </b-dropdown-item>
    </b-dropdown>
</template>

<script>
export const NoAction = 1;
export const Restart = 2;
export const RunProgram = 4;
export const RestartRunProgram = Restart | RunProgram;

export const actions = [
    {
        text: 'No Action',
        value: 1,
        hint: 'Service will stop if an error occurs.',
        icon: 'stop',
    },
    {
        text: 'Restart',
        value: 2,
        hint: 'Service will restart if an error occurs.',
        icon: 'restart',
    },
    {
        text: 'Run Program',
        value: 4,
        hint: 'Service will stop and run a program if an error occurs.',
        icon: 'apps',
    },
    {
        text: 'Restart & Run Program',
        value: 6,
        hint: 'Service will restart and run a program if an error occurs.',
        icon: 'repeat',
    },
];

export default {
    data() {
        return {
            selected: actions.find(a => a.value === this.action),
            actions,
        };
    },
    props: {
        action: {
            type: Number,
            default: NoAction,
        },
    },
    methods: {
        changeSelection: function () {
            //console.log(arguments, arguments);
            console.log(this.selected.text, this.selected);
        },
    },
};
</script>

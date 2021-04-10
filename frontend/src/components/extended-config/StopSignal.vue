<template>
    <b-field label="Stop Signals" class="stop-signal">
        <b-field class="stop-signal__item" message="Sends a Ctrl-C to the process">
            <b-checkbox @input="checkboxChanged($event, signals.CtrlC)" :value="ctrlChecked"> Ctrl-C </b-checkbox>
        </b-field>
        <b-field class="stop-signal__item" message="Sends a WM_CLOSE message to the process">
            <b-checkbox @input="checkboxChanged($event, signals.Close)" :value="closeChecked"> WM_CLOSE </b-checkbox>
        </b-field>
        <b-field class="stop-signal__item" message="Sends a WM_QUIT message to the process">
            <b-checkbox @input="checkboxChanged($event, signals.Quit)" :value="quitChecked"> WM_QUIT </b-checkbox>
        </b-field>
    </b-field>
</template>

<script>
export const CtrlC = 1;
export const Quit = 2;
export const Close = 4;

export default {
    props: {
        value: {
            type: Number,
            required: true,
        },
    },
    data: function () {
        return {
            sendCtrlC: false,
            sendQuit: false,
            sendClose: false,
            signals: {
                Quit,
                Close,
                CtrlC,
            },
        };
    },
    methods: {
        checkboxChanged: function (e, sig) {
            let signal = this.value;
            if (e) signal |= sig;
            else signal &= ~sig;

            this.$emit('input', signal);
        },
    },
    computed: {
        ctrlChecked: function () {
            if ((this.value & CtrlC) === CtrlC) return true;

            return false;
        },
        quitChecked: function () {
            if ((this.value & Quit) === Quit) return true;

            return false;
        },
        closeChecked: function () {
            if ((this.value & Close) === Close) return true;

            return false;
        },
    },
};
</script>

<style lang="scss">
.stop-signal {
    .stop-signal__item {
        &:not(:first-child) {
            margin-left: 0.8rem;
        }
    }
}
</style>
<template>
    <section class="recovery-action-editor">
        <!-- ExitCode -->

        <b-field
            label="Exit Code:"
            class="recovery-action-editor__number-item"
            :type="validExitCode ? '' : 'is-danger'"
            :message="validExitCodeMsg"
        >
            <b-numberinput :disabled="!isNew" v-model="model.exit_code" :controls="false"></b-numberinput>
        </b-field>

        <!-- Action & Delay -->
        <div class="columns">
            <div class="column is-half">
                <b-field label="Action:" message="Action to execute on exit">
                    <action-dropdown v-model="model.action"></action-dropdown>
                </b-field>
            </div>
            <div class="column is-half" v-if="showRestart">
                <b-field
                    label="Delay:"
                    class="recovery-action-editor__number-item"
                    :message="errors.delay || 'Seconds to wait until action takes place'"
                    :type="errors.delay ? 'is-danger' : ''"
                >
                    <b-numberinput v-model="model.delay" :controls="false"></b-numberinput>
                </b-field>
            </div>
        </div>

        <!-- MaxRestarts & ResetAfter -->
        <div class="columns" v-if="showRestart">
            <div class="column is-half">
                <b-field
                    label="Max Restarts:"
                    class="recovery-action-editor__number-item"
                    :message="errors.max_restarts || 'Maximal amount of restarts'"
                    :type="errors.max_restarts ? 'is-danger' : ''"
                >
                    <b-numberinput v-model="model.max_restarts" :controls="false"></b-numberinput>
                </b-field>
            </div>
            <div class="column is-half">
                <b-field
                    label="Reset After:"
                    class="recovery-action-editor__number-item"
                    :message="errors.reset_after || 'Resets restart counter after n seconds'"
                    :type="errors.reset_after ? 'is-danger' : ''"
                >
                    <b-numberinput v-model="model.reset_after" :controls="false"></b-numberinput>
                </b-field>
            </div>
        </div>

        <!-- Run Program Section -->
        <div v-if="showProgram">
            <!-- Executable Path field -->
            <b-field label="Executable Path:" :message="errors.program" :type="errors.program ? 'is-danger' : ''">
                <b-input v-model="model.program" expanded></b-input>
                <p class="control">
                    <button class="button is-primary" @click="showDialog()">Select</button>
                </p>
            </b-field>

            <!-- Arguments field -->
            <b-field label="Arguments:">
                <b-taginput
                    v-model="model.arguments"
                    icon="format-text-variant"
                    ellipsis
                    allow-duplicates
                    placeholder="Add"
                    type="is-primary"
                >
                </b-taginput>
            </b-field>
        </div>
    </section>
</template>

<script>
import ActionDropdown from './ActionDropdown';
import { NoAction, Restart, RestartRunProgram, RunProgram } from './ActionDropdown.vue';
import { createMsg } from '../../util';

export const ValidationChanged = 'validationChanged';

export default {
    components: { ActionDropdown },
    props: {
        value: Object,
        isNew: Boolean,
        excludedExitCodes: Array,
    },
    data() {
        return {
            model: {
                exit_code: 0,
                action: 1,
                delay: 0,
                max_restarts: 0,
                reset_after: 0,
                program: '',
                arguments: [],
                ...this.value,
            },
            isValid: false,
            errors: {
                delay: '',
                max_restarts: '',
                reset_after: '',
                program: '',
            },
        };
    },
    mounted: function () {
        this.isValid = this.validateModel();
        this.$emit(ValidationChanged, this.isValid);
    },
    computed: {
        showProgram: function () {
            if (this.model.action === RunProgram || this.model.action === RestartRunProgram) return true;

            return false;
        },
        showRestart: function () {
            if (this.model.action === Restart || this.model.action === RestartRunProgram) return true;

            return false;
        },
        validExitCode: function () {
            if (!Array.isArray(this.excludedExitCodes) || !this.isNew) return true;

            if (typeof this.model.exit_code !== 'number') return false;            
            return this.excludedExitCodes.find((e) => e === this.model.exit_code) === undefined;
        },
        validExitCodeMsg: function () {
            if (this.validExitCode) return '';

            if (typeof this.model.exit_code === 'number') return 'Exit code already used';

            return 'Empty exit code is not valid';
        },
    },
    methods: {
        validateModel: function () {
            this.errors.delay = typeof this.model.delay === 'number' ? null : 'Must be a number';

            this.errors.max_restarts = typeof this.model.max_restarts === 'number' ? null : 'Must be a number';

            this.errors.reset_after = typeof this.model.reset_after === 'number' ? null : 'Must be a number';

            if (this.model.action === RunProgram || this.model.action === RestartRunProgram) {
                this.errors.program =
                    typeof this.model.program === 'string' && this.model.program.length > 0 ? '' : '* Required';
            } else {
                this.errors.program = '';
            }

            for (const err in this.errors) {
                if (this.errors[err]) return false;
            }

            return this.validExitCode;
        },
        showDialog: function () {
            window.backend.Services.ShowFileDialog({dir: false})
                .then((val) => this.model.program = val)
                .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')));
        },
    },
    watch: {
        isValid: function (val) {
            this.$emit(ValidationChanged, val);
        },
        'model.action': function (val) {
            if (val === NoAction || val === Restart) {
                this.model.program = '';
                this.model.arguments = [];
            }

            if (val === NoAction || val === RunProgram) {
                this.model.delay = 0;
                this.model.max_restarts = 0;
                this.model.reset_after = 0;
            }
        },
        model: {
            handler: function () {
                this.isValid = this.validateModel();

                if (this.isValid) {
                    this.$emit('input', {
                        exit_code: this.model.exit_code,
                        action: this.model.action,
                        delay: this.model.delay,
                        reset_after: this.model.reset_after,
                        max_restarts: this.model.max_restarts,
                        program: this.model.program,
                        arguments: this.model.arguments,
                    });
                }
            },
            deep: true,
        },
    },
};
</script>

<style lang="scss">
.recovery-action-editor {
    .recovery-action-editor__number-item {
        max-width: 250px;
    }
}
</style>

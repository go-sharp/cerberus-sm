<template>
    <section class="recovery-action-editor">
        <!-- ExitCode -->

        <b-field
            label="ExitCode:"
            class="recovery-action-editor__number-item"
            :type="validExitCode ? '' : 'is-danger'"
            :message="validExitCodeMsg"
        >
            <b-numberinput :disabled="!isNew" v-model="model.exitCode" :controls="false"></b-numberinput>
        </b-field>

        <!-- Action & Delay -->
        <div class="columns">
            <div class="column is-half">
                <b-field label="Action:" message="Action to execute on exit">
                    <action-dropdown v-model="model.action"></action-dropdown>
                </b-field>
            </div>
            <div class="column is-half">
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
        <div class="columns">
            <div class="column is-half">
                <b-field
                    label="Max Restarts:"
                    class="recovery-action-editor__number-item"
                    :message="errors.maxRestarts || 'Maximal amount of restarts'"
                    :type="errors.maxRestarts ? 'is-danger' : ''"
                >
                    <b-numberinput v-model="model.maxRestarts" :controls="false"></b-numberinput>
                </b-field>
            </div>
            <div class="column is-half">
                <b-field
                    label="Reset After:"
                    class="recovery-action-editor__number-item"
                    :message="errors.resetAfter || 'Resets restart counter after n seconds'"
                    :type="errors.resetAfter ? 'is-danger' : ''"
                >
                    <b-numberinput v-model="model.resetAfter" :controls="false"></b-numberinput>
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
                    v-model="model.args"
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
        excludeErrorCodes: Array,
    },
    data() {
        return {
            model: {
                exitCode: 0,
                action: 1,
                delay: 0,
                maxRestarts: 0,
                resetAfter: 0,
                program: '',
                args: [],
                ...this.value,
            },
            isValid: false,
            errors: {
                delay: '',
                maxRestarts: '',
                resetAfter: '',
                program: '',
            },
        };
    },
    computed: {
        showProgram: function () {
            if (this.model.action === RunProgram || this.model.action === RestartRunProgram) return true;

            return false;
        },
        validExitCode: function () {
            if (!Array.isArray(this.excludeErrorCodes)) return true;

            if (typeof this.model.exitCode !== 'number') return false;
            return !this.excludeErrorCodes.find((e) => e === this.model.exitCode);
        },
        validExitCodeMsg: function () {
            if (this.validExitCode) return '';

            if (typeof this.model.exitCode === 'number') return 'Exit code already used';

            return 'Empty exit code is not valid';
        },
    },
    methods: {
        validateModel: function () {
            this.errors.delay = typeof this.model.delay === 'number' ? null : 'Must be a number';

            this.errors.maxRestarts = typeof this.model.maxRestarts === 'number' ? null : 'Must be a number';

            this.errors.resetAfter = typeof this.model.resetAfter === 'number' ? null : 'Must be a number';

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
        showDialog: function (isWorkDir) {
            window.backend.Services.ShowFileDialog({
                dir: !!isWorkDir,
            })
                .then((val) => {
                    const prop = isWorkDir ? 'work_dir' : 'exe_path';
                    this.model[prop] = val;
                })
                .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')));
        },
    },
    watch: {
        isValid: function (val) {
            this.$emit(ValidationChanged, val);
        },
        action: function (val) {
            if (val === NoAction || val == Restart) {
                this.program = '';
                this.args = [];
            }
        },
        model: {
            handler: function () {
                this.isValid = this.validateModel();

                if (this.isValid) {
                    this.$emit('input', {
                        exitCode: this.model.exitCode,
                        action: this.model.action,
                        delay: this.model.delay,
                        resetAfter: this.model.resetAfter,
                        maxRestarts: this.model.maxRestarts,
                        program: this.model.program,
                        args: this.model.args,
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

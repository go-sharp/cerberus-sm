<template>
    <div class="recovery-actions">
        <div class="recovery-actions__menu">
            <div class="is-flex-grow-1">Configured Recovery Actions</div>
            <div class="">
                <b-button v-if="!selected" @click="addAction" size="is-small" class="is-primary" icon-left="plus">
                    Add
                </b-button>
                <b-button
                    v-if="selected"
                    @click="updateAction"
                    :disabled="!canSave"
                    size="is-small"
                    class="is-primary"
                    icon-left="content-save-cog"
                >
                    Update
                </b-button>
                <b-button v-if="selected" @click="cancelUpdate" size="is-small" icon-left="cancel"> Cancel </b-button>
            </div>
        </div>
        <div class="recovery-actions__divider"></div>
        <section class="columns">
            <div class="column is-one-half">
                <div class="recovery-actions__items">
                    <template v-for="(item, index) in sortedDataSource">
                        <action-item
                            class="recovery-actions__item"
                            :key="index"
                            :exitCode="item.exit_code"
                            :recoveryAction="item.action"
                            :disabled="!!selected"
                            @edit="editItem(item)"
                            @delete="deleteItem(item)"
                        ></action-item>
                    </template>
                </div>
            </div>
            <div class="column is-one-half">
                <recovery-action-editor
                    v-if="selected"
                    :isNew="isNew"
                    :excludedExitCodes="excludedExitCodes"
                    v-model="selected"
                    @validationChanged="canSave = $event"
                ></recovery-action-editor>
            </div>
        </section>
    </div>
</template>

<script>
import RecoveryActionEditor from './RecoveryActionEditor.vue';
import ActionItem from './ActionItem.vue';
import { NoAction, Restart, RestartRunProgram, RunProgram } from './ActionDropdown.vue';

let actionsSource = [
    { exit_code: 2, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 3, action: RestartRunProgram, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 5, action: RunProgram, delay: 0, max_restarts: 0, reset_after: 0, program: 'c:\\temp\\myprogram.exe' },
    { exit_code: 6, action: NoAction, delay: 0, max_restarts: 0, reset_after: 0 },
    { exit_code: 8, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 123, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
];

export default {
    components: { RecoveryActionEditor, ActionItem },
    props: {
        recoveryActions: {
            type: Array,
            required: false,
        },
    },
    data() {
        return {
            action: { exit_code: 45, delay: 444 },
            selected: null,
            dataSource: actionsSource,
            isNew: false,
            canSave: false,
        };
    },
    computed: {
        excludedExitCodes: function () {
            return this.dataSource.map((a) => a.exit_code);
        },
        sortedDataSource: function () {
            const sorted = [...this.dataSource];
            return sorted.sort((first, second) => {
                if (first.exit_code > second.exit_code) return 1;
                if (first.exit_code < second.exit_code) return -1;
                return 0;
            });
        },
    },
    methods: {
        cancelUpdate: function () {
            this.isNew = false;
            this.selected = null;
        },
        updateAction: function () {
            if (this.isNew) {
                this.dataSource.push(this.selected);
            } else {
                const idx = this.dataSource.findIndex((a) => a.exit_code === this.selected.exit_code);
                this.dataSource[idx] = this.selected;
            }

            this.selected = null;
            this.emitUpdate();
        },
        emitUpdate: function () {
            this.$emit('update:recoveryActions', this.dataSource);
        },
        deleteItem: function (item) {
            const idx = this.dataSource.findIndex((i) => i.exit_code === item.exit_code);
            if (idx < 0) return;
            this.dataSource.splice(idx, 1);
            this.emitUpdate();
        },
        editItem: function (item) {
            this.isNew = false;
            this.selected = item;
        },
        addAction: function () {
            this.isNew = true;
            this.selected = {
                exit_code: 0,
                action: NoAction,
                delay: 0,
                reset_after: 0,
                max_restarts: 0,
                program: '',
                args: [],
            };
        },
    },
};
</script>

<style lang="scss" scoped>
@import '../../styles';

.recovery-actions {
    .recovery-actions__menu {
        padding-left: 5px;
        padding-right: 5px;
        display: flex;
        justify-content: space-between;

        & > *:first-child {
            align-self: flex-end;
        }

        & > *:last-child {
            align-self: flex-start;

            & > * {
                margin-left: 5px;
            }
        }
    }

    .recovery-actions__divider {
        height: 1px;
        margin: 5px;
        background-color: $background-contrast;
    }

    .recovery-actions__items {
        overflow: auto;
        max-height: 70vh;
        padding: 10px;
    }

    .recovery-actions__item {
        // box-shadow: 0 0.5em 1.5em -0.125em #a2a2a2, 0 0px 0 1px #a2a2a2;
        box-shadow: 0 0 0.2em 1px $background-contrast, 0 0px 0 1px $background-contrast;
        padding: 5px;
        margin-left: 5px;

        & + .recovery-actions__item {
            margin-top: 10px;
        }

        &:last-child {
            margin-bottom: 1em;
        }
    }
}
</style>
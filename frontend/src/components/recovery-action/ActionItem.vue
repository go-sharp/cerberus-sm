<template>
    <section class="recovery-action-item">
        <div class="recovery-action-item__content">
            <div class="recovery-action-item__group">
                <b-field label="Exit Code:" custom-class="is-small">{{ exitCode }}</b-field>
                <b-field label="Action:" custom-class="is-small">
                    <b-icon :icon="action.icon" size="is-small"></b-icon>
                    {{ action.text }}
                </b-field>
            </div>
        </div>
        <div class="recovery-action-item__menu">
            <b-button
                type="is-primary"
                size="is-small"
                icon-right="pencil"
                :disabled="disabled"
                @click="$emit('edit', exitCode)"
            />
            <b-button
                type="is-danger"
                size="is-small"
                icon-right="delete"
                :disabled="disabled"
                @click="$emit('delete', exitCode)"
            />
        </div>
    </section>
</template>

<script>
import { actions } from './ActionDropdown.vue';
export default {
    props: {
        exitCode: {
            type: Number,
            required: true,
        },
        recoveryAction: {
            type: Number,
            required: true,
        },
        disabled: Boolean,
    },
    computed: {
        action: function () {            
            const action = actions.find((a) => a.value === this.recoveryAction);
            return action ? { text: action.text, icon: action.icon } : {};
        },
    },
};
</script>

<style lang="scss" scoped>
.flex-wrap {
    flex-wrap: wrap;
}

.recovery-action-item {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .recovery-action-item__group {
        display: flex;
        justify-content: flex-start;
        flex-flow: row wrap;

        & > * {
            min-width: 30%;
        }

        & > *:not(:first-child) {
            margin-left: 5px;
        }
    }

    .recovery-action-item__content {
        flex: 1 0 auto;
        margin-right: 15px;
    }

    .recovery-action-item__menu {
        display: flex;
        justify-content: flex-end;

        & > *:not(:first-child) {
            margin-left: 5px;
        }
    }
}
</style>

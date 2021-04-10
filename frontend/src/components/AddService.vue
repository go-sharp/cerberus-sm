<template>
    <page :title="title">
        <!-- Page Menu -->
        <template v-slot:menu>
            <div class="buttons">
                <b-button
                    size="is-small"
                    type="is-primary"
                    icon-left="content-save-cog"
                    :aria-disabled="!isValid"
                    :disabled="!isValid"
                    @click="installService"
                >
                    Install
                </b-button>
                <b-button size="is-small" icon-left="cancel" tag="router-link" :to="{ name: OverviewRoute }">
                    Cancel
                </b-button>
            </div>
        </template>
        <!-- Page content -->
        <b-tabs v-model="activeTab" :animated="false">
            <b-tab-item value="base" label="Base Configuration">
                <edit-service-base v-model="model" :is-new="true" :field-messages="validations" />
            </b-tab-item>
            <b-tab-item value="extended" label="Extended Configuration" class="extended-configuration">
                <div class="extended-configuration__item">
                    <b-field label="Start Type">
                        <start-type-dropdown :value="model.start_type"></start-type-dropdown>
                    </b-field>
                    <stop-signal v-model="model.stop_signal"></stop-signal>
                    <dependencies-list v-model="model.dependencies"></dependencies-list>
                </div>
                <div class="extended-configuration__item">
                    <service-user :password.sync="model.password" :userName="model.service_user"></service-user>
                </div>
            </b-tab-item>
            <b-tab-item value="recovery" label="Recovery Actions">
                <recovery-actions :recoveryActions.sync="model.recovery_actions"></recovery-actions>
            </b-tab-item>
        </b-tabs>
    </page>
</template>

<script>
import { isLoading, createMsg } from '../util';
import { OverviewRoute } from '../routes';
import EditServiceBase from './EditServiceBase.vue';
import { createEditBaseModel } from './models/service';
import RecoveryActions from './recovery-action/RecoveryActions.vue';
import { NoAction, Restart, RestartRunProgram, RunProgram } from './recovery-action/ActionDropdown.vue';
import StartTypeDropdown from './extended-config/StartTypeDropdown.vue';
import StopSignal from './extended-config/StopSignal.vue';
import DependenciesList from './extended-config/DependenciesList.vue';
import ServiceUser from './extended-config/ServiceUser.vue';

const isValidServiceName = new RegExp(/^[a-zA-Z0-9_+\-!]+$/);

let actionSource = [
    { exit_code: 2, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 3, action: RestartRunProgram, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 5, action: RunProgram, delay: 0, max_restarts: 0, reset_after: 0, program: 'c:\\temp\\myprogram.exe' },
    { exit_code: 6, action: NoAction, delay: 0, max_restarts: 0, reset_after: 0 },
    { exit_code: 8, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
    { exit_code: 123, action: Restart, delay: 15, max_restarts: 1, reset_after: 0 },
];

export default {
    components: { EditServiceBase, RecoveryActions, StartTypeDropdown, StopSignal, DependenciesList, ServiceUser },
    data() {
        return {
            OverviewRoute,
            model: createEditBaseModel(),
            validations: {
                name: { msg: '* required', hasValidated: false },
                exe_path: { msg: '* required', hasValidated: false },
            },
            displayEdited: false,
            activeTab: 'base',
            actionSource: [...actionSource],
            stopSignal: 0,
            dependencies: [],
        };
    },
    watch: {
        'model.name': function (val) {
            if (!this.displayEdited) {
                this.isNameUpdating = true;
                this.model.display_name = val;
            }

            // Validation
            let v = {
                msg: '',
                type: 'is-success',
                hasValidated: true,
            };
            if (!(this.model.name && this.model.name.length > 0)) {
                v = {
                    ...v,
                    msg: '* required',
                    type: 'is-danger',
                };
            } else if (!isValidServiceName.test(this.model.name)) {
                v = {
                    ...v,
                    msg: 'allowed characters: [a-zA-Z0-9_+-!]',
                    type: 'is-danger',
                };
            }

            this.validations = { ...this.validations, name: v };
        },
        'model.display_name': function (val) {
            if (this.isNameUpdating) {
                this.isNameUpdating = false;
                return;
            }

            if (!this.displayEdited && val) this.displayEdited = true;
            else if (this.displayEdited && !val) this.displayEdited = false;
        },
        'model.exe_path': function () {
            // Validation
            let v = {
                msg: '',
                type: 'is-success',
                hasValidated: true,
            };
            if (!(this.model.exe_path && this.model.exe_path.length > 0)) {
                v = {
                    ...v,
                    msg: '* required',
                    type: 'is-danger',
                };
            }

            this.validations = { ...this.validations, exe_path: v };
        },
    },
    computed: {
        title: function () {
            if (!this.model.name) return 'New Service';
            return this.model.name;
        },
        isValid: function () {
            let not_validated = 0;
            for (const key in this.validations) {
                const { type, hasValidated } = this.validations[key];
                if (type && type === 'is-danger') return false;
                if (hasValidated === false) not_validated++;
            }

            return not_validated <= 0;
        },
    },
    methods: {
        installService: function () {
            isLoading(true);
            console.log(this.model);
            window.backend.Services.InstallService(this.model)
                .then(() => {
                    isLoading(false);
                    this.$buefy.toast.open(createMsg('Successfully installed service.', 'success'));
                    this.$router.push({ name: OverviewRoute });
                })
                .catch((reason) => {
                    isLoading(false);
                    this.$buefy.toast.open(createMsg(reason, 'error'));
                });
        },
    },
};
</script>

<style lang="scss" >
@import '../styles';
.extended-configuration {
    display: flex;
    flex-flow: row wrap;
    height: 100%;
    overflow-y: auto;

    .extended-configuration__item {
        &:not(:last-child) {
            margin-right: 2rem;
        }
        flex: 1 1 300px;
        min-width: 300px;
    }
}

.b-tabs {
    height: 100%;
    max-height: 100%;
    display: flex;
    flex-flow: column nowrap;

    .tabs {
        ul {
            margin: 0;
            border-bottom-color: $background-contrast;
            li {
                &.is-active a {
                    color: $primary;
                    border-bottom-color: $primary;
                    &:hover {
                        color: $primary;
                        border-bottom-color: $primary;
                    }
                }

                a {
                    &:hover {
                        color: white;
                        border-bottom-color: white;
                    }
                    color: $background-contrast;
                    border-bottom-color: $background-contrast;
                }
            }
        }
    }

    .tab-content.tab-content {
        padding: 1rem;
        height: 100%;
        max-height: 100%;
        overflow: hidden;

        &.is-transitioning {
            max-height: 100%;
            height: 100%;

            .tab-item {
                height:200%;                
            }
        }

        .tab-item {
            max-height: 100%;
            height: 100%;
            overflow-y: auto;

            & > * {
                max-width: 100%;
            }
        }
    }
}
</style>
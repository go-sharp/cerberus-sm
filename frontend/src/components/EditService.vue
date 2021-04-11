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
                    v-if="isNew"
                >
                    Install
                </b-button>
                <b-button
                    size="is-small"
                    type="is-primary"
                    icon-left="content-save-cog"
                    @click="updateService"
                    :disabled="!(hasChanges && isValid)"
                    v-if="!isNew"
                >
                    Update
                </b-button>
                <b-button size="is-small" icon-left="cancel" tag="router-link" :to="{ name: OverviewRoute }">
                    Cancel
                </b-button>
            </div>
        </template>
        <!-- Page content -->
        <b-tabs v-model="activeTab" :animated="false" v-if="isReady">
            <b-tab-item value="base" label="Base Configuration">
                <edit-service-base v-model="model" :is-new="isNew" :field-messages="validations" />
            </b-tab-item>
            <b-tab-item value="extended" label="Extended Configuration" class="extended-configuration">
                <div class="extended-configuration__item">
                    <b-field label="Start Type">
                        <start-type-dropdown v-model="model.start_type"></start-type-dropdown>
                    </b-field>
                    <stop-signal v-model="model.stop_signal"></stop-signal>
                    <dependencies-list v-model="model.dependencies" :exclude="model.name"></dependencies-list>
                </div>
                <div class="extended-configuration__item">
                    <service-user :password.sync="model.password" :userName.sync="model.service_user"></service-user>
                </div>
            </b-tab-item>
            <b-tab-item value="recovery" label="Recovery Actions">
                <recovery-actions :recoveryActions.sync="model.recovery_actions"></recovery-actions>
            </b-tab-item>
        </b-tabs>
    </page>
</template>

<script>
import { isLoading, createMsg, isEqualObj } from '../util';
import { OverviewRoute } from '../routes';
import EditServiceBase from './EditServiceBase.vue';
import { createEditBaseModel } from './models/service';
import RecoveryActions from './recovery-action/RecoveryActions.vue';
import StartTypeDropdown from './extended-config/StartTypeDropdown.vue';
import StopSignal from './extended-config/StopSignal.vue';
import DependenciesList from './extended-config/DependenciesList.vue';
import ServiceUser from './extended-config/ServiceUser.vue';

const isValidServiceName = new RegExp(/^[a-zA-Z0-9_+\-!]+$/);

export default {
    components: { EditServiceBase, RecoveryActions, StartTypeDropdown, StopSignal, DependenciesList, ServiceUser },
    props: {
        serviceName: String,
    },
    data() {
        return {
            OverviewRoute,
            model: createEditBaseModel(),
            currentSvc: createEditBaseModel(),
            validations: this.serviceName
                ? {}
                : {
                      name: { msg: '* required', hasValidated: false },
                      exe_path: { msg: '* required', hasValidated: false },
                  },
            displayEdited: false,
            activeTab: 'base',
            isNew: !this.serviceName,
            isReady: !this.serviceName,
        };
    },
    watch: {
        'model.name': function (val) {
            if (!this.isNew) return;
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
            if (!this.isNew) return;
            if (this.isNameUpdating) {
                this.isNameUpdating = false;
                return;
            }

            if (!this.displayEdited && val) this.displayEdited = true;
            else if (this.displayEdited && !val) this.displayEdited = false;
        },
        'model.exe_path': function () {
            if (!this.isNew) return;
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
            if (this.serviceName) return this.serviceName;

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
        hasChanges: function () {
            return !isEqualObj(this.model, this.currentSvc);
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
        updateService: function () {        
            /* Todo: Implement Backend UpdateService method */
            isLoading(true);
            window.backend.Services.UpdateService(this.model)
                .then(() => {
                    this.$buefy.toast.open(createMsg('Successfully updated service.', 'success'));
                    this.$router.push({ name: OverviewRoute });
                })
                .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')))
                .finally(() => isLoading(false));
        },
    },
    created: function () {
        if (this.serviceName && !this.isNew) {
            isLoading(true);
            window.backend.Services.LoadService(this.serviceName)
                .then((s) => {
                    // Ensure reference model has the same properties
                    this.currentSvc = { ...createEditBaseModel(), ...s };
                    this.model = { ...createEditBaseModel(), ...s };
                    this.isReady = true;
                })
                .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')))
                .finally(() => isLoading(false));
        }
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
    align-content: flex-start;

    .extended-configuration__item {
        &:not(:last-child) {
            margin-right: 2rem;
            margin-bottom: 1rem;
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
                height: 200%;
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
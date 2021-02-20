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
                <b-button
                    size="is-small"
                    icon-left="cancel"
                    tag="router-link"
                    :to="{ name: OverviewRoute }"
                >
                    Cancel
                </b-button>
            </div>
        </template>
        <!-- Page content -->
        <b-tabs v-model="activeTab">
            <b-tab-item value="base" label="Base Configuration">
                <edit-service-base
                    v-model="model"
                    :is-new="true"
                    :field-messages="validations"
                />
            </b-tab-item>
            <b-tab-item value="extended" label="Extended Configuration">
                Here comes the extended configuration!
            </b-tab-item>
            <b-tab-item value="recovery" label="Recovery Actions">
                Here comes the recovery actions!
            </b-tab-item>
        </b-tabs>
    </page>
</template>

<script>
import { isLoading, createMsg } from '../util';
import { OverviewRoute } from '../routes';
import EditServiceBase from './EditServiceBase.vue';
import { createEditBaseModel } from './models/service';

const isValidServiceName = new RegExp(/^[a-zA-Z0-9_+\-!]+$/);

export default {
    components: { EditServiceBase },
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
            window.backend.Services.InstallService(this.model)
                .then(() => {
                    isLoading(false);
                    this.$buefy.toast.open(
                        createMsg('Successfully installed service.', 'success')
                    );
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
.b-tabs {
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
        padding: 16px 0;
    }
}
</style>
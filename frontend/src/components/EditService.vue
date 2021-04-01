<template>
    <page :title="model.name">
        <!-- Page Menu -->
        <template v-slot:menu>
            <div class="buttons">
                <b-button
                    size="is-small"
                    type="is-primary"
                    icon-left="content-save-cog"
                    @click="updateService"
                    :disabled="!hasChanges"
                >
                    Update
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
                <edit-service-base v-model="model" :is-new="false" />
            </b-tab-item>
            <b-tab-item value="extended" label="Extended Configuration">
                Here comes the extended configuration!
            </b-tab-item>
            <b-tab-item value="recovery" label="Recovery Actions">
                <recovery-actions></recovery-actions>
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

export default {
    components: { EditServiceBase, RecoveryActions },
    props: {
        svc: String,
    },
    data() {
        return {
            OverviewRoute,
            model: createEditBaseModel(),
            currentSvc: createEditBaseModel(),
        };
    },
    computed: {
        hasChanges: function () {
            return !isEqualObj(this.model, this.currentSvc);
        },
    },
    methods: {
        updateService: function () {
            this.$buefy.toast.open(
                createMsg("This feature isn't implemented yet!", 'info')
            );
            /* Todo: Implement Backend UpdateService method
            isLoading(true);
            window.backend.Services.UpdateService(this.model)
                .then(() => {
                    this.$buefy.toast.open(
                        createMsg('Successfully updated service.', 'success')
                    );
                    this.$router.push({ name: OverviewRoute });
                })
                .catch((reason) =>
                    this.$buefy.toast.open(createMsg(reason, 'error'))
                )
                .finally(() => isLoading(false));
                */
        },
    },
    created: function () {
        isLoading(true);
        window.backend.Services.LoadService(this.svc)
            .then((s) => {
                this.currentSvc = s;
                for (const key in this.model) {
                    this.model[key] = this.currentSvc[key];
                }
            })
            .catch((reason) =>
                this.$buefy.toast.open(createMsg(reason, 'error'))
            )
            .finally(() => isLoading(false));
    },
};
</script>

<style lang="scss" >
@import '../styles';
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
        padding: 16px 0;
        height: 100%;
        max-height: 100%;
        overflow-y: auto;

        &.is-transitioning {            
            max-height: 100%;
            height: 100%;

            .tab-item {
                height: 0;
            }
        }

        .tab-item {
            max-height: 100%;
            height: 100%;

            & > * {
                max-width: 100%;
            }
        }
    }
}
</style>
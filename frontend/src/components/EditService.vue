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
        <edit-service-base v-model="model" :is-new="false" />
    </page>
</template>

<script>
import { isLoading, createMsg, isEqualObj } from '../util';
import { OverviewRoute } from '../routes';
import EditServiceBase from './EditServiceBase.vue';
import { createEditBaseModel } from './models/service';

export default {
    components: { EditServiceBase },
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
        hasChanges: function() {
            return !isEqualObj(this.model, this.currentSvc);            
        }
    },
    methods: {
        updateService: function () {
            this.$buefy.toast.open(createMsg("This feature isn't implemented yet!", 'info'))
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

<style>
</style>
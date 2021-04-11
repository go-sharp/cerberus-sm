<template>
    <page title="Service Overview">
        <template v-slot:menu>
            <div class="buttons">
                <b-tooltip label="Reloads all configured services." position="is-left" size="is-small" :delay="1000">
                    <b-button @click="refreshServices" size="is-small" type="is-light" icon-left="refresh">
                        Refresh
                    </b-button>
                </b-tooltip>
            </div>
        </template>
        <section v-if="!loading && services.length === 0" class="hero is-primary is-small">
            <div class="hero-body">
                <div class="container">
                    <h2 class="title">No installed services found!</h2>
                    <h3 class="subtitle">Click on "New Service" and install your first cerberus service.</h3>
                </div>
            </div>
        </section>
        <b-table
            v-if="services.length !== 0"
            :data="services"
            default-sort-direction="asc"
            sort-icon="arrow-up"
            sort-icon-size="is-small"
            default-sort="name"
            striped
            narrowed
        >
            <b-table-column
                field="name"
                label="Name"
                header-class="is-sticky-column-one"
                cell-class="is-sticky-column-one"
                sortable
                v-slot="props"
            >
                {{ props.row.name }}
            </b-table-column>

            <b-table-column field="display_name" label="Display Name" sortable v-slot="props">
                {{ props.row.display_name }}
            </b-table-column>

            <b-table-column field="description" label="Description" v-slot="props">
                {{ props.row.description }}
            </b-table-column>

            <b-table-column field="state" label="State" sortable centered v-slot="props" cell-class="table-icon">
                <StateIcon :state="props.row.state" />
            </b-table-column>

            <b-table-column
                field="start_type"
                label="StartType"
                sortable
                centered
                v-slot="props"
                cell-class="table-icon"
            >
                <StartTypeIcon :startType="props.row.start_type" />
            </b-table-column>

            <b-table-column label="Actions" centered v-slot="props">
                <b-dropdown position="is-bottom-left" aria-role="list" :mobile-modal="false">
                    <!-- Dropdown Button/Trigger -->
                    <button class="button is-primary is-small" type="button" slot="trigger">
                        <b-icon size="is-small" icon="cog"></b-icon>
                    </button>

                    <!-- Menu Items -->
                    <b-dropdown-item
                        v-if="props.row.state === 1"
                        class="action-item"
                        aria-role="listitem"
                        @click="startService(props.row.name)"
                    >
                        <b-icon icon="play"></b-icon>
                        <span>Start</span>
                    </b-dropdown-item>
                    <b-dropdown-item
                        v-if="props.row.state === 4"
                        class="action-item"
                        aria-role="listitem"
                        @click="stopService(props.row.name)"
                    >
                        <b-icon icon="stop"></b-icon>
                        <span>Stop</span>
                    </b-dropdown-item>
                    <b-dropdown-item
                        v-if="props.row.state === 4"
                        class="action-item"
                        aria-role="listitem"
                        @click="restartService(props.row.name)"
                    >
                        <b-icon icon="restart"></b-icon>
                        <span>Restart</span>
                    </b-dropdown-item>
                    <b-dropdown-item class="action-item" aria-role="listitem" @click="editService(props.row.name)">
                        <b-icon icon="pencil"></b-icon>
                        <span>Edit</span>
                    </b-dropdown-item>
                    <b-dropdown-item class="action-item" aria-role="listitem" @click="deleteService(props.row.name)">
                        <b-icon type="is-danger" icon="delete"></b-icon>
                        <span>Delete</span>
                    </b-dropdown-item>
                </b-dropdown>
            </b-table-column>
        </b-table>
    </page>
</template>

<script>
import { isLoading, createMsg } from '../util';
import StateIcon from './StateIcon.vue';
import StartTypeIcon from './extended-config/StartTypeIcon.vue';
import { EditRoute } from '../routes';

export default {
    components: {
        StateIcon,
        StartTypeIcon,
    },
    data() {
        return {
            services: [],
            loading: false,
        };
    },
    methods: {
        isLoading: function (loading) {
            this.loading = loading;
            isLoading(loading);
        },
        refreshServices: function () {
            this.isLoading(true)
            window.backend.Services.LoadOverviewServices()
                .then((svc) => {
                    this.services = svc;
                })
                .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')))
                .finally(() => this.isLoading(false));
        },
        startService: function (svc) {
            this.runBackgroundTask(svc, 'StartService', 'started', 'start');
        },
        stopService: function (svc) {
            this.runBackgroundTask(svc, 'StopService', 'stopped', 'stop');
        },
        restartService: function (svc) {
            this.runBackgroundTask(svc, 'RestartService', 'restarted', 'restart');
        },
        editService: function (svc) {
            this.$router.push({ name: EditRoute, params: { serviceName: svc } });
        },
        deleteService: function (svc) {
            this.runBackgroundTask(svc, 'DeleteService', 'deleted', 'delete');
        },
        runBackgroundTask: function (svc, method, successMsg, errorMsg) {
            this.isLoading(true);
            window.backend.Services[method](svc)
                .then(() => {
                    this.$buefy.toast.open(createMsg(`Successfully ${successMsg} service '${svc}'`, 'success'));
                })
                .catch((reason) => {
                    this.$buefy.toast.open(
                        createMsg(`Failed to ${errorMsg} service '${svc}': ${reason}`, 'error', 5000)
                    );
                })
                .finally(() => {
                    this.isLoading(false);
                    this.refreshServices();
                });
        },
    },
    // Lifecycle hooks
    created: function () {
        this.refreshServices();
    },
};
</script>

<style lang="scss">
@import '../styles';

.action-item {
    display: flex;
    align-items: center;

    span:nth-child(0n + 2) {
        margin-left: 15px;
    }
}

.table-icon {
    color: $primary;
}

@media (max-width: 768px) {
    .b-table .table-wrapper.has-mobile-cards tr {
        box-shadow: 0 0 0.2em 1px $background-contrast, 0 0px 0 1px $background-contrast;
    }
}

@media screen and (max-width: 1023px) {
    .table-wrapper {
        overflow: visible;
    }
}
</style>
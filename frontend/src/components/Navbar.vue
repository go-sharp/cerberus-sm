<template>
    <div class="level is-mobile">
        <div class="level-left">
            <div class="level-item">
                <h1 class="title">Cerberus</h1>
                <span>Service Manager {{ version }}</span>
            </div>
        </div>
        <div class="level-right">
            <div class="level-item">
                <b-button
                    type="is-primary is-small"
                    icon-left="home"
                    tag="router-link"
                    :to="{ name: this.OverviewRoute }"
                    :disabled="!isAdmin"
                >
                    Services
                </b-button>
                <!-- <router-link :to="{ name: this.OverviewRoute }">Overview</router-link> -->
            </div>
            <div class="level-item">
                <b-button
                    type="is-primary is-small"
                    icon-left="plus"
                    tag="router-link"
                    :disabled="!isAdmin"
                    :to="{ name: this.AddRoute }"
                >
                    New Service
                </b-button>
                <!-- <router-link :to="{ name: this.AddRoute }">Add Service</router-link> -->
            </div>
        </div>
    </div>
</template>

<script>
import { OverviewRoute, AddRoute } from '../routes';

export default {
    props: {
        version: {
            type: String,
            default: ''
        },
        isAdmin: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            OverviewRoute,
            AddRoute
        };
    },
};
</script>

<style scoped lang="scss">
@import '../styles';

.level {
    padding: 0.5rem;
}

.title {
    & + span {
        color: $background-contrast;
        margin-left: 5px;
    }
}
</style>

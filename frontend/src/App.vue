<template>
    <div id="app">
        <navbar id="navbar" :isAdmin="isAdmin" :version="version"></navbar>
        <section v-if="!isAdmin" class="hero is-danger is-small">
            <div class="hero-body">
                <div class="container">
                    <h2 class="title">Need administrator rights!</h2>
                    <h3 class="subtitle">Start Cerberus Service Manager with administration rights.</h3>
                </div>
            </div>
        </section>
        <div v-if="isAdmin" class="main">
            <router-view></router-view>
        </div>
        <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
    </div>
</template>

<script>
import './assets/css/main.css';
import Navbar from './components/Navbar.vue';
import { LOADING_EVENT } from './util';

export default {
    name: 'app',
    components: {
        navbar: Navbar,
    },
    data() {
        return {
            isLoading: false,
            isAdmin: false,
            version: ''
        };
    },
    // Lifecylce hooks
    created: function () {
        window.wails.Events.On(LOADING_EVENT, (isLoading) => (this.isLoading = !!isLoading));

        window.backend
            .Version()
            .then((v) => (this.version = v))
            .catch((reason) => console.error(reason));

        window.backend.Services.IsAdmin()
            .then((isAdmin) => (this.isAdmin = isAdmin))
            .catch((reason) => console.error(reason));
    },
};
</script>

<style lang="scss">
// Import custom colors and styles
@import 'styles';

@import '~bulma';
@import '~buefy/src/scss/buefy';

.main {
    height: 100%;
    background-color: $background-color;
    overflow: hidden;
}

#navbar {
    margin-bottom: 2rem;
}

.title {
    color: $primary;
}

.label,
.help {
    color: $background-contrast;
}

.main .b-table .table thead tr th {
    color: $background-contrast;
}

/* Input */
.input,
.textarea,
.taginput .taginput-container.is-focusable,
.select select {
    background-color: $background-color;
    color: $background-contrast;
    border-color: $background-contrast;
}

.autocomplete .dropdown-item.is-hovered {
    background-color: $primary;
    color: $white;
    &:hover {
        color: $black;
    }
}

[disabled].input,
[disabled].textarea,
.taginput [disabled].taginput-container.is-focusable,
.select select[disabled],
fieldset[disabled] .input,
fieldset[disabled] .textarea,
fieldset[disabled] .taginput .taginput-container.is-focusable,
.taginput fieldset[disabled] .taginput-container.is-focusable,
fieldset[disabled] .select select,
.select fieldset[disabled] select {
    background-color: $input-background-color;
    border-color: $background-contrast;
}

.control.has-icons-left .input:focus ~ .icon,
.control.has-icons-left .select:focus ~ .icon,
.control.has-icons-right .input:focus ~ .icon,
.control.has-icons-right .select:focus ~ .icon {
    color: $primary;
}

.taginput.control .autocomplete.control {
    overflow-x: hidden;
}

/* Buefy dropdown menu */
.dropdown-content {
    background-color: $background-color;
    border-collapse: separate;
    box-shadow: 0 0.2em 1.5em -0.125em #a2a2a2, 0 0px 0 1px #a2a2a2;

    .dropdown-item {
        color: $background-contrast;

        &:hover {
            background-color: $background-contrast;
            .media {
                h3 {
                    color: $black;
                }
            }
        }

        &.is-active {
            &:hover {
                .media {
                    color: $black;
                    h3 {
                        color: $black;
                    }
                }
            }
            .media {
                h3 {
                    color: $white;
                }
            }
        }

        .media {
            h3 {
                color: $background-contrast;
                margin-bottom: 0;
            }
        }
    }
}

.control-label:hover {
    color: $foreground-color;
}

.checkbox:hover {
    color: $foreground-color;
}
</style>
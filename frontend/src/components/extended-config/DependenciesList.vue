<template>
    <b-field label="Service Dependencies" message="Start typing to get service suggestions">
        <b-taginput
            v-model="dependencies"
            :data="filteredServices"
            autocomplete
            :allow-new="false"
            :open-on-focus="false"
            :keep-first="false"
            :clear-on-select="true"
            field="display_name"
            icon="cogs"
            placeholder="Add a dependency"
            @typing="getFilteredServices"
            @select="getFilteredServices('')"
            type="is-primary"
        >
        </b-taginput>
    </b-field>
</template>

<script>
import { createMsg } from '../../util';
export default {
    props: {
        value: Array,
        exclude: [String, Array],
    },
    data: function () {
        return {
            dependencies: [],
            availableSvcs: [],
            filteredServices: [],
        };
    },
    methods: {
        getFilteredServices: function (text) {
            const regex = new RegExp(text, 'i');
            this.filteredServices = this.availableSvcs.filter((s) => {
                if (this.dependencies.indexOf(s) >= 0) return false;

                if (text && !s.name.match(regex)) {
                    return false;
                }

                return true;
            });
        },
    },
    watch: {
        dependencies: function (current) {
            this.$emit(
                'input',
                current.map((s) => s.name)
            );
        },
    },
    mounted: function () {
        window.backend.Services.GetDependOnSvc()
            .then((svcs) => {
                if (this.exclude) {
                    this.availableSvcs = svcs.filter((e) => {
                        if(typeof this.exclude === 'string' && e.name === this.exclude) return false;
                        if(Array.isArray(this.exclude) && this.exclude.indexOf(e.name) >= 0) return false;

                        return true;
                    });
                } else this.availableSvcs = svcs;

                if (this.value) {
                    for (const svc of this.value) {
                        const s = this.availableSvcs.find((e) => e.name === svc);
                        if (s) {
                            this.dependencies.push(s);
                        }
                    }
                }
                this.getFilteredServices('');
            })
            .catch((reason) => this.$buefy.toast.open(createMsg(reason, 'error')));
    },
};
</script>

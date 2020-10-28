<template>
    <section class="columns">
        <div class="column is-one-half">
            <!-- Name field -->
            <b-field
                label="Name:"
                :message="isNew ? messages.name.msg : ''"
                :type="messages.name.type"
            >
                <b-input
                    v-model.trim="model.name"
                    :aria-disabled="!isNew"
                    :disabled="!isNew"
                ></b-input>
            </b-field>
            <!-- Display Name field -->
            <b-field
                label="Display Name:"
                :message="messages.display_name.msg"
                :type="messages.display_name.type"
            >
                <b-input v-model.trim="model.display_name"></b-input>
            </b-field>
            <!-- Description field -->
            <b-field
                label="Description:"
                :message="messages.description.msg"
                :type="messages.description.type"
            >
                <b-input
                    maxlength="200"
                    type="textarea"
                    v-model="model.description"
                ></b-input>
            </b-field>
        </div>
        <div class="column is-one-half">
            <!-- Executable Path field -->
            <b-field
                label="Executable Path:"
                :message="isNew ? messages.exe_path.msg : ''"
                :type="messages.exe_path.type"
            >
                <b-input
                    v-model="model.exe_path"
                    expanded
                    :aria-disabled="!isNew"
                    :disabled="!isNew"
                ></b-input>
                <p class="control">
                    <button
                        class="button is-primary"
                        @click="showDialog()"
                        :aria-disabled="!isNew"
                        :disabled="!isNew"
                    >
                        Select
                    </button>
                </p>
            </b-field>
            <!-- Working Directory field -->
            <b-field
                label="Working Directory:"
                :message="messages.work_dir.msg"
                :type="messages.work_dir.type"
            >
                <b-input v-model="model.work_dir" expanded></b-input>
                <p class="control">
                    <button class="button is-primary" @click="showDialog(true)">
                        Select
                    </button>
                </p>
            </b-field>
            <!-- Arguments field -->
            <b-field
                label="Arguments:"
                :message="messages.args.msg"
                :type="messages.args.type"
            >
                <b-taginput
                    v-model="model.args"
                    icon="format-text-variant"
                    ellipsis
                    allow-duplicates
                    placeholder="Add"
                >
                </b-taginput>
            </b-field>
            <!-- Envrionment field -->
            <b-field
                label="Environment:"
                :message="messages.env.msg"
                :type="messages.env.type"
            >
                <b-taginput
                    v-model="model.env"
                    ellipsis
                    icon="currency-usd"
                    allow-duplicates
                    placeholder="Add"
                >
                </b-taginput>
            </b-field>
        </div>
    </section>
</template>

<script>
import { createMsg } from '../util';

const initMessages = {
    name: { msg: '* required', type: '' },
    display_name: { msg: '', type: '' },
    description: { msg: '', type: '' },
    exe_path: { msg: '* required', type: '' },
    work_dir: { msg: '', type: '' },
    args: { msg: '', type: '' },
    env: { msg: '', type: '' },
};

export default {
    data: function () {
        return {
            model: {
                name: '',
                display_name: '',
                description: '',
                exe_path: '',
                work_dir: '',
                args: [],
                env: [],
            },
        };
    },
    computed: {
        messages: function () {
            let msgs = this.fieldMessages || {};
            for (const key in initMessages) {
                msgs[key] = msgs[key]
                    ? { ...initMessages[key], ...msgs[key] }
                    : initMessages[key];
            }

            return msgs;
        },
    },
    watch: {
        model: {
            handler(val) {
                this.$emit('input', val);
            },
            deep: true,
        },
    },
    props: {
        value: Object,
        isNew: {
            type: Boolean,
            default: true,
        },
        fieldMessages: {
            type: Object,
        },
    },
    methods: {
        showDialog: function (isWorkDir) {
            window.backend.Services.ShowFileDialog({
                dir: !!isWorkDir,
            })
                .then((val) => {
                    const prop = isWorkDir ? 'work_dir' : 'exe_path';
                    this.model[prop] = val;
                })
                .catch((reason) =>
                    this.$buefy.toast.open(createMsg(reason, 'error'))
                );
        },
    },
    // Lifecycle hooks
    created: function () {
        this.model = { ...this.model, ...this.value };
    },
};
</script>
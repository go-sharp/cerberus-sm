<template>
  <section class="columns">
    <div class="column is-one-half">
      <b-field label="Name:" :message="isNew ? '* required' : ''">
        <b-input
          v-model="model.name"
          required
          :aria-disabled="!isNew"
          :disabled="!isNew"
        ></b-input>
      </b-field>
      <b-field label="Display Name:">
        <b-input v-model="model.display_name" @input="test"></b-input>
      </b-field>
      <b-field label="Description:">
        <b-input maxlength="200" type="textarea"></b-input>
      </b-field>
    </div>
    <div class="column is-one-half">
      <b-field
        label="Executable Path:"
        :message="isNew ? '* required' : ''"
      >
        <b-input
          v-model.trim="model.exe_path"
          @input="exeError = model.exe_path === '' ? 'is-danger' : ''"
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
      <b-field label="Working Directory:">
        <b-input v-model="model.work_dir" expanded></b-input>
        <p class="control">
          <button class="button is-primary" @click="showDialog(true)">
            Select
          </button>
        </p>
      </b-field>
      <b-field label="Arguments:">
        <b-taginput
          v-model="model.args"
          icon="format-text-variant"
          ellipsis
          allow-duplicates
          placeholder="Add"
        >
        </b-taginput>
      </b-field>
      <b-field label="Environment:">
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
import { createMsg } from "../util";

export default {
  data: function () {
    return {
      model: {
        name: "",
        display_name: "",
        description: "",
        exe_path: "",
        work_dir: "",
        args: [],
        env: [],
      },
    };
  },
  props: {
    value: Object,
    isNew: {
      type: Boolean,
      default: true,
    },
  },
  methods: {
    test(event) {
      window.alert(JSON.stringify(event.target));
      //event.target.checkValidity();
      //window.alert(JSON.stringify(arguments));
    },
    showDialog: function (isWorkDir) {
      window.backend.Services.ShowFileDialog({
        dir: !!isWorkDir,
      })
        .then((val) => {
          const prop = isWorkDir ? "work_dir" : "exe_path";
          this.model[prop] = val;
        })
        .catch((reason) => this.$buefy.toast.open(createMsg(reason, "error")));
    },
  },
  // Lifecycle hooks
  created: function () {
    this.model = { ...this.model, ...this.value };
  },
};
</script>

<style>
</style>
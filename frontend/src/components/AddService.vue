<template>
  <page title="New Service">
    <!-- Page Menu -->
    <template v-slot:menu>
      <div class="buttons">
        <b-button
          size="is-small"
          type="is-primary"
          icon-left="content-save-cog"
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
    <edit-service-base v-model="m" @input="(event) => alert(event)" :is-new="false"/>
  </page>
</template>

<script>
import { OverviewRoute } from "../routes";
import EditServiceBase from "./EditServiceBase.vue";

export default {
  components: { EditServiceBase },
  data() {
    return {
      OverviewRoute,
      m: {name: 'hello', args: ['asd']}
    };
  },
};
</script>

<style>
</style>
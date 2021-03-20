<template>
    <section class="columns">
        <div class="column is-one-half">
              <recovery-action-editor
         
            ></recovery-action-editor>

        </div>
        <div class="column is-one-half">
            <recovery-action-editor
                :isNew="true"
                :excludeErrorCodes="[11, 23]"
                v-model="action"
                @validationChanged="abortEditing"
            ></recovery-action-editor>
        </div>
    </section>
</template>

<script>
import RecoveryActionEditor from './RecoveryActionEditor.vue';

export default {
    components: { RecoveryActionEditor },
    data() {
        return {
            action: { exitCode: 45, delay: 444 },            
        };
    },
    methods: {
        abortEditing: function(evt) {
            console.log('abort clicked...', evt);
        }
    }    
};
</script>

<style>
</style>
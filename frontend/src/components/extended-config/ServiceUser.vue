<template>
    <section>
        <b-field label="Service User"> </b-field>
        <b-field>
            <b-checkbox v-model="useSystem"> Use Local System </b-checkbox>
        </b-field>
        <b-field
            label="User"
            custom-class="is-small"
            :type="userHasError"
            :message="userHasError ? 'User name is required' : ''"
        >
            <b-input v-model="name" :disabled="useSystem"></b-input>
        </b-field>
        <b-field label="Password" message="Leave empty to use existing password" custom-class="is-small">
            <b-input type="password" v-model="pwd" :password-reveal="false" :disabled="useSystem"></b-input>
        </b-field>
    </section>
</template>

<script>
export default {
    props: {
        userName: String,
        password: String,
    },
    data: function () {
        return {
            pwd: this.password,
            name: this.userName,
            useSystem: this.userName === 'LocalSystem',
        };
    },
    computed: {
        userHasError: function () {
            return this.name && this.name.length > 0 ? '' : 'is-danger';
        },
    },
    watch: {
        useSystem: function (current) {
            if (current) {
                this.name = 'LocalSystem';
                this.pwd = '';
            }
        },
        pwd: function (current) {
            this.$emit('update:password', current);
        },
        name: function (current) {
            this.$emit('update:userName', current);
        },
    },
};
</script>

<style>
</style>
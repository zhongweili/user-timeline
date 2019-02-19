<template>
    <v-toolbar dark color="teal">
      <v-spacer></v-spacer>
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="Search for users on Github"
        prepend-inner-icon="search"
        v-model="query"
        @keyup.enter="onSearchSubmition"
      ></v-text-field>
      <v-spacer></v-spacer>
      <div  v-if="loginStatus" >
        <button @click.prevent="logout">Logout</button>
      </div>
      <div v-else >
        <button @click.prevent="login">Login</button>
      </div>
    </v-toolbar>
</template>

<script>
export default {
    data() {
      return {
        query: null,
      };
    },
    props: ['defaultQuery', 'loginStatus'],
    methods: {
      onSearchSubmition() {
        this.$emit('search-submitted', this.query);
      },
      async logout () {
        await this.$auth.logout()
        this.$router.push('/')
      },
      async login () {
        const isAuthenticated = await this.$auth.isAuthenticated();
        isAuthenticated && this.$router.push('/me');
        this.$auth.loginRedirect('/me')
      }
    }
}
</script>

<template>
  <div>
    <SearchBar :loginStatus='0' defaultQuery='okta' v-on:search-submitted="githubQuery" />
    <v-container grid-list-md fluid class="grey lighten-4" >
        <v-tabs-items style="width:100%" v-model="tabs">
          <v-tab-item :key="1">
            <v-layout row wrap>
              <v-flex v-for="user in users" :key="user.id" md4>
                <GitHubUser :user="user" />
              </v-flex>
            </v-layout>
          </v-tab-item>
        </v-tabs-items>
    </v-container>
  </div>
</template>

<script>
import SearchBar from './SearchBar.vue'
import GitHubUser from './GithubUser.vue'
import githubClient from '../githubClient'

export default {
  name: 'Home',
  components: { SearchBar, GitHubUser },
  data() {
    return {
      users: [],
      tabs: 0
    }
  },
  methods: {
    githubQuery(query) {
      this.tabs = 1;
      githubClient
        .getJSONUsers(query)
        .then((response) => {
          this.users = response.items
        })
    },
  },
}
</script>

<style>
 .v-tabs__content {
   padding-bottom: 2px;
 }
</style>

<template>
  <div>
    <SearchBar :loginStatus='1' defaultQuery='okta' v-on:search-submitted="githubQuery" />
    <v-container grid-list-md fluid class="grey lighten-4" >
        <v-tabs
        slot="extension"
        v-model="tabs"
        centered
        color="teal"
        text-color="white"
        slider-color="white"
        >
        <v-tab class="white--text" :key="2">
          FOLLOWINGS
        </v-tab>
        <v-tab class="white--text" :key="1">
          USERS
        </v-tab>
      </v-tabs>
        <v-tabs-items style="width:100%" v-model="tabs">
          <v-tab-item :key="2">
            <v-layout row wrap>
              <v-flex v-for="following in allFollowings" :key="following.id" md4 >
                <GitHubUser :user="following" />
              </v-flex>
            </v-layout>
          </v-tab-item>
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
import { mapMutations, mapGetters, mapActions } from 'vuex'

export default {
  name: 'Me',
  components: { SearchBar, GitHubUser },
  data() {
    return {
      tabs: 0
    }
  },
  computed: mapGetters(['allFollowings', 'users']),
  created() {
    this.getFollowings();
  },
  methods: {
    githubQuery(query) {
      this.tabs = 1;
      githubClient
        .getJSONUsers(query)
        .then(response => this.resetUsers(response.items) )
    },
    ...mapMutations(['resetUsers']),
    ...mapActions(['getFollowings']),
  },
}
</script>

<style>
 .v-tabs__content {
   padding-bottom: 2px;
 }
</style>

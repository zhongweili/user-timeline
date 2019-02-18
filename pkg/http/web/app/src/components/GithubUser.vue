<template>
  <v-card >
    <v-card-title primary-title>
      <div class="user-card-content">
        <h3 class="headline mb-0">
          <a :href="'https://github.com/' + user.login">{{user.login}}</a>
        </h3>
      </div>
      <div class="box">
            <img v-bind:src="user.avatar_url"/>
        </div>
    </v-card-title>
    <v-card-actions>
      <v-chip>
        <router-link :to="{ name: 'timeline-details', params: { name: user.login }}" >Timeline</router-link>
      </v-chip>
      <v-spacer></v-spacer>
      <v-btn @click.prevent="toggleFollowing(user)"  flat icon color="pink">
        <v-icon v-if="isFollowing(user)">favorite</v-icon>
        <v-icon v-else>favorite_border</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  data() {
    return {}
  },
  props: ['user'],
  methods: {
    isFollowing(user) {
      return this.$store.getters.isFollowing(user);
    },
    ...mapActions(['toggleFollowing'])
  }
}
</script>

<style>
 .user-card-content {
   height: 90px;
   overflow: scroll;
 }
 .box{
   position: absolute;
   right: 5%;
   padding-top: 25%;
   overflow: hidden;
   width: 25%;
 }
 .box img{
   position: absolute;
   top: 0;
   width: 100%;
   height: 100%;
 }
</style>
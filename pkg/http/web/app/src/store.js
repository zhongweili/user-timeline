import Vue from 'vue';
import Vuex from 'vuex';

import APIClient from './apiClient';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    followings: {},
    users: [],
  },
  mutations: {
    resetUsers (state, users) {
        state.users = users;
      },
    resetFollowings(state, followings) {
      state.followings = followings;
    }
  },
  getters: {
    allFollowings(state) {
      return Object.values(state.followings);
    },
    followings(state) {
      return state.followings;
    },
    users(state) {
        return state.users;
    },
    isFollowing(state) {
      return (user)=> {
        return !!state.followings[user.id];
      }
    }
  },
  actions: {
    getFollowings ({commit}) {
      APIClient.getFollowings().then((data) => {
        commit('resetFollowings', data.reduce((acc, following) => {
                               return {[following.id]: following, ...acc}
                             }, {}))
      })
    },
    updateFollowing({ commit, state }, user) {
      const followings = { ...state.following, [user.id]: user };

      return APIClient
        .updateFollowing(user)
        .then(() => {
          commit('resetFollowings', followings)
        });
    },
    toggleFollowing({ commit, state }, user) {
      if (!state.followings[user.id]) {
        return APIClient
          .createFollowing(user)
          .then(following => commit('resetFollowings', { [following.id]: following, ...state.followings }))
      }

      const followings = Object.entries(state.followings).reduce((acc, [id, following]) => {
                      return (id == user.id) ? acc
                                                 : { [id]: following, ...acc };
                    }, {});

      return APIClient
        .deleteFollowing(user)
        .then(() => commit('resetFollowings', followings));
    }
  }
});

export default store;

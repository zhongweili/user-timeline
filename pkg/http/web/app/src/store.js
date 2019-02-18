import Vue from 'vue';
import Vuex from 'vuex';

import APIClient from './apiClient';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    kudos: {},
    repos: [],
    followings: {},
    users: [],
  },
  mutations: {
    resetRepos (state, repos) {
      state.repos = repos;
    },
    resetUsers (state, users) {
        state.users = users;
      },
    resetKudos(state, kudos) {
      state.kudos = kudos;
    },
    resetFollowings(state, followings) {
      state.followings = followings;
    }
  },
  getters: {
    allKudos(state) {
      return Object.values(state.kudos);
    },
    allFollowings(state) {
      return Object.values(state.followings);
    },
    followings(state) {
      return state.followings;
    },
    kudos(state) {
      return state.kudos;
    },
    repos(state) {
      return state.repos;
    },
    users(state) {
        return state.users;
      },
    isKudo(state) {
      return (repo)=> {
        return !!state.kudos[repo.id];
      };
    },
    isFollowing(state) {
      return (user)=> {
        return !!state.followings[user.id];
      }
    }
  },
  actions: {
    getKudos ({commit}) {
      APIClient.getKudos().then((data) => {
        commit('resetKudos', data.reduce((acc, kudo) => {
                               return {[kudo.id]: kudo, ...acc}
                             }, {}))
      })
    },
    updateKudo({ commit, state }, repo) {
      const kudos = { ...state.kudos, [repo.id]: repo };

      return APIClient
        .updateKudo(repo)
        .then(() => {
          commit('resetKudos', kudos)
        });
    },
    toggleKudo({ commit, state }, repo) {
      if (!state.kudos[repo.id]) {
        return APIClient
          .createKudo(repo)
          .then(kudo => commit('resetKudos', { [kudo.id]: kudo, ...state.kudos }))
      }

      const kudos = Object.entries(state.kudos).reduce((acc, [repoId, kudo]) => {
                      return (repoId == repo.id) ? acc
                                                 : { [repoId]: kudo, ...acc };
                    }, {});

      return APIClient
        .deleteKudo(repo)
        .then(() => commit('resetKudos', kudos));
    },
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

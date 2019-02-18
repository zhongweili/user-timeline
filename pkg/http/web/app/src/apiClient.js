import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'http://localhost:4444';
const client = axios.create({
  baseURL: BASE_URI,
  json: true
});

const APIClient =  {
  createKudo(repo) {
    return this.perform('post', '/kudos', repo);
  },

  deleteKudo(repo) {
    return this.perform('delete', `/kudos/${repo.id}`);
  },

  updateKudo(repo) {
    return this.perform('put', `/kudos/${repo.id}`, repo);
  },

  getKudos() {
    return this.perform('get', '/kudos');
  },

  getKudo(repo) {
    return this.perform('get', `/kudo/${repo.id}`);
  },

  createFollowing(user) {
    return this.perform('post', '/kudos', user);
  },

  deleteFollowing(user) {
    return this.perform('delete', `/kudos/${user.id}`);
  },

  updateFollowing(user) {
    return this.perform('put', `/kudos/${user.id}`, user);
  },

  getFollowings() {
    return this.perform('get', '/kudos');
  },

  getFollowing(user) {
    return this.perform('get', `/kudo/${user.id}`);
  },

  async perform (method, resource, data) {
    let accessToken = await Vue.prototype.$auth.getAccessToken()
    return client({
      method,
      url: resource,
      data,
      headers: {
        Authorization: `Bearer ${accessToken}`
      }
    }).then(req => {
      return req.data
    })
  }
}

export default APIClient;

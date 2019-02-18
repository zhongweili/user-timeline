import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'http://localhost:4444';
const client = axios.create({
  baseURL: BASE_URI,
  json: true
});

const APIClient =  {
  createFollowing(user) {
    return this.perform('post', '/followings', user);
  },

  deleteFollowing(user) {
    return this.perform('delete', `/followings/${user.id}`);
  },

  updateFollowing(user) {
    return this.perform('put', `/followings/${user.id}`, user);
  },

  getFollowings() {
    return this.perform('get', '/followings');
  },

  getFollowing(user) {
    return this.perform('get', `/following/${user.id}`);
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

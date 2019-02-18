const API_URL = "https://api.github.com/search/users"
export default {
  getJSONUsers(query) {
    return fetch(`${API_URL}?q=` + query).then(response => response.json());
  }
}

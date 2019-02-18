const API_URL1 = "https://api.github.com/search/repositories"
const API_URL2 = "https://api.github.com/search/users"
export default {
  getJSONRepos(query) {
    return fetch(`${API_URL1}?q=` + query).then(response => response.json());
  },
  getJSONUsers(query) {
    return fetch(`${API_URL2}?q=` + query).then(response => response.json());
  }
}

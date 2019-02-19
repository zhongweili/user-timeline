import Vue from 'vue';
import VueRouter from 'vue-router';
import Auth from '@okta/okta-vue'
import Home from './components/Home';
import Me from './components/Me';
import TimelineDetails from './components/TimelineDetails';

Vue.use(VueRouter);
Vue.use(Auth, {
  issuer: 'https://dev-179387.okta.com/oauth2/default',
  client_id: '0oa9x42szoVX3cGOu356',
  redirect_uri: 'http://localhost:8080/implicit/callback',
  scope: 'openid profile email'
})

export default new VueRouter({
 mode: 'history',
 routes: [
   { path: '/me', component: Me, meta: { requiresAuth: true }},
   { name: 'home', path: '/', component: Home },
   { name: 'timeline-details', path: '/timeline/:name', component: TimelineDetails },
   { path: '/implicit/callback', component: Auth.handleCallback() }
 ]
});

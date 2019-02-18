<template>
    <div class="my-container">
        <h2 style="margin-bottom:20px" class="align-items-center timeline">{{name}}'s Timeline</h2>
        <div v-for="element in event" :key="element.id" >
            <div class="row item-row">
                <v-layout row class="px-4">
                    <v-flex xs9>
                        
                        <img style="float:left" v-bind:src="element.actor.avatar_url">
                    
                        <span v-if="element.type=='ForkEvent'">
                            <a class="link" :href="'https://github.com/' + element.actor.display_login" target="_blank">{{element.actor.display_login}}</a>
                                forked <a class="link" :href="element.payload.forkee.html_url" target="_blank">{{element.payload.forkee.full_name}}</a> from 
                            <a class="link" :href="'https://github.com/' + element.repo.name" target="_blank">{{element.repo.name}}</a>
                        </span>
                        <span v-else>
                            <a class="link" :href="'https://github.com/' + element.actor.display_login" target="_blank">{{element.actor.display_login}}</a>
                                {{getEventType(element.type, element.payload)}}
                            <a class="link" :href="'https://github.com/' + element.repo.name" target="_blank">{{element.repo.name}}</a>
                        </span>

                    </v-flex>

                    <v-flex xs3 align-self-center>
                        <span class="time"> {{from(element.created_at)}}</span>
                    </v-flex>
                </v-layout>
            </div>
        </div>
    </div>
</template>

<script>
import dayjs from 'dayjs'
import relativeTime from "dayjs/plugin/relativeTime"

export default {
  data() {
    return {
      name: String,
      event: {},
      eventType: String
    }
  },
  watch: {
    '$route': 'fetchData'
  },
  created() {
    dayjs.extend(relativeTime)
    this.fetchData();
  },
  methods: {
    fetchData() {
      fetch('https://api.github.com/users/' + this.$route.params.name + '/received_events')
        .then(response => response.json())
        .then((response) => {
          this.event = response
          this.name = this.$route.params.name
        })
    },
    getEventType(type, payload) {
      let event = ''
      switch(type) {
        case 'WatchEvent':
            event = 'starred'
            break
        case 'ForkEvent':
            event = `forked <a href="${payload.forkee.html_url}" class="link">${payload.forkee.full_name}</a> from `
            break
        case 'PublicEvent':
            event = 'made public'
            break
        case 'CreateEvent':
            event = 'created a repository'
            break
        case 'PullRequestEvent':
            event = 'opened a pull request in'
            break
        case 'PushEvent':
            event = 'pushed a commit'
            break;
        default:
            event = ''
        }
        return event
    },
    from(time) {
        let timeDisplay = ''
        timeDisplay = dayjs(time).fromNow()
        return timeDisplay
    }

  }
}
</script>

<style>
.my-container {
    margin: 5%
}

.loader {
    border: 5px solid #f3f3f3; /* Light grey */
    border-top: 5px solid #3498db; /* Blue */
    border-radius: 50%;
    width: 3em;
    height: 3em;
    animation: spin 3s linear infinite;
    display: none;
}

@keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}

.link	{
    color: #000;
    font-weight: bold;
}

.time {
    float: right;
    color: #888;
}

img {
    width: 50px;
    height: 50px;
    margin-right: 10px;
    border-radius: 5px;
}

span {
    font-size: 1.15rem;
}

.item-row {
    border-bottom: 1px solid #ccc;
    padding: 10px 0px;
    width:100%
}
</style>

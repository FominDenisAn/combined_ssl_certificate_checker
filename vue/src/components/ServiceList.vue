<template>
  <div class="service-list">
    <h3>Service List</h3>
    <ul>
      <li v-for="(hosts, serviceName) in services" :key="serviceName">
        <strong>{{ serviceName }}</strong>
        <ul>
          <li v-for="host in hosts" :key="host">
            <a href="#" @click.prevent="selectHost(host)">{{ host }}</a>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      services: {},
    };
  },
  created() {
    this.fetchServices();
  },
  methods: {
    async fetchServices() {
      try {
        const response = await axios.get('http://localhost:8080/services');
        this.services = response.data;
      } catch (error) {
        console.error('Error fetching services:', error);
      }
    },
    selectHost(host) {
      axios.get(`http://localhost:8080/certificate?host=${host}`)
        .then(response => {
          this.$emit('host-selected', response.data);
        })
        .catch(error => {
          console.error('Error fetching certificate info:', error);
        });
    },
  },
};
</script>

<style scoped>
/* Add your styles here */
</style>

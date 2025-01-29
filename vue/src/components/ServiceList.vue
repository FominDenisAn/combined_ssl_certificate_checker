<template>
  <div class="service-list">
    <h3>Service List</h3>
    <ul>
      <li v-for="(hosts, serviceName) in services" :key="serviceName">
        <button @click="toggleService(serviceName)">
          {{ serviceName }}
          <span v-if="expandedServices.includes(serviceName)">-</span>
          <span v-else>+</span>
        </button>
        <ul v-if="expandedServices.includes(serviceName)">
          <li v-for="host in hosts" :key="host">
            <a href="#" @click.prevent="selectHost(host)">
              {{ host }} 
              <Indicator :value="cpuUsage[host]" />
              <Indicator :value="ramUsage[host]" />
            </a>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import Indicator from './Indicator.vue';

export default {
  components: { Indicator },
  data() {
    return {
      services: {},
      expandedServices: [],
      cpuUsage: {},
      ramUsage: {},
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
    toggleService(serviceName) {
      if (this.expandedServices.includes(serviceName)) {
        this.expandedServices = this.expandedServices.filter(s => s !== serviceName);
      } else {
        this.expandedServices.push(serviceName);
      }
    },
    async selectHost(host) {
      try {
        const response = await axios.get(`http://localhost:8080/certificate?host=${host}`);
        const certInfo = response.data;

        // Пример получения данных о нагрузке CPU и RAM
        const systemInfo = await axios.get(`http://localhost:8080/systeminfo?host=${host}`);
        this.cpuUsage[host] = systemInfo.data.cpu_usage;
        this.ramUsage[host] = systemInfo.data.ram_usage;

        this.$emit('host-selected', certInfo);
      } catch (error) {
        console.error('Error fetching certificate info:', error);
      }
    },
  },
};
</script>

<style scoped>
.service-list {
  margin-top: 20px;
}

button {
  display: block;
  width: 100%;
  text-align: left;
  margin-bottom: 10px;
  cursor: pointer;
}

ul {
  list-style-type: none;
  padding-left: 20px;
}

a {
  display: block;
  margin-bottom: 5px;
  color: black;
  text-decoration: none;
}
</style>

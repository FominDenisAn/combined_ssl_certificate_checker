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
        const certResponse = await axios.get(`http://localhost:8080/certificate?host=${host}`);
        const systemInfoResponse = await axios.get(`http://localhost:8080/systeminfo?host=${host}`);

        const certInfo = certResponse.data;
        const systemInfo = systemInfoResponse.data;

        // Объединяем данные из двух запросов
        const combinedData = {
          ...certInfo,
          cpu_usage: systemInfo.cpu_usage,
          ram_usage: systemInfo.ram_usage,
        };

        this.$emit('host-selected', combinedData);
      } catch (error) {
        console.error('Error fetching certificate info or system info:', error);
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
  background-color: #e0e0e0;
  border: none;
  padding: 10px;
  font-size: 16px;
}

button:hover {
  background-color: #d0d0d0;
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
  font-size: 14px;
}
</style>

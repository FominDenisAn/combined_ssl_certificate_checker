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
        <transition-group name="fade" tag="ul" v-if="expandedServices.includes(serviceName)" class="host-list">
          <li v-for="host in hosts" :key="host" class="host-item">
            <a href="#" @click.prevent="selectHost(host)">
              {{ host }} 
              <MyIndicator :value="cpuUsage[host]" />
              <MyIndicator :value="ramUsage[host]" />
            </a>
          </li>
        </transition-group>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import MyIndicator from './MyIndicator.vue';

export default {
  components: { MyIndicator },
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

.host-list {
  list-style-type: none;
  padding-left: 20px;
}

.host-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 5px;
  padding: 5px 0;
  border: 1px solid white;
  transition: all 0.5s ease-in-out;
}

.host-item a {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-grow: 1;
  color: black;
  text-decoration: none;
  font-size: 14px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
  opacity: 0;
}
</style>

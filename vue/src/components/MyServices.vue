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
            <a href="#" @click.prevent="selectHost(host)" class="host-link">
              <span class="host-name">{{ host }}</span>
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
import MyIndicator from '@/components/MyIndicator.vue';

export default {
  name: 'MyServices',
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

        // Заполняем данные о нагрузке для каждого хоста
        for (const serviceName in this.services) {
          for (const host of this.services[serviceName]) {
            const systemInfoResponse = await axios.get(`http://localhost:8080/systeminfo?host=${host}`);
            const systemInfo = systemInfoResponse.data;
            this.cpuUsage[host] = systemInfo.cpu_usage;
            this.ramUsage[host] = systemInfo.ram_usage;
          }
        }
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
    selectHost(host) {
      this.$emit('host-selected', host); // Отправляем выбранного хоста в родительский компонент
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
  align-items: center;
  margin-bottom: 5px;
  padding: 5px 0;
  transition: all 0.5s ease-in-out;
}

.host-link {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-grow: 1;
  color: black;
  text-decoration: none;
  font-size: 14px;
}

.host-name {
  flex-grow: 1;
  margin-right: 10px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>

<template>
  <div class="dashboard">
    <h3>Dashboard</h3>
    <p>На этой странице отображаются основные метрики системы.</p>

    <!-- Пример: карточки с метриками -->
    <div class="metrics">
      <div class="metric-card">
        <h4>Total Services</h4>
        <p>{{ totalServices }}</p>
      </div>
      <div class="metric-card">
        <h4>Total Hosts</h4>
        <p>{{ totalHosts }}</p>
      </div>
      <div class="metric-card">
        <h4>Average CPU Usage</h4>
        <p>{{ averageCPUUsage }}%</p>
      </div>
      <div class="metric-card">
        <h4>Average RAM Usage</h4>
        <p>{{ averageRAMUsage }}%</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'MyDashboard',
  data() {
    return {
      services: {}, // Данные о сервисах
      hosts: [],    // Массив данных о хостах
    };
  },
  computed: {
    totalServices() {
      return Object.keys(this.services).length;
    },
    totalHosts() {
      return Object.values(this.services).flat().length;
    },
    averageCPUUsage() {
      if (this.hosts.length === 0) return 0;
      const total = this.hosts.reduce((sum, host) => sum + host.cpu_usage, 0);
      return Math.round(total / this.hosts.length);
    },
    averageRAMUsage() {
      if (this.hosts.length === 0) return 0;
      const total = this.hosts.reduce((sum, host) => sum + host.ram_usage, 0);
      return Math.round(total / this.hosts.length);
    },
  },
  created() {
    this.fetchServices();
  },
  methods: {
    async fetchServices() {
      try {
        const response = await axios.get('http://localhost:8080/services');
        this.services = response.data;

        // Преобразуем данные в массив хостов
        const hostsArray = [];
        for (const serviceName in this.services) {
          this.services[serviceName].forEach(host => {
            hostsArray.push({ host, serviceName });
          });
        }

        // Получаем информацию о системной нагрузке для каждого хоста
        for (const host of hostsArray) {
          const systemInfoResponse = await axios.get(`http://localhost:8080/systeminfo?host=${host.host}`);
          const systemInfo = systemInfoResponse.data;
          host.cpu_usage = systemInfo.cpu_usage;
          host.ram_usage = systemInfo.ram_usage;
        }

        this.hosts = hostsArray;
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      }
    },
  },
};
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.metrics {
  display: flex;
  justify-content: space-around;
  margin-top: 20px;
}

.metric-card {
  background-color: #f9f9f9;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  text-align: center;
  width: 20%;
}

.metric-card h4 {
  margin-bottom: 10px;
}
</style>

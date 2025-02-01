<template>
  <div class="analytics">
    <h3>Analytics</h3>
    <p>На этой странице будут отображаться аналитические данные о хостах.</p>

    <!-- Пример: таблица с данными -->
    <table v-if="hosts.length > 0" class="data-table">
      <thead>
        <tr>
          <th>Host</th>
          <th>CPU Usage (%)</th>
          <th>RAM Usage (%)</th>
          <th>Threat Level</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(host, index) in hosts" :key="index" :class="['threat-level', getThreatLevelClass(host)]">
          <td>{{ host.host }}</td>
          <td>{{ host.cpu_usage }}%</td>
          <td>{{ host.ram_usage }}%</td>
          <td>{{ getThreatLevel(host) }}</td>
        </tr>
      </tbody>
    </table>
    <p v-else>No data available.</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'MyAnalytics',
  data() {
    return {
      hosts: [], // Массив данных о хостах
    };
  },
  created() {
    this.fetchHosts();
  },
  methods: {
    async fetchHosts() {
      try {
        const response = await axios.get('http://localhost:8080/services');
        const services = response.data;

        // Преобразуем данные в массив хостов
        const hostsArray = [];
        for (const serviceName in services) {
          services[serviceName].forEach(host => {
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
        console.error('Error fetching analytics data:', error);
      }
    },
    getThreatLevel(host) {
      if (!host.valid_until) return '';
      const end = new Date(host.valid_until);
      const now = new Date();
      const diff = Math.floor((end - now) / (1000 * 60 * 60 * 24));
      if (diff <= 0) {
        return 'Critical';
      } else if (diff <= 7) {
        return 'High';
      } else if (diff <= 30) {
        return 'Medium';
      } else {
        return 'Low';
      }
    },
    getThreatLevelClass(host) {
      const level = this.getThreatLevel(host);
      return level.toLowerCase();
    },
  },
};
</script>

<style scoped>
.analytics {
  padding: 20px;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.data-table th,
.data-table td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.data-table tr.threat-level.critical {
  background-color: #f8d7da;
}

.data-table tr.threat-level.high {
  background-color: #fff3cd;
}

.data-table tr.threat-level.medium {
  background-color: #d4edda;
}

.data-table tr.threat-level.low {
  background-color: #cce5ff;
}
</style>

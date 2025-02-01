<template>
  <div class="analytics">
    <h3>Analytics</h3>
    <canvas id="cpuUsageChart"></canvas>
    <canvas id="ramUsageChart"></canvas>
  </div>
</template>

<script>
import { Chart } from 'chart.js/auto';
import axios from 'axios';

export default {
  name: 'MyAnalytics',
  data() {
    return {
      cpuUsageData: [],
      ramUsageData: [],
      hostLabels: [],
    };
  },
  async created() {
    await this.fetchHosts();
    this.createCharts();
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
            this.hostLabels.push(host); // Добавляем метки для графиков
          });
        }

        // Получаем информацию о системной нагрузке для каждого хоста
        for (const host of hostsArray) {
          const systemInfoResponse = await axios.get(`http://localhost:8080/systeminfo?host=${host.host}`);
          const systemInfo = systemInfoResponse.data;
          this.cpuUsageData.push(systemInfo.cpu_usage);
          this.ramUsageData.push(systemInfo.ram_usage);
        }
      } catch (error) {
        console.error('Error fetching analytics data:', error);
      }
    },
    createCharts() {
      // График CPU Usage
      const cpuCtx = document.getElementById('cpuUsageChart').getContext('2d');
      new Chart(cpuCtx, {
        type: 'bar',
        data: {
          labels: this.hostLabels,
          datasets: [
            {
              label: 'CPU Usage (%)',
              data: this.cpuUsageData,
              backgroundColor: 'rgba(75, 192, 192, 0.6)',
              borderColor: 'rgba(75, 192, 192, 1)',
              borderWidth: 1,
            },
          ],
        },
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });

      // График RAM Usage
      const ramCtx = document.getElementById('ramUsageChart').getContext('2d');
      new Chart(ramCtx, {
        type: 'bar',
        data: {
          labels: this.hostLabels,
          datasets: [
            {
              label: 'RAM Usage (%)',
              data: this.ramUsageData,
              backgroundColor: 'rgba(255, 99, 132, 0.6)',
              borderColor: 'rgba(255, 99, 132, 1)',
              borderWidth: 1,
            },
          ],
        },
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    },
  },
};
</script>

<style scoped>
.analytics {
  padding: 20px;
}

canvas {
  max-width: 100%;
  margin-bottom: 20px;
}
</style>

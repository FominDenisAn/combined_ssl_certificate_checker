<template>
  <div v-if="selectedHost">
    <h3>Certificate Information for {{ selectedHost.host }}</h3>
    <p><strong>Issuer:</strong> {{ selectedHost.issuer }}</p>
    <p><strong>Days Left:</strong> <span :class="['threat-level', threatLevelClass]">{{ daysLeft }} days</span></p>
    <p><strong>Start Date:</strong> {{ formattedStartDate }}</p>
    <p><strong>End Date:</strong> {{ formattedEndDate }}</p>
    <p><strong>Serial Number:</strong> {{ selectedHost.serial_number }}</p>
    <p><strong>System Time:</strong> {{ selectedHost.system_time }}</p>
    <p><strong>Uptime:</strong> {{ selectedHost.uptime }} seconds</p>
    <p><strong>CPU Usage:</strong> {{ selectedHost.cpu_usage }}%</p>
    <p><strong>RAM Usage:</strong> {{ selectedHost.ram_usage }}%</p>
    <p><strong>Threat Level:</strong> {{ threatLevel }}</p>
    <p><strong>Target Path:</strong> {{ selectedHost.target_path }}</p>
  </div>
  <div v-else>
    <p>Select a host to view certificate information.</p>
  </div>
</template>

<script>
import dayjs from 'dayjs';

export default {
  props: ['selectedHost'],
  computed: {
    daysLeft() {
      if (!this.selectedHost || !this.selectedHost.valid_until) {
        return '';
      }
      const end = dayjs(this.selectedHost.valid_until);
      const now = dayjs();
      const diff = end.diff(now, 'day');
      return diff;
    },
    formattedStartDate() {
      if (!this.selectedHost || !this.selectedHost.valid_from) {
        return '';
      }
      return dayjs(this.selectedHost.valid_from).format('YYYY-MM-DD HH:mm:ss');
    },
    formattedEndDate() {
      if (!this.selectedHost || !this.selectedHost.valid_until) {
        return '';
      }
      return dayjs(this.selectedHost.valid_until).format('YYYY-MM-DD HH:mm:ss');
    },
    threatLevel() {
      if (!this.selectedHost || !this.selectedHost.valid_until) {
        return '';
      }
      const end = dayjs(this.selectedHost.valid_until);
      const now = dayjs();
      const diff = end.diff(now, 'day');
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
    threatLevelClass() {
      if (!this.selectedHost || !this.selectedHost.valid_until) {
        return '';
      }
      const end = dayjs(this.selectedHost.valid_until);
      const now = dayjs();
      const diff = end.diff(now, 'day');
      if (diff <= 0) {
        return 'critical';
      } else if (diff <= 7) {
        return 'high';
      } else if (diff <= 30) {
        return 'medium';
      } else {
        return 'low';
      }
    },
  },
};
</script>

<style scoped>
/* Add your styles here */
.threat-level {
  font-weight: bold;
  color: green;
}

.threat-level.critical {
  color: red;
}

.threat-level.high {
  color: orange;
}

.threat-level.medium {
  color: yellow;
}
</style>

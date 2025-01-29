<template>
  <div>
    <h3>Certificate Information</h3>
    <div v-if="selectedHost">
      <p><strong>Host:</strong> {{ selectedHost.host }}</p>
      <p><strong>Issuer:</strong> {{ selectedHost.issuer }}</p>
      <p><strong>Valid From:</strong> {{ selectedHost.valid_from }}</p>
      <p><strong>Valid Until:</strong> {{ selectedHost.valid_until }}</p>
      <p><strong>Serial Number:</strong> {{ selectedHost.serial_number }}</p>
      <p><strong>System Time:</strong> {{ selectedHost.system_time }}</p>
      <p><strong>Uptime:</strong> {{ selectedHost.uptime }} seconds</p>
      <p><strong>CPU Usage:</strong> {{ selectedHost.cpu_usage }}%</p>
      <p><strong>RAM Usage:</strong> {{ selectedHost.ram_usage }}%</p>
    </div>
    <div v-else>
      <p>Select a host to view certificate information.</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ['selectedHost'],
};
</script>

<style scoped>
/* Add your styles here */
</style>

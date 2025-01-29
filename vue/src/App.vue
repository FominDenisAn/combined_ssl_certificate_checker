<template>
  <div id="app">
    <div class="sidebar">
      <ServiceList @host-selected="handleHostSelected" />
      <ImportExport />
    </div>
    <div class="content">
      <CertificateInfo :selectedHost="selectedHost" />
    </div>
  </div>
</template>

<script>
import ServiceList from './components/ServiceList.vue';
import CertificateInfo from './components/CertificateInfo.vue';
import ImportExport from './components/ImportExport.vue';

export default {
  name: 'App',
  components: {
    ServiceList,
    CertificateInfo,
    ImportExport,
  },
  data() {
    return {
      selectedHost: null,
    };
  },
  methods: {
    handleHostSelected(host) {
      this.selectedHost = host;
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  display: flex;
  height: 100vh;
  margin: 0;
}

.sidebar {
  width: 300px; /* Фиксируем ширину */
  background-color: #f4f4f4;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
}

.content {
  flex-grow: 1;
  padding: 20px;
  transition: opacity 0.5s ease-in-out;
}
</style>

<template>
  <div id="app">
    <div class="controls">
      <ImportExport />
    </div>
    <div class="sidebar">
      <ServiceList @host-selected="handleHostSelected" />
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
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin: 0;
}

.controls {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #f0f0f0;
  border-bottom: 1px solid #ccc;
}

.sidebar {
  width: 300px; /* Фиксированная ширина */
  background-color: #f4f4f4;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
  flex-grow: 1;
}

.content {
  flex-grow: 1;
  padding: 20px;
  transition: opacity 0.5s ease-in-out;
  background-color: #ffffff;
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

.indicator-container {
  display: inline-block;
  position: relative;
  width: 100px;
  height: 20px;
  border: 1px solid #ccc;
  margin-right: 10px;
}

.indicator-bar {
  height: 100%;
}

span {
  position: absolute;
  right: 5px;
  top: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
</style>

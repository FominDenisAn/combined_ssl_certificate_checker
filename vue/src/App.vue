<template>
  <div id="app">
    <div class="controls">
      <ImportExport />
    </div>
    <div class="main-container">
      <div class="sidebar">
        <ServiceList @host-selected="handleHostSelected" />
      </div>
      <div class="content">
        <CertificateInfo :selectedHost="selectedHost" />
      </div>
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

.main-container {
  display: flex;
  flex-grow: 1;
}

.sidebar {
  width: 350px; /* Увеличиваем ширину */
  background-color: #f4f4f4;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
}

.content {
  flex-grow: 1;
  padding: 20px;
  transition: opacity 0.5s ease-in-out;
  background-color: #ffffff;
  overflow-y: auto; /* Добавляем скроллбар для контента справа */
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
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: black;
  text-decoration: none;
  font-size: 14px;
  padding: 5px 0;
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
  display: inline-flex;
  align-items: center;
  width: 100px;
  height: 20px;
  border: 1px solid #ccc;
  margin-right: 10px;
}

.indicator-bar {
  height: 100%;
  width: 0;
  transition: width 0.5s ease-in-out;
}

span {
  margin-left: 5px;
  font-size: 12px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
  opacity: 0;
}
</style>

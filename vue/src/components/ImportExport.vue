<template>
  <div class="import-export">
    <h3>Import/Export Services</h3>
    <textarea v-model="jsonData" placeholder="Paste JSON data here"></textarea>
    <button @click="importData">Import</button>
    <button @click="exportData">Export</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      jsonData: '',
    };
  },
  methods: {
    async importData() {
      try {
        const parsedData = JSON.parse(this.jsonData);
        await axios.post('http://localhost:8080/import', parsedData);
        alert('Data imported successfully');
      } catch (error) {
        console.error('Error importing data:', error);
        alert('Failed to import data');
      }
    },
    async exportData() {
      try {
        const response = await axios.get('http://localhost:8080/export');
        this.jsonData = JSON.stringify(response.data, null, 2);
      } catch (error) {
        console.error('Error exporting data:', error);
        alert('Failed to export data');
      }
    },
  },
};
</script>

<style scoped>
.import-export {
  margin-top: 20px;
}

textarea {
  width: 100%;
  height: 150px;
}

button {
  margin-top: 10px;
  padding: 10px 20px;
  cursor: pointer;
  background-color: #e0e0e0;
  border: none;
  margin-right: 10px;
}

button:hover {
  background-color: #d0d0d0;
}

.controls {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #f0f0f0;
  border-bottom: 1px solid #ccc;
}
</style>

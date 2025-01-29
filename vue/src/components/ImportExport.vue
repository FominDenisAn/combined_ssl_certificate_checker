<template>
  <div class="import-export">
    <button @click="exportData" class="export-button">Export</button>
    <input type="file" id="importFile" @change="importData" style="display:none;" ref="importFileInput"/>
    <button @click="triggerFileInput" class="import-button">Import</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  methods: {
    triggerFileInput() {
      this.$refs.importFileInput.click();
    },
    async importData(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = async (e) => {
          try {
            const jsonData = JSON.parse(e.target.result);
            await axios.post('http://localhost:8080/import', jsonData);
            alert('Data imported successfully');
          } catch (error) {
            console.error('Error importing data:', error);
            alert('Failed to import data');
          }
        };
        reader.readAsText(file);
      }
    },
    async exportData() {
      try {
        const response = await axios.get('http://localhost:8080/export');
        const jsonData = JSON.stringify(response.data, null, 2);
        const blob = new Blob([jsonData], { type: 'application/json' });
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = 'services.json';
        link.click();
        URL.revokeObjectURL(link.href);
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
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #f0f0f0;
  border-bottom: 1px solid #ccc;
}

.export-button, .import-button {
  padding: 10px 20px;
  cursor: pointer;
  background-color: #e0e0e0;
  border: none;
  font-size: 14px;
}

.export-button:hover, .import-button:hover {
  background-color: #d0d0d0;
}
</style>

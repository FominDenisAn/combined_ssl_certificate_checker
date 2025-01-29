<template>
  <div class="indicator-container">
    <div class="indicator-bar" :style="{ width: value + '%', backgroundColor: getColor(value) }"></div>
    <span>{{ value }}%</span>
  </div>
</template>

<script>
export default {
  props: ['value'],
  methods: {
    getColor(value) {
      if (value < 30) {
        return 'green';
      } else if (value < 70) {
        return 'orange';
      } else {
        return 'red';
      }
    },
  },
};
</script>

<style scoped>
.indicator-container {
  display: inline-flex;
  align-items: center;
  width: 100px;
  height: 20px;
  border: 1px solid #ccc;
  margin-right: 10px;
  padding: 2px;
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
</style>

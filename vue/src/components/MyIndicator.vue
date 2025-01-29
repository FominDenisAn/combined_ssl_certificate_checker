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

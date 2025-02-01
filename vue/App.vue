<template>
  <div id="app">
    <!-- Меню навигации -->
    <nav class="navbar">
      <router-link to="/services" class="nav-item">Services</router-link>
      <router-link to="/dashboard" class="nav-item">Dashboard</router-link>
      <router-link to="/analytics" class="nav-item">Analytics</router-link>
    </nav>

    <!-- Основной контент -->
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin: 0;
}

.navbar {
  background-color: #333;
  padding: 10px;
  display: flex;
  justify-content: space-around;
}

.nav-item {
  color: white;
  text-decoration: none;
  font-size: 16px;
  padding: 10px;
}

.nav-item.router-link-exact-active {
  color: #42b983;
  border-bottom: 2px solid #42b983;
}

main {
  flex-grow: 1;
  padding: 20px;
  background-color: #ffffff;
}
</style>

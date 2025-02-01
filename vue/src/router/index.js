import { createRouter, createWebHistory } from 'vue-router';

// Импортируем компоненты из правильной директории (views)
import MyServices from '../views/MyServices.vue';
import MyDashboard from '../views/MyDashboard.vue';
import MyAnalytics from '../views/MyAnalytics.vue';

const routes = [
  {
    path: '/services',
    name: 'MyServices',
    component: MyServices,
  },
  {
    path: '/dashboard',
    name: 'MyDashboard',
    component: MyDashboard,
  },
  {
    path: '/analytics',
    name: 'MyAnalytics',
    component: MyAnalytics,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

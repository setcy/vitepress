import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/a/:pathMatch(.*)*',
      name: 'admin',
      component: () => import('../views/AdminView.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'content',
      component: () => import('../views/ContentView.vue')
    },
  ]
})

export default router

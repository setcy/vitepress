import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // {
    //   path: '/',
    //   name: 'home',
    //   component: () => import('../views/HomeView.vue')
    // },
    {
      path: '/a/:catchAll(.*)*',
      name: 'admin',
      component: () => import('../views/AdminView.vue')
    },
    {
      path: '/',
      name: 'content',
      component: () => import('../views/ContentView.vue'),
      children: [
        {
          path: '/:catchAll(.*)*',
          name: 'content',
          component: () => import('../components/content/VPContent.vue')
        }
      ]
    },
  ]
})

export default router

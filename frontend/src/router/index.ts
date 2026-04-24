/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/index.vue'
import { useAppStore } from '@/stores/app'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Index,
    },
    {
      path: '/auth',
      component: () => import('@/pages/auth.vue'),
      name: 'auth',
    }
  ],
})

router.beforeEach((to, from, next) => {
  const appStore = useAppStore()

  if (to.path !== '/auth' && !appStore.isAuthenticated) {
    next('/auth')
  } else {
    next()
  }
})

export default router

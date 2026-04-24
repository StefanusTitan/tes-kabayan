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
    },
    {
      path: '/pembeli',
      component: () => import('@/pages/pembeli/index.vue'),
      name: 'pembeli',
    },
    // {
    //   path: '/barang',
    //   component: () => import('@/pages/barang.vue'),
    //   name: 'barang',
    // },
    // {
    //   path: '/transaksi',
    //   component: () => import('@/pages/transaksi.vue'),
    //   name: 'transaksi',
    // },
    // {
    //   path: '/insights',
    //   component: () => import('@/pages/insights.vue'),
    //   name: 'insights',
    // }
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

<template>
  <v-app-bar app flat border class="px-2" density="comfortable">
    <div class="flex px-2">
      <v-btn to="/" variant="text" prepend-icon="mdi-storefront-outline" class="mr-2 text-none" rounded="pill">
        Warung Kabayan
      </v-btn>

      <v-spacer />

      <div class="nav-links d-flex align-center flex-nowrap ga-1">
        <v-btn
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          :variant="isActive(item.path) ? 'flat' : 'text'"
          class="text-none"
          color="primary"
          rounded="pill"
        >
          {{ item.label }}
        </v-btn>

        <v-btn
          color="error"
          variant="text"
          class="text-none"
          rounded="pill"
          prepend-icon="mdi-logout"
          :loading="isLoggingOut"
          @click="handleLogout"
        >
          Logout
        </v-btn>
      </div>
    </div>
  </v-app-bar>
</template>

<script setup lang="ts">
  import { computed, ref } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { logout } from '@/api/auth'
  import { useAppStore } from '@/stores/app'

  type NavItem = {
    path: string
    label: string
  }

  const router = useRouter()
  const route = useRoute()
  const appStore = useAppStore()
  const isLoggingOut = ref(false)

  const navItems = computed<NavItem[]>(() => {
    return router.getRoutes()
      .filter(r => {
        const label = r.meta?.navLabel
        return typeof label === 'string' && label.length > 0 && !r.path.includes(':')
      })
      .map(r => ({
        path: r.path,
        label: String(r.meta.navLabel),
      }))
      .sort((a, b) => a.path.localeCompare(b.path))
  })

  const isActive = (path: string) => {
    if (path === '/') return route.path === '/'

    return route.path === path || route.path.startsWith(`${path}/`)
  }

  const handleLogout = async () => {
    if (isLoggingOut.value) {
      return
    }

    try {
      isLoggingOut.value = true
      await logout()
      appStore.isAuthenticated = false
      appStore.user = null
      await router.push('/auth')
    } catch (error) {
      console.error('Logout failed from navigation:', error)
    } finally {
      isLoggingOut.value = false
    }
  }
</script>

<style scoped>
  .nav-links {
    overflow-x: auto;
    padding-bottom: 2px;
  }

  .nav-links :deep(.v-btn) {
    flex: 0 0 auto;
  }
</style>
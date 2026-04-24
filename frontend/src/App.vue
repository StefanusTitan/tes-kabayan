<template>
  <v-app>
    <HorizontalNav v-if="showNav" />

    <v-main :class="showNav ? 'py-6' : 'flex items-center justify-center'">
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
  import { computed } from 'vue'
  import { useRoute } from 'vue-router'
  import HorizontalNav from '@/components/HorizontalNav.vue'

  const route = useRoute()

  const showNav = computed(() => !route.meta?.hideNav)
</script>

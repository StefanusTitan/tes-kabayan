<template>
  <v-container class="h-full flex items-center" max-width="900">
    <div>
      <div class="mb-8 text-center">
        <div class="font-light -mb-1">Welcome to</div>
        <h1 class="text-6xl font-heading font-bold">Warung Kabayan</h1>
      </div>

      <div class="grid md:grid-cols-2 gap-4">
        <v-card
          v-for="link in links"
          :to="link.route"
          :key="link.route"
          :subtitle="link.subtitle"
          :title="link.title"
          class="
            h-full py-3 rounded-3xl transition-[border-radius] hover:rounded-lg
            flex items-center [&>.v-card-item]:w-full
            group"
          variant="flat"
        >
          <template #prepend>
            <v-avatar :icon="link.icon" size="60" variant="tonal" class="ml-2 mr-4" />
          </template>
          <template #append>
            <v-icon class="ml-1 opacity-0 transition group-hover:opacity-90 group-hover:translate-x-1" icon="mdi-arrow-right" />
          </template>
          <template #subtitle>
            <div class="line-clamp-2 text-wrap">{{ link.subtitle }}</div>
          </template>
        </v-card>
      </div>
    </div>
  </v-container>
</template>

<script setup lang="ts">
  const links = [
    {
      route: '/pembeli',
      icon: 'mdi-account-group',
      subtitle: 'Informasi mengenai pembeli pada Warung Kabayan.',
      title: 'Pembeli',
    },
    {
      route: '/barang',
      icon: 'mdi-package-variant-closed',
      subtitle: 'Eksplor semua barang pada Warung Kabayan.',
      title: 'Barang',
    },
    {
      route: '/transaksi',
      icon: 'mdi-swap-horizontal',
      subtitle: 'Masukkan data transaksi pada Warung Kabayan.',
      title: 'Transaksi',
    },
    {
      route: '/insights',
      icon: 'mdi-view-dashboard',
      subtitle: 'Dapatkan wawasan tentang operasional Warung Kabayan.',
      title: 'Insights',
    },
  ]
</script>

<style scoped>
@reference "../styles/tailwind.css";

/*
  1. mixing helper classes and @apply for demonstration purposes only
  2. the classes below are NOT wrapped in any CSS layer, so they "win" over everything else
*/
.hero-card {
  @apply py-3 md:pr-[120px] w-full transition-none;
}

:deep(.v-card) {
  @apply bg-gray-200;
  @apply dark:bg-black dark:bg-linear-to-r dark:from-primary/50 dark:to-primary/30 dark:text-white/80;
}
</style>

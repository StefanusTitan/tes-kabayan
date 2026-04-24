<template>
  <v-container class="min-h-screen p-4">
    <v-card class="mx-auto w-full max-w-5xl !rounded-xl shadow-2xl">
      <v-card-title class="d-flex align-center justify-between py-4 w-100">
        <span class="text-h6">Data Pembeli</span>
        <v-btn color="primary" @click="isCreateModalOpen = true">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-card-title>

      <v-card-text class="pt-0">
        <v-row class="mb-2" align="center">
          <v-col cols="12" md="9">
            <v-text-field
              v-model="searchNama"
              label="Cari nama pembeli"
              variant="outlined"
              density="comfortable"
              hide-details
              @keyup.enter="handleSearch"
            />
          </v-col>
          <v-col cols="12" md="3" class="d-flex">
            <v-btn
              color="primary"
              block
              :loading="loading"
              prepend-icon="mdi-magnify"
              @click="handleSearch"
            >
              Cari
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>

      <v-table>
        <thead>
          <tr>
            <th class="text-left">ID</th>
            <th class="text-left">Nama</th>
            <th class="text-left">Alamat</th>
            <th class="text-left">No. Telepon</th>
            <th class="text-left">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="pembeli in pembelis" :key="pembeli.id">
            <td>{{ pembeli.id }}</td>
            <td>{{ pembeli.nama }}</td>
            <td>{{ pembeli.alamat }}</td>
            <td>{{ pembeli.no_telp }}</td>
            <td>
              <v-btn size="small" variant="text" color="primary" icon="mdi-pencil" @click="openEditModal(pembeli)" />
            </td>
          </tr>
          <tr v-if="!loading && pembelis.length === 0">
            <td colspan="5" class="py-6 text-center text-medium-emphasis">
              Data pembeli tidak ditemukan.
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-card>

    <PembeliCreateModal
      v-model="isCreateModalOpen"
      @saved="handleMutated"
    />
    <PembeliUpdateModal
      v-model="isUpdateModalOpen"
      :selected-pembeli="selectedPembeli"
      @saved="handleMutated"
    />
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import PembeliCreateModal from '@/components/PembeliCreateModal.vue'
import PembeliUpdateModal from '@/components/PembeliUpdateModal.vue'
import { getSemuaPembeli, type Pembeli } from '@/api/pembeli'

const pembelis = ref<Pembeli[]>([])
const searchNama = ref('')
const loading = ref(false)
const isCreateModalOpen = ref(false)
const isUpdateModalOpen = ref(false)
const selectedPembeli = ref<Pembeli | null>(null)

const fetchPembeli = async () => {
  try {
    loading.value = true
    const response = await getSemuaPembeli(searchNama.value.trim())
    pembelis.value = response
  } catch (error) {
    console.error('Error fetching pembeli:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = async () => {
  await fetchPembeli()
}

const openEditModal = (pembeli: Pembeli) => {
  selectedPembeli.value = pembeli
  isUpdateModalOpen.value = true
}

const handleMutated = async () => {
  await fetchPembeli()
}

onMounted(async () => {
  await fetchPembeli()
})
</script>
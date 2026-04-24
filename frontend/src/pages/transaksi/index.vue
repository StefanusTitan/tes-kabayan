<template>
  <v-container class="min-h-screen p-4">
    <v-card class="mx-auto mb-4 w-full max-w-6xl !rounded-xl shadow-2xl">
      <v-card-title class="py-4 text-h6">Tambah Transaksi</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" md="4">
            <v-select
              v-model="form.pembeliId"
              :items="pembeliOptions"
              item-title="label"
              item-value="value"
              label="Pembeli"
              variant="outlined"
              density="comfortable"
              :loading="loadingFormDeps"
              hide-details="auto"
            />
          </v-col>
          <v-col cols="12" md="4">
            <v-select
              v-model="form.barangId"
              :items="barangOptions"
              item-title="label"
              item-value="value"
              label="Barang"
              variant="outlined"
              density="comfortable"
              :loading="loadingFormDeps"
              hide-details="auto"
            />
          </v-col>
          <v-col cols="12" md="2">
            <v-text-field
              v-model.number="form.quantity"
              type="number"
              min="1"
              label="Qty"
              variant="outlined"
              density="comfortable"
              hide-details="auto"
            />
          </v-col>
          <v-col cols="12" md="2" class="d-flex">
            <v-btn
              block
              color="primary"
              :loading="creating"
              prepend-icon="mdi-content-save"
              @click="handleCreate"
            >
              Simpan
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <v-card class="mx-auto w-full max-w-6xl !rounded-xl shadow-2xl">
      <v-card-title class="d-flex align-center justify-space-between py-4">
        <span class="text-h6">Data Transaksi</span>
      </v-card-title>

      <v-card-text class="pt-0">
        <v-row class="mb-2" align="center">
          <v-col cols="12" md="9">
            <v-text-field
              v-model="search"
              label="Cari berdasarkan nama pembeli / barang"
              variant="outlined"
              density="comfortable"
              hide-details
              @keyup.enter="fetchTransaksi"
            />
          </v-col>
          <v-col cols="12" md="3" class="d-flex">
            <v-btn
              color="primary"
              block
              :loading="loadingTable"
              prepend-icon="mdi-magnify"
              @click="fetchTransaksi"
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
            <th class="text-left">Tanggal</th>
            <th class="text-left">Pembeli</th>
            <th class="text-left">Barang</th>
            <th class="text-left">Qty</th>
            <th class="text-left">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in transaksis" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ formatDate(item.created_at) }}</td>
            <td>{{ item.pembeli?.nama || '-' }}</td>
            <td>{{ item.barang?.nama || '-' }}</td>
            <td>{{ item.quantity }}</td>
            <td>
              <v-btn
                size="small"
                variant="text"
                color="error"
                icon="mdi-delete"
                :loading="deletingId === item.id"
                @click="handleDelete(item)"
              />
            </td>
          </tr>
          <tr v-if="!loadingTable && transaksis.length === 0">
            <td colspan="6" class="py-6 text-center text-medium-emphasis">
              Data transaksi tidak ditemukan.
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-card>
  </v-container>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { createTransaksi, deleteTransaksi, getSemuaTransaksi, type Transaksi } from '@/api/transaksi'
import { getSemuaPembeli } from '@/api/pembeli'
import { getSemuaBarang } from '@/api/barang'

const transaksis = ref<Transaksi[]>([])
const search = ref('')
const loadingTable = ref(false)
const creating = ref(false)
const deletingId = ref<number | null>(null)
const loadingFormDeps = ref(false)

const pembelis = ref<Array<{ id: number; nama: string }>>([])
const barangs = ref<Array<{ id: number; nama: string }>>([])

const form = reactive({
  pembeliId: null as number | null,
  barangId: null as number | null,
  quantity: 1,
})

const pembeliOptions = computed(() => {
  return pembelis.value.map((p) => ({
    label: `${p.nama} (#${p.id})`,
    value: p.id,
  }))
})

const barangOptions = computed(() => {
  return barangs.value.map((b) => ({
    label: `${b.nama} (#${b.id})`,
    value: b.id,
  }))
})

const formatDate = (value: string) => {
  return new Intl.DateTimeFormat('id-ID', {
    dateStyle: 'medium',
    timeStyle: 'short',
  }).format(new Date(value))
}

const fetchTransaksi = async () => {
  try {
    loadingTable.value = true
    transaksis.value = await getSemuaTransaksi({
      search: search.value.trim() || undefined,
    })
  } catch (error) {
    console.error('Error fetching transaksi:', error)
  } finally {
    loadingTable.value = false
  }
}

const fetchFormDependencies = async () => {
  try {
    loadingFormDeps.value = true
    const [pembeliData, barangData] = await Promise.all([
      getSemuaPembeli(),
      getSemuaBarang(),
    ])

    pembelis.value = pembeliData
    barangs.value = barangData

    if (!form.pembeliId && pembeliData.length > 0) {
      form.pembeliId = pembeliData[0].id
    }

    if (!form.barangId && barangData.length > 0) {
      form.barangId = barangData[0].id
    }
  } catch (error) {
    console.error('Error fetching transaksi form dependencies:', error)
  } finally {
    loadingFormDeps.value = false
  }
}

const resetForm = () => {
  form.quantity = 1
}

const handleCreate = async () => {
  if (!form.pembeliId || !form.barangId || form.quantity < 1) {
    window.alert('Pilih pembeli, barang, dan quantity minimal 1.')
    return
  }

  try {
    creating.value = true
    await createTransaksi({
      pembeli_id: form.pembeliId,
      barang_id: form.barangId,
      quantity: Number(form.quantity),
    })
    resetForm()
    await fetchTransaksi()
  } catch (error) {
    console.error('Error creating transaksi:', error)
  } finally {
    creating.value = false
  }
}

const handleDelete = async (item: Transaksi) => {
  const isConfirmed = window.confirm(`Hapus transaksi #${item.id}?`)
  if (!isConfirmed) {
    return
  }

  try {
    deletingId.value = item.id
    await deleteTransaksi(item.id)
    await fetchTransaksi()
  } catch (error) {
    console.error('Error deleting transaksi:', error)
  } finally {
    deletingId.value = null
  }
}

onMounted(async () => {
  await Promise.all([fetchFormDependencies(), fetchTransaksi()])
})
</script>

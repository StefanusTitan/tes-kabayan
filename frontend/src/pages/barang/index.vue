<template>
	<v-container class="min-h-screen p-4">
		<v-card class="mx-auto w-full max-w-5xl !rounded-xl shadow-2xl">
			<v-card-title class="d-flex align-center justify-between py-4 w-100">
				<span class="text-h6">Data Barang</span>
				<v-btn color="primary" @click="isCreateModalOpen = true">
					<v-icon>mdi-plus</v-icon>
				</v-btn>
			</v-card-title>

			<v-card-text class="pt-0">
				<v-row class="mb-2" align="center">
					<v-col cols="12" md="9">
						<v-text-field
							v-model="searchNama"
							label="Cari nama barang"
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
						<th class="text-left">Deskripsi</th>
						<th class="text-left">Harga</th>
						<th class="text-left">Stok</th>
						<th class="text-left">Aksi</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="barang in barangs" :key="barang.id">
						<td>{{ barang.id }}</td>
						<td>{{ barang.nama }}</td>
						<td>{{ barang.deskripsi || '-' }}</td>
						<td>{{ formatCurrency(barang.harga) }}</td>
						<td>{{ barang.stock }}</td>
						<td>
							<v-btn
								size="small"
								variant="text"
								color="primary"
								icon="mdi-pencil"
								@click="openEditModal(barang)"
							/>
							<v-btn
								size="small"
								variant="text"
								color="error"
								icon="mdi-delete"
								:loading="deletingId === barang.id"
								@click="handleDelete(barang)"
							/>
						</td>
					</tr>
					<tr v-if="!loading && barangs.length === 0">
						<td colspan="6" class="py-6 text-center text-medium-emphasis">
							Data barang tidak ditemukan.
						</td>
					</tr>
				</tbody>
			</v-table>
		</v-card>

		<BarangCreateModal
			v-model="isCreateModalOpen"
			@saved="handleMutated"
		/>
		<BarangUpdateModal
			v-model="isUpdateModalOpen"
			:selected-barang="selectedBarang"
			@saved="handleMutated"
		/>
	</v-container>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import BarangCreateModal from '@/components/BarangCreateModal.vue'
import BarangUpdateModal from '@/components/BarangUpdateModal.vue'
import { deleteBarang, getSemuaBarang, type Barang } from '@/api/barang'

const barangs = ref<Barang[]>([])
const searchNama = ref('')
const loading = ref(false)
const deletingId = ref<number | null>(null)
const isCreateModalOpen = ref(false)
const isUpdateModalOpen = ref(false)
const selectedBarang = ref<Barang | null>(null)

const formatCurrency = (value: number) => {
	return new Intl.NumberFormat('id-ID', {
		style: 'currency',
		currency: 'IDR',
		maximumFractionDigits: 0
	}).format(value || 0)
}

const fetchBarang = async () => {
	try {
		loading.value = true
		const response = await getSemuaBarang(searchNama.value.trim())
		barangs.value = response
	} catch (error) {
		console.error('Error fetching barang:', error)
	} finally {
		loading.value = false
	}
}

const handleSearch = async () => {
	await fetchBarang()
}

const openEditModal = (barang: Barang) => {
	selectedBarang.value = barang
	isUpdateModalOpen.value = true
}

const handleMutated = async () => {
	await fetchBarang()
}

const handleDelete = async (barang: Barang) => {
	const isConfirmed = window.confirm(`Hapus barang \"${barang.nama}\"?`)
	if (!isConfirmed) {
		return
	}

	try {
		deletingId.value = barang.id
		await deleteBarang(barang.id)
		await fetchBarang()
	} catch (error) {
		console.error('Error deleting barang:', error)
	} finally {
		deletingId.value = null
	}
}

onMounted(async () => {
	await fetchBarang()
})
</script>

<template>
	<v-container class="min-h-screen p-4">
		<v-card class="mx-auto mb-4 w-full max-w-6xl !rounded-xl shadow-2xl">
			<v-card-title class="d-flex align-center justify-space-between py-4">
				<span class="text-h6">Insights Penjualan</span>
			</v-card-title>

			<v-card-text>
				<v-row>
					<v-col cols="12" md="3">
						<v-select
							v-model="period"
							:items="periodOptions"
							item-title="label"
							item-value="value"
							label="Periode"
							variant="outlined"
							density="comfortable"
							hide-details
						/>
					</v-col>
					<v-col v-if="period === 'range'" cols="12" md="3">
						<v-text-field
							v-model="startDate"
							type="date"
							label="Tanggal Mulai"
							variant="outlined"
							density="comfortable"
							hide-details
						/>
					</v-col>
					<v-col v-if="period === 'range'" cols="12" md="3">
						<v-text-field
							v-model="endDate"
							type="date"
							label="Tanggal Akhir"
							variant="outlined"
							density="comfortable"
							hide-details
						/>
					</v-col>
					<v-col cols="12" md="3" class="d-flex">
						<v-btn
							color="primary"
							block
							prepend-icon="mdi-chart-bar"
							:loading="loading"
							@click="fetchData"
						>
							Terapkan
						</v-btn>
					</v-col>
				</v-row>
			</v-card-text>
		</v-card>

		<v-card class="mx-auto mb-4 w-full max-w-6xl !rounded-xl shadow-2xl">
			<v-card-title class="py-4 text-subtitle-1">Grafik Jumlah Terjual</v-card-title>
			<v-card-text>
				<div v-if="chartData.length === 0" class="py-8 text-center text-medium-emphasis">
					Belum ada data untuk ditampilkan.
				</div>
				<div v-else class="chart-wrap">
					<v-sparkline
						:model-value="sparklineValues"
						color="primary"
						line-width="3"
						smooth
						auto-draw
						padding="18"
						height="180"
					/>
					<div v-if="!hasAnySales" class="mt-4 text-center text-medium-emphasis">
						Belum ada penjualan di periode ini (semua nilai masih 0).
					</div>
					<div class="chart-meta-grid">
						<div v-for="point in chartData" :key="point.key" class="chart-meta-item">
							<strong>{{ point.value }}</strong>
							<span>{{ point.label }}</span>
						</div>
					</div>
				</div>
			</v-card-text>
		</v-card>

		<v-card class="mx-auto w-full max-w-6xl !rounded-xl shadow-2xl">
			<v-card-title class="py-4 text-subtitle-1">Daftar Barang Terjual</v-card-title>
			<v-table>
				<thead>
					<tr>
						<th class="text-left">Tanggal</th>
						<th class="text-left">Barang</th>
						<th class="text-left">Dibeli Oleh</th>
						<th class="text-left">Quantity</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="item in transaksis" :key="item.id">
						<td>{{ formatDate(item.created_at) }}</td>
						<td>{{ item.barang?.nama || '-' }}</td>
						<td>{{ item.pembeli?.nama || '-' }}</td>
						<td>{{ item.quantity }}</td>
					</tr>
					<tr v-if="!loading && transaksis.length === 0">
						<td colspan="4" class="py-6 text-center text-medium-emphasis">
							Tidak ada transaksi pada periode ini.
						</td>
					</tr>
				</tbody>
			</v-table>
		</v-card>
	</v-container>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import { getSemuaTransaksi, type Transaksi, type TransaksiFilter } from '@/api/transaksi'

type PeriodOption = 'weekly' | 'monthly' | 'range'

const loading = ref(false)
const transaksis = ref<Transaksi[]>([])
const period = ref<PeriodOption>('weekly')
const startDate = ref('')
const endDate = ref('')

const periodOptions = [
	{ label: 'Mingguan (7 hari terakhir)', value: 'weekly' },
	{ label: 'Bulanan (30 hari terakhir)', value: 'monthly' },
	{ label: 'Rentang Tanggal', value: 'range' },
]

const toInputDate = (value: Date) => {
	const year = value.getFullYear()
	const month = `${value.getMonth() + 1}`.padStart(2, '0')
	const day = `${value.getDate()}`.padStart(2, '0')
	return `${year}-${month}-${day}`
}

const formatDate = (value: string) => {
	return new Intl.DateTimeFormat('id-ID', {
		dateStyle: 'medium',
		timeStyle: 'short',
	}).format(new Date(value))
}

const formatShortDate = (value: Date) => {
	return new Intl.DateTimeFormat('id-ID', {
		day: '2-digit',
		month: '2-digit',
	}).format(value)
}

const normalizeDate = (value: Date) => {
	return new Date(value.getFullYear(), value.getMonth(), value.getDate())
}

const addDays = (value: Date, days: number) => {
	const next = new Date(value)
	next.setDate(next.getDate() + days)
	return next
}

const getFilterByPeriod = (): TransaksiFilter => {
	const today = normalizeDate(new Date())

	if (period.value === 'weekly') {
		return {
			start_date: toInputDate(addDays(today, -6)),
			end_date: toInputDate(today),
		}
	}

	if (period.value === 'monthly') {
		return {
			start_date: toInputDate(addDays(today, -29)),
			end_date: toInputDate(today),
		}
	}

	if (!startDate.value || !endDate.value) {
		return {}
	}

	return {
		start_date: startDate.value,
		end_date: endDate.value,
	}
}

const buildDailyBuckets = () => {
	const buckets: Array<{ key: string; label: string; value: number }> = []
	const filter = getFilterByPeriod()

	if (!filter.start_date || !filter.end_date) {
		return buckets
	}

	let cursor = normalizeDate(new Date(filter.start_date))
	const until = normalizeDate(new Date(filter.end_date))

	while (cursor <= until) {
		buckets.push({
			key: toInputDate(cursor),
			label: formatShortDate(cursor),
			value: 0,
		})
		cursor = addDays(cursor, 1)
	}

	return buckets
}

const chartData = computed(() => {
	const buckets = buildDailyBuckets()
	const byDate = new Map(buckets.map((item) => [item.key, item]))

	for (const transaksi of transaksis.value) {
		const key = toInputDate(new Date(transaksi.created_at))
		const bucket = byDate.get(key)
		if (bucket) {
			bucket.value += transaksi.quantity
		}
	}

	return buckets
})

const sparklineValues = computed(() => chartData.value.map((item) => item.value))

const hasAnySales = computed(() => sparklineValues.value.some((value) => value > 0))

const fetchData = async () => {
	if (period.value === 'range' && (!startDate.value || !endDate.value)) {
		window.alert('Isi tanggal mulai dan tanggal akhir terlebih dahulu.')
		return
	}

	try {
		loading.value = true
		transaksis.value = await getSemuaTransaksi(getFilterByPeriod())
	} catch (error) {
		console.error('Error fetching insights data:', error)
	} finally {
		loading.value = false
	}
}

const initDefaultRange = () => {
	const today = normalizeDate(new Date())
	startDate.value = toInputDate(addDays(today, -6))
	endDate.value = toInputDate(today)
}

initDefaultRange()
fetchData()
</script>

<style scoped>
.chart-wrap {
	width: 100%;
	overflow-x: auto;
	padding-top: 8px;
}


.chart-meta-grid {
	margin-top: 14px;
	display: flex;
	flex-wrap: wrap;
	gap: 8px;
	justify-content: center;
}

.chart-meta-item {
	min-width: 62px;
	display: flex;
	flex-direction: column;
	align-items: center;
	padding: 4px 6px;
	border-radius: 8px;
	background: rgba(var(--v-theme-on-surface), 0.04);
}


.chart-meta-item strong {
	font-size: 11px;
	color: rgb(var(--v-theme-primary));
}


.chart-meta-item span {
	font-size: 11px;
	color: rgba(var(--v-theme-on-surface), 0.8);
}
</style>

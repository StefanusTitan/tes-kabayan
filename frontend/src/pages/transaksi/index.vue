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
              :disabled="loadingFormDeps || cartItems.length > 0"
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
              color="secondary"
              @click="addToCart"
            >
              <v-icon>mdi-cart-plus</v-icon>
            </v-btn>
          </v-col>
        </v-row>

        <v-divider class="my-4" />

        <div class="d-flex flex-wrap align-center justify-space-between ga-2 mb-2">
          <div class="text-body-2 text-medium-emphasis">
            Pembeli aktif: <strong>{{ activePembeliName }}</strong>
          </div>
          <div class="text-body-2 text-medium-emphasis">
            Total item: <strong>{{ cartTotalItems }}</strong> | Total nilai: <strong>{{ formatCurrency(cartTotalPrice) }}</strong>
          </div>
        </div>

        <v-table>
          <thead>
            <tr>
              <th class="text-left">Barang</th>
              <th class="text-left">Qty</th>
              <th class="text-left">Harga</th>
              <th class="text-left">Subtotal</th>
              <th class="text-left">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in cartItems" :key="item.barangId">
              <td>{{ item.barangNama }}</td>
              <td>{{ item.quantity }}</td>
              <td>{{ formatCurrency(item.harga) }}</td>
              <td>{{ formatCurrency(item.harga * item.quantity) }}</td>
              <td>
                <v-btn
                  size="small"
                  variant="text"
                  color="error"
                  icon="mdi-delete"
                  @click="removeFromCart(item.barangId)"
                />
              </td>
            </tr>
            <tr v-if="cartItems.length === 0">
              <td colspan="5" class="py-4 text-center text-medium-emphasis">
                Keranjang kosong. Tambahkan barang dulu sebelum checkout.
              </td>
            </tr>
          </tbody>
        </v-table>

        <div class="mt-4 d-flex flex-wrap justify-end ga-2">
          <v-btn
            variant="text"
            color="error"
            prepend-icon="mdi-cart-off"
            :disabled="cartItems.length === 0 || creating"
            @click="clearCart"
          >
            Kosongkan
          </v-btn>
          <v-btn
            color="primary"
            prepend-icon="mdi-content-save"
            :loading="creating"
            :disabled="cartItems.length === 0"
            @click="handleCheckout"
          >
            Simpan Semua Item
          </v-btn>
        </div>
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
import { getSemuaPembeli, type Pembeli } from '@/api/pembeli'
import { getSemuaBarang, type Barang } from '@/api/barang'

const transaksis = ref<Transaksi[]>([])
const search = ref('')
const loadingTable = ref(false)
const creating = ref(false)
const deletingId = ref<number | null>(null)
const loadingFormDeps = ref(false)

const pembelis = ref<Pembeli[]>([])
const barangs = ref<Barang[]>([])

const form = reactive({
  pembeliId: null as number | null,
  barangId: null as number | null,
  quantity: 1,
})

type CartItem = {
  barangId: number
  barangNama: string
  harga: number
  stock: number
  quantity: number
}

const cartItems = ref<CartItem[]>([])

const pembeliOptions = computed(() => {
  return pembelis.value.map((p) => ({
    label: `${p.nama} (#${p.id})`,
    value: p.id,
  }))
})

const barangOptions = computed(() => {
  return barangs.value.map((b) => ({
    label: `${b.nama} (#${b.id}) - stok ${b.stock}`,
    value: b.id,
  }))
})

const activePembeliName = computed(() => {
  const selected = pembelis.value.find((p) => p.id === form.pembeliId)
  return selected ? `${selected.nama} (#${selected.id})` : '-'
})

const cartTotalItems = computed(() => {
  return cartItems.value.reduce((total, item) => total + item.quantity, 0)
})

const cartTotalPrice = computed(() => {
  return cartItems.value.reduce((total, item) => total + (item.harga * item.quantity), 0)
})

const getErrorMessage = (error: unknown) => {
  if (error instanceof Error) {
    return error.message
  }
  return String(error)
}

const isInsufficientStockError = (error: unknown) => {
  return getErrorMessage(error).toLowerCase().includes('insufficient stock')
}

const formatDate = (value: string) => {
  return new Intl.DateTimeFormat('id-ID', {
    dateStyle: 'medium',
    timeStyle: 'short',
  }).format(new Date(value))
}

const formatCurrency = (value: number) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
  }).format(value)
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
  form.barangId = null
  form.quantity = 1
}

const addToCart = () => {
  if (!form.pembeliId || !form.barangId || form.quantity < 1) {
    window.alert('Pilih pembeli, barang, dan quantity minimal 1.')
    return
  }

  const selectedBarang = barangs.value.find((b) => b.id === form.barangId)
  if (!selectedBarang) {
    window.alert('Barang tidak ditemukan.')
    return
  }

  const existingItem = cartItems.value.find((item) => item.barangId === selectedBarang.id)
  const existingQty = existingItem ? existingItem.quantity : 0
  const requestedQty = Number(form.quantity)

  if (existingQty + requestedQty > selectedBarang.stock) {
    window.alert(`Stok ${selectedBarang.nama} tidak cukup. Maksimal ${selectedBarang.stock}.`)
    return
  }

  if (existingItem) {
    existingItem.quantity += requestedQty
  } else {
    cartItems.value.push({
      barangId: selectedBarang.id,
      barangNama: selectedBarang.nama,
      harga: selectedBarang.harga,
      stock: selectedBarang.stock,
      quantity: requestedQty,
    })
  }

  resetForm()
}

const removeFromCart = (barangId: number) => {
  cartItems.value = cartItems.value.filter((item) => item.barangId !== barangId)
}

const clearCart = () => {
  cartItems.value = []
}

const handleCheckout = async () => {
  if (!form.pembeliId) {
    window.alert('Pilih pembeli terlebih dahulu.')
    return
  }

  if (cartItems.value.length === 0) {
    window.alert('Keranjang masih kosong.')
    return
  }

  try {
    creating.value = true
    const currentCartItems = [...cartItems.value]
    const payloads = cartItems.value.map((item) => ({
      pembeli_id: Number(form.pembeliId),
      barang_id: item.barangId,
      quantity: item.quantity,
    }))

    const results = await Promise.allSettled(payloads.map((payload) => createTransaksi(payload)))

    const failedItems: CartItem[] = []
    const stockFailedItems: CartItem[] = []
    let successCount = 0

    results.forEach((result, index) => {
      if (result.status === 'fulfilled') {
        successCount += 1
      } else {
        const failedItem = currentCartItems[index]
        failedItems.push(failedItem)

        if (isInsufficientStockError(result.reason)) {
          stockFailedItems.push(failedItem)
        }

        console.error(`Error creating transaksi for item #${failedItem.barangId}:`, result.reason)
      }
    })

    cartItems.value = failedItems

    if (failedItems.length === 0) {
      window.alert(`Berhasil menyimpan ${successCount} item transaksi.`)
    } else if (stockFailedItems.length === failedItems.length) {
      const failedNames = stockFailedItems.map((item) => item.barangNama).join(', ')
      window.alert(
        `Berhasil menyimpan ${successCount} item. Stok tidak cukup untuk: ${failedNames}. ` +
        'Item yang gagal tetap ada di keranjang.',
      )
    } else {
      window.alert(`Berhasil menyimpan ${successCount} item. ${failedItems.length} item gagal dan tetap ada di keranjang.`)
    }

    await Promise.all([fetchTransaksi(), fetchFormDependencies()])
  } catch (error) {
    console.error('Error creating transaksi checkout:', error)
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

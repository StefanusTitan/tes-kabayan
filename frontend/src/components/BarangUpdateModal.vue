<template>
  <v-dialog :model-value="modelValue" max-width="640" @update:model-value="onDialogChange">
    <v-card>
      <v-card-title class="text-h6">Ubah Barang</v-card-title>
      <v-card-text>
        <v-label class="mb-1" for="nama">Nama <span class="!text-red-500">*</span></v-label>
        <v-text-field
          v-model="form.nama"
          variant="outlined"
          class="mb-3"
          required
        />
        <v-textarea
          v-model="form.deskripsi"
          label="Deskripsi"
          variant="outlined"
          class="mb-3"
          rows="3"
          auto-grow
        />
        <v-label class="mb-1" for="harga">Harga <span class="!text-red-500">*</span></v-label>
        <v-text-field
          v-model.number="form.harga"
          type="number"
          variant="outlined"
          class="mb-3"
          min="0"
        />
        <v-label class="mb-1" for="stock">Stok <span class="!text-red-500">*</span></v-label>
        <v-text-field
          v-model.number="form.stock"
          type="number"
          variant="outlined"
          min="0"
        />
      </v-card-text>
      <v-card-actions class="justify-end">
        <v-btn variant="text" :disabled="loading" @click="closeModal">Batal</v-btn>
        <v-btn color="primary" :loading="loading" :disabled="!form.nama.trim() || !selectedBarang" @click="submitForm">
          Perbarui
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { updateBarang, type Barang, type BarangPayload } from '@/api/barang'

interface Props {
  modelValue: boolean
  selectedBarang: Barang | null
}

const props = defineProps<Props>()
const emit = defineEmits<{
  (event: 'update:modelValue', value: boolean): void
  (event: 'saved'): void
}>()

const loading = ref(false)
const form = ref<BarangPayload>({
  nama: '',
  deskripsi: '',
  harga: 0,
  stock: 0
})

const closeModal = () => {
  emit('update:modelValue', false)
}

const onDialogChange = (value: boolean) => {
  emit('update:modelValue', value)
}

const submitForm = async () => {
  if (!props.selectedBarang || !form.value.nama.trim()) {
    return
  }

  try {
    loading.value = true
    await updateBarang(props.selectedBarang.id, {
      nama: form.value.nama.trim(),
      deskripsi: form.value.deskripsi.trim(),
      harga: Math.max(0, Number(form.value.harga) || 0),
      stock: Math.max(0, Number(form.value.stock) || 0)
    })
    emit('saved')
    closeModal()
  } catch (error) {
    console.error('Failed to update barang:', error)
  } finally {
    loading.value = false
  }
}

watch(
  () => [props.modelValue, props.selectedBarang] as const,
  ([isOpen, selectedBarang]) => {
    if (!isOpen || !selectedBarang) {
      return
    }

    form.value = {
      nama: selectedBarang.nama,
      deskripsi: selectedBarang.deskripsi,
      harga: selectedBarang.harga,
      stock: selectedBarang.stock
    }
  },
  { immediate: true }
)
</script>
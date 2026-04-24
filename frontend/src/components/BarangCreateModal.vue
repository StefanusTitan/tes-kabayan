<template>
  <v-dialog :model-value="modelValue" max-width="640" @update:model-value="onDialogChange">
    <v-card>
      <v-card-title class="text-h6">Tambah Barang</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="form.nama"
          label="Nama"
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
        <v-text-field
          v-model.number="form.harga"
          label="Harga"
          type="number"
          variant="outlined"
          class="mb-3"
          min="0"
        />
        <v-text-field
          v-model.number="form.stock"
          label="Stok"
          type="number"
          variant="outlined"
          min="0"
        />
      </v-card-text>
      <v-card-actions class="justify-end">
        <v-btn variant="text" :disabled="loading" @click="closeModal">Batal</v-btn>
        <v-btn color="primary" :loading="loading" :disabled="!form.nama.trim()" @click="submitForm">
          Simpan
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { createBarang, type BarangPayload } from '@/api/barang'

interface Props {
  modelValue: boolean
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

const resetForm = () => {
  form.value = {
    nama: '',
    deskripsi: '',
    harga: 0,
    stock: 0
  }
}

const closeModal = () => {
  emit('update:modelValue', false)
}

const onDialogChange = (value: boolean) => {
  emit('update:modelValue', value)
}

const submitForm = async () => {
  if (!form.value.nama.trim()) {
    return
  }

  try {
    loading.value = true
    await createBarang({
      nama: form.value.nama.trim(),
      deskripsi: form.value.deskripsi.trim(),
      harga: Math.max(0, Number(form.value.harga) || 0),
      stock: Math.max(0, Number(form.value.stock) || 0)
    })
    emit('saved')
    closeModal()
  } catch (error) {
    console.error('Failed to create barang:', error)
  } finally {
    loading.value = false
  }
}

watch(
  () => props.modelValue,
  (isOpen) => {
    if (isOpen) {
      resetForm()
    }
  }
)
</script>
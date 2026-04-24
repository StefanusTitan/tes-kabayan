<template>
  <v-dialog :model-value="modelValue" max-width="640" @update:model-value="onDialogChange">
    <v-card>
      <v-card-title class="text-h6">Tambah Barang</v-card-title>
      <v-form ref="formRef" v-model="isFormValid" @submit.prevent="submitForm">
        <v-card-text>
          <v-text-field
            v-model="form.nama"
            label="Nama"
            variant="outlined"
            class="mb-3"
            :rules="namaRules"
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
            :rules="numberRules"
          />
          <v-text-field
            v-model.number="form.stock"
            label="Stok"
            type="number"
            variant="outlined"
            min="0"
            :rules="numberRules"
          />
        </v-card-text>
        <v-card-actions class="justify-end">
          <v-btn variant="text" :disabled="loading" @click="closeModal">Batal</v-btn>
          <v-btn color="primary" type="submit" :loading="loading" :disabled="!canSubmit">
            Simpan
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { computed, nextTick, ref, watch } from 'vue'
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
const formRef = ref<{ validate: () => Promise<{ valid: boolean }>; resetValidation: () => void } | null>(null)
const isFormValid = ref(false)
const form = ref<BarangPayload>({
  nama: '',
  deskripsi: '',
  harga: 0,
  stock: 0
})

const namaRules = [(value: string) => !!value?.trim() || 'Nama wajib diisi']
const numberRules = [(value: number | string) => Number(value) >= 0 || 'Nilai tidak boleh negatif']
const canSubmit = computed(() => isFormValid.value && !loading.value)

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
  const validation = await formRef.value?.validate()
  if (!validation?.valid) {
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
  async (isOpen) => {
    if (isOpen) {
      resetForm()
      await nextTick()
      formRef.value?.resetValidation()
    }
  }
)
</script>
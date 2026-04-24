<template>
  <v-dialog :model-value="modelValue" max-width="640" @update:model-value="onDialogChange">
    <v-card>
      <v-card-title class="text-h6">Ubah Pembeli</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="form.nama"
          label="Nama"
          variant="outlined"
          class="mb-3"
          required
        />
        <v-text-field
          v-model="form.alamat"
          label="Alamat"
          variant="outlined"
          class="mb-3"
        />
        <v-text-field
          v-model="form.no_telp"
          label="No. Telepon"
          variant="outlined"
        />
      </v-card-text>
      <v-card-actions class="justify-end">
        <v-btn variant="text" :disabled="loading" @click="closeModal">Batal</v-btn>
        <v-btn color="primary" :loading="loading" :disabled="!form.nama.trim() || !selectedPembeli" @click="submitForm">
          Perbarui
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { updatePembeli, type Pembeli, type PembeliPayload } from '@/api/pembeli'

interface Props {
  modelValue: boolean
  selectedPembeli: Pembeli | null
}

const props = defineProps<Props>()
const emit = defineEmits<{
  (event: 'update:modelValue', value: boolean): void
  (event: 'saved'): void
}>()

const loading = ref(false)
const form = ref<PembeliPayload>({
  nama: '',
  alamat: '',
  no_telp: ''
})

const closeModal = () => {
  emit('update:modelValue', false)
}

const onDialogChange = (value: boolean) => {
  emit('update:modelValue', value)
}

const submitForm = async () => {
  if (!props.selectedPembeli || !form.value.nama.trim()) {
    return
  }

  try {
    loading.value = true
    await updatePembeli(props.selectedPembeli.id, {
      nama: form.value.nama.trim(),
      alamat: form.value.alamat.trim(),
      no_telp: form.value.no_telp.trim()
    })
    emit('saved')
    closeModal()
  } catch (error) {
    console.error('Failed to update pembeli:', error)
  } finally {
    loading.value = false
  }
}

watch(
  () => [props.modelValue, props.selectedPembeli] as const,
  ([isOpen, selectedPembeli]) => {
    if (!isOpen || !selectedPembeli) {
      return
    }

    form.value = {
      nama: selectedPembeli.nama,
      alamat: selectedPembeli.alamat,
      no_telp: selectedPembeli.no_telp
    }
  },
  { immediate: true }
)
</script>

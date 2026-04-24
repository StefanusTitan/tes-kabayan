<template>
  <v-dialog :model-value="modelValue" max-width="640" @update:model-value="onDialogChange">
    <v-card>
      <v-card-title class="text-h6">Ubah Pembeli</v-card-title>
      <v-form ref="formRef" v-model="isFormValid" @submit.prevent="submitForm">
        <v-card-text>
          <v-label class="mb-1" for="nama">Nama <span class="!text-red-500">*</span></v-label>
          <v-text-field
            v-model="form.nama"
            variant="outlined"
            class="mb-3"
            :rules="namaRules"
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
            :rules="noTelpRules"
          />
        </v-card-text>
        <v-card-actions class="justify-end">
          <v-btn variant="text" :disabled="loading" @click="closeModal">Batal</v-btn>
          <v-btn color="primary" type="submit" :loading="loading" :disabled="!canSubmit">
            Perbarui
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { computed, nextTick, ref, watch } from 'vue'
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
const formRef = ref<{ validate: () => Promise<{ valid: boolean }>; resetValidation: () => void } | null>(null)
const isFormValid = ref(false)
const form = ref<PembeliPayload>({
  nama: '',
  alamat: '',
  no_telp: ''
})

const namaRules = [(value: string) => !!value?.trim() || 'Nama wajib diisi']
const noTelpRules = [
  (value: string) => {
    if (!value) return true // Allow empty if not mandatory
    const pattern = /^[0-9]+$/
    return pattern.test(value) || 'Nomor telepon harus berupa angka'
  }
]
const canSubmit = computed(() => isFormValid.value && !loading.value && !!props.selectedPembeli)

const closeModal = () => {
  emit('update:modelValue', false)
}

const onDialogChange = (value: boolean) => {
  emit('update:modelValue', value)
}

const submitForm = async () => {
  if (!props.selectedPembeli) {
    return
  }

  const validation = await formRef.value?.validate()
  if (!validation?.valid) {
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
  async ([isOpen, selectedPembeli]) => {
    if (!isOpen || !selectedPembeli) {
      return
    }

    form.value = {
      nama: selectedPembeli.nama,
      alamat: selectedPembeli.alamat,
      no_telp: selectedPembeli.no_telp
    }

    await nextTick()
    formRef.value?.resetValidation()
  },
  { immediate: true }
)
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-4">
    
    <v-card class="w-full max-w-md !rounded-xl shadow-2xl">
      <div class="bg-blue-600 p-6 text-white">
        <h1 class="text-2xl font-bold">{{ isLogin ? 'Login' : 'Register' }}</h1>
      </div>

      <v-card-text class="p-8"> <v-form @submit.prevent="handleSubmit">
          <v-text-field
            v-model="formData.username"
            label="Username"
            variant="outlined"
            class="mb-4" 
          ></v-text-field>

          <v-text-field
            v-model="formData.password"
            label="Password"
            type="password"
            variant="outlined"
            class="mb-6"
          ></v-text-field>

          <v-btn 
            type="submit" 
            color="primary" 
            block 
            class="!h-12 !text-lg font-bold uppercase tracking-wider"
          >
            {{ isLogin ? 'Sign In' : 'Sign Up' }}
          </v-btn>
        </v-form>

        <div class="mt-8 text-center">
          <button 
            @click="changeState"
            class="text-blue-500 hover:text-blue-700 transition-colors font-medium"
          >
            {{ isLogin ? "Don't have an account? Register" : "Already have an account? Login" }}
          </button>
        </div>
      </v-card-text>
    </v-card>

  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { register, login } from '@/api/auth'
import { useAppStore } from '@/stores/app'
import { useRouter } from 'vue-router'

interface AuthForm {
  username: string
  password: string
}

const isLogin = useAppStore().isAuthenticated ? ref(true) : ref(false)
const router = useRouter()
const formData = ref<AuthForm>({
  username: '',
  password: ''
})

const changeState = () => {
  isLogin.value = !isLogin.value
}

const handleSubmit = async () => {
  if (isLogin.value) {
    try {
      const res = await login(formData.value.username, formData.value.password)
      router.push('/')
      console.log('Login successful:', res)
    } catch (error) {
      console.error('Login failed:', error)
    }
  } else {
    try {
      const res = await register(formData.value.username, formData.value.password)
      console.log('Registration successful:', res)
    } catch (error) {
      console.error('Registration failed:', error)
    }
  }
}
</script>
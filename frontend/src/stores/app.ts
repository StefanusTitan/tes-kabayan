// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    isAuthenticated: false,
    user: null as { id: number; username: string } | null,
  }),
})

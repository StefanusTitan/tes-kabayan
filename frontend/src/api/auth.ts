// src/api/auth.ts
import { api } from "@/api/api";
import { useAppStore } from "@/stores/app";

export const login = async (username: string, password: string) => {
  try {
    const response = await api.post('/auth/login', { username, password });
    const appStore = useAppStore();
    appStore.isAuthenticated = true;
    return response;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

export const logout = async () => {
  try {
    await api.post('/auth/logout');
    const appStore = useAppStore();
    appStore.isAuthenticated = false;
  } catch (error) {
    console.error('Logout failed:', error);
    throw error;
  }
};

export const register = async (username: string, password: string) => {
  try {
    const response = await api.post('/auth/register', { username, password });
    return response;
  } catch (error) {
    console.error('Registration failed:', error);
    throw error;
  }
};
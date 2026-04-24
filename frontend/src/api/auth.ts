// src/api/auth.ts
import { api } from "@/api/api";
import { useAppStore } from "@/stores/app";

type AuthUser = { id: number; username: string };
type LoginResponse = { message: string; data: AuthUser };
type MeResponse = { data: AuthUser };

export const login = async (username: string, password: string) => {
  try {
    const response = await api.post<LoginResponse>('/auth/login', { username, password });
    const appStore = useAppStore();
    appStore.isAuthenticated = true;
    appStore.user = response?.data ?? null;
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
    appStore.user = null;
  } catch (error) {
    console.error('Logout failed:', error);
    throw error;
  }
};

export const getCurrentUser = async () => {
  try {
    const response = await api.get<MeResponse>('/auth/me');
    const appStore = useAppStore();
    appStore.isAuthenticated = true;
    appStore.user = response?.data ?? null;
    return response;
  } catch (error) {
    const appStore = useAppStore();
    appStore.isAuthenticated = false;
    appStore.user = null;
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
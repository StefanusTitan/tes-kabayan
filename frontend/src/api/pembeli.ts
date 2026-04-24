import { api } from "@/api/api";

export interface Pembeli {
  id: number;
  nama: string;
  alamat: string;
  no_telp: string;
}

export interface PembeliPayload {
  nama: string;
  alamat: string;
  no_telp: string;
}

export const getSemuaPembeli = async (search = ""): Promise<Pembeli[]> => {
  try {
    const params = search.trim() ? { search } : undefined;
    const response = await api.get<Pembeli[]>('/pembeli', { params });
    return response ?? [];
  } catch (error) {
    console.error('Failed to fetch pembeli:', error);
    throw error;
  }
};

export const createPembeli = async (data: PembeliPayload): Promise<Pembeli> => {
  try {
    const response = await api.post<Pembeli>('/pembeli', data);
    if (!response) {
      throw new Error('Failed to create pembeli: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to create pembeli:', error);
    throw error;
  }
};

export const updatePembeli = async (id: number, data: PembeliPayload): Promise<Pembeli> => {
  try {
    const response = await api.put<Pembeli>(`/pembeli/${id}`, data);
    if (!response) {
      throw new Error('Failed to update pembeli: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to update pembeli:', error);
    throw error;
  }
}

export const deletePembeli = async (id: number) => {
  try {
    await api.delete(`/pembeli/${id}`);
  } catch (error) {
    console.error('Failed to delete pembeli:', error);
    throw error;
  }
}
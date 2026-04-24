import { api } from "@/api/api";

export interface Barang {
  id: number;
  nama: string;
  deskripsi: string;
  harga: number;
  stock: number;
}

export interface BarangPayload {
  nama: string;
  deskripsi: string;
  harga: number;
  stock: number;
}

export const getSemuaBarang = async (search = ""): Promise<Barang[]> => {
  try {
    const params = search.trim() ? { search } : undefined;
    const response = await api.get<Barang[]>('/barang', { params });
    return response ?? [];
  } catch (error) {
    console.error('Failed to fetch barang:', error);
    throw error;
  }
};

export const createBarang = async (data: BarangPayload): Promise<Barang> => {
  try {
    const response = await api.post<Barang>('/barang', data);
    if (!response) {
      throw new Error('Failed to create barang: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to create barang:', error);
    throw error;
  }
};

export const updateBarang = async (id: number, data: BarangPayload): Promise<Barang> => {
  try {
    const response = await api.put<Barang>(`/barang/${id}`, data);
    if (!response) {
      throw new Error('Failed to update barang: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to update barang:', error);
    throw error;
  }
}

export const deleteBarang = async (id: number) => {
  try {
    await api.delete(`/barang/${id}`);
  } catch (error) {
    console.error('Failed to delete barang:', error);
    throw error;
  }
}
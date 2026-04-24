import { api } from "@/api/api";

export interface TransaksiItem {
  id: number;
  transaksi_id: number;
  produk_id: number;
  jumlah: number;
  harga: number;
}

export interface Transaksi {
  id: number;
  pembeli_id: number;
  tanggal: string;
  total: number;
  items: TransaksiItem[];
}

export interface TransaksiPayload {
  pembeli_id: number;
  tanggal: string;
  items: {
    produk_id: number;
    jumlah: number;
    harga: number;
  }[];
}

export interface TransaksiFilter {
  search?: string;
  pembeli_id?: number;
  start_date?: string; // format: YYYY-MM-DD
  end_date?: string;   // format: YYYY-MM-DD
}

export const getSemuaTransaksi = async (filter: TransaksiFilter = {}): Promise<Transaksi[]> => {
  try {
    const params: Record<string, string> = {};
    if (filter.search?.trim()) params.search = filter.search.trim();
    if (filter.pembeli_id) params.pembeli_id = String(filter.pembeli_id);
    if (filter.start_date) params.start_date = filter.start_date;
    if (filter.end_date) params.end_date = filter.end_date;

    const response = await api.get<Transaksi[]>('/transaksi', {
      params: Object.keys(params).length ? params : undefined,
    });
    return response ?? [];
  } catch (error) {
    console.error('Failed to fetch transaksi:', error);
    throw error;
  }
};

export const getTransaksiByID = async (id: number): Promise<Transaksi> => {
  try {
    const response = await api.get<Transaksi>(`/transaksi/${id}`);
    if (!response) {
      throw new Error('Failed to fetch transaksi: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to fetch transaksi:', error);
    throw error;
  }
};

export const createTransaksi = async (data: TransaksiPayload): Promise<Transaksi> => {
  try {
    const response = await api.post<Transaksi>('/transaksi', data);
    if (!response) {
      throw new Error('Failed to create transaksi: empty response');
    }
    return response;
  } catch (error) {
    console.error('Failed to create transaksi:', error);
    throw error;
  }
};

export const deleteTransaksi = async (id: number): Promise<void> => {
  try {
    await api.delete(`/transaksi/${id}`);
  } catch (error) {
    console.error('Failed to delete transaksi:', error);
    throw error;
  }
};

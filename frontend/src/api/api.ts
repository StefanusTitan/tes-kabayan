const BASE_URL: string = import.meta.env.VITE_API_URL;

if (!BASE_URL) {
  console.error('VITE_API_URL is not set. Check frontend/.env and restart the dev server.');
}

console.log('API Base URL:', BASE_URL);

interface FetchOptions extends Omit<RequestInit, 'body'> {
  body?: unknown;
  params?: Record<string, string | number | boolean | null | undefined>;
}

type ApiErrorBody = {
  error?: string;
  message?: string;
  detail?: string;
}

async function fetchClient<T>(endpoint: string, options: FetchOptions = {}): Promise<T | null> {
  const { body, params, ...requestOptions } = options;

  let url = `${BASE_URL}${endpoint}`;
  if (params) {
    const query = new URLSearchParams();
    for (const [key, value] of Object.entries(params)) {
      if (value !== undefined && value !== null) {
        query.append(key, String(value));
      }
    }
    const queryString = query.toString();
    if (queryString) {
      url = `${url}${url.includes('?') ? '&' : '?'}${queryString}`;
    }
  }

  const defaultHeaders: HeadersInit = {
    'Content-Type': 'application/json',
  };

  const config: RequestInit = {
    ...requestOptions,
    credentials: 'include',
    headers: {
      ...defaultHeaders,
      ...(options.headers as HeadersInit),
    },
  };

  if (body && typeof body === 'object') {
    config.body = JSON.stringify(body);
  } else if (body) {
    config.body = body as BodyInit;
  }

  try {
    const response = await fetch(url, config);

    if (!response.ok) {
      const errorData = await response.json().catch(() => null) as ApiErrorBody | null;
      const errorMessage =
        errorData?.error ||
        errorData?.message ||
        errorData?.detail ||
        `Error: ${response.status} ${response.statusText}`;
      const apiError = new Error(errorMessage) as Error & { status: number; data: ApiErrorBody | null };
      apiError.status = response.status;
      apiError.data = errorData;
      throw apiError;
    }

    if (response.status === 204) {
      return null;
    }

    return (await response.json()) as T;

  } catch (error) {
    const err = error as Error;
    console.error(`[API Error] ${options.method || 'GET'} ${endpoint}:`, err.message);
    throw err;
  }
}

export const api = {
  get: <T>(
    endpoint: string,
    options?: {
      headers?: HeadersInit;
      params?: Record<string, string | number | boolean | null | undefined>;
    },
  ) => fetchClient<T>(endpoint, { method: 'GET', headers: options?.headers, params: options?.params }),

  post: <T>(endpoint: string, body?: unknown, headers?: HeadersInit) => 
    fetchClient<T>(endpoint, { method: 'POST', body, headers }),

  put: <T>(endpoint: string, body?: unknown, headers?: HeadersInit) => 
    fetchClient<T>(endpoint, { method: 'PUT', body, headers }),

  delete: <T>(endpoint: string, headers?: HeadersInit) => 
    fetchClient<T>(endpoint, { method: 'DELETE', headers }),
};
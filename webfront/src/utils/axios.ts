import type {AxiosRequestConfig} from 'axios';
import axios from 'axios';

/**
 * 根据运行环境获取基础请求URL
 */
export const getUrl = (): string => {
  const value: string = "http://localhost:8081/api"
  return value == 'getCurrentDomain' ? window.location.protocol + '//' + window.location.host : value
}

export function createAxios<Data = any, T = ApiPromise<Data>>(
    axiosConfig: AxiosRequestConfig,
    mock: boolean = false,
): T {
  const axiosInstance = axios.create({
    baseURL: mock ? window.location.protocol + '//' + window.location.host : getUrl(),
    timeout: 1000 * 10,
    headers: {},
    responseType: 'json',
  });
  return axiosInstance(axiosConfig) as unknown as T;
}

interface ApiResponse<T = any> {
  status: number;
  data: T;
  statusText: string;
  time: number;
}

type ApiPromise<T = any> = Promise<ApiResponse<T>>;
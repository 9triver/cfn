import type { AxiosRequestConfig } from 'axios';
import axios from 'axios';

/**
 * 根据运行环境获取基础请求URL
 */
export const getUrl = (): string => {
    const value: string = import.meta.env.VITE_AXIOS_BASE_URL as string
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
    return axiosInstance(axiosConfig) as T;
}

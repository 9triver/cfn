import {createAxios} from "@/utils/axios";

export const getTotalResourceCount = () => {
    return createAxios({
        url: '/resources/total',
        method: 'get'
    });
}

export const getAvailableResourceCount = () => {
    return createAxios({
        url: '/resources/available',
        method: 'get'
    });
}
import axios, { AxiosRequestConfig } from "axios";
import { useQuery, useMutation } from '@tanstack/react-query';
import { storageGetValue } from "./utils/storage";

const backendUrl = process.env.REACT_APP_BACKEND_URL;

const config: AxiosRequestConfig = {
    withCredentials: true,
    baseURL: backendUrl,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Access-Control-Allow-Credentials': 'true',
    },
};

const http = axios.create(config);

http.interceptors.request.use(async (conf) => {
    let c = { ...conf };
    const token = await storageGetValue("token");
    if (token) {
        c = {
            ...c,
            headers: {
                ...c.headers,
                Authorization: "Bearer " + token,
            }
        }
    }
    return c;
});

interface ISignupResponse {
    token: string;
}

export const useSignup = () => {
    return useMutation(["signup"], async (params: {
        username: string,
        password: string,
    }) => {
        const response = await http.post<ISignupResponse>("/sup", {
            uname: params.username,
            pword: params.password,
        });
        return response.data;
    });
};

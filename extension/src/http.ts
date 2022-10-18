import axios, { AxiosRequestConfig } from "axios";
import { useQuery, useMutation } from '@tanstack/react-query';

const backendUrl = "http://localhost:8080";

const config: AxiosRequestConfig = {
    withCredentials: false,
    baseURL: backendUrl,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    //   'Access-Control-Allow-Credentials': 'true',
    },
};

const http = axios.create(config);

http.interceptors.request.use((conf) => {
    return {
        ...conf,
    };
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

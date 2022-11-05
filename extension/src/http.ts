import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
    IAddCommentResponse,
    IAddReplyResponse,
    ICreateTopicResponse,
    IGetCommentsResponse,
    IGetReactionsResponse,
    IGetRepliesResponse,
    IGetTopicsResponse,
    IPostReactionResponse,
    ISignupResponse
} from "./interfaces/model";
import { ChromeExtension } from "./data-structure/chrome-extension";

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

const chromeExt = new ChromeExtension();

http.interceptors.request.use(async (conf) => {
    let c = { ...conf };
    const token = await chromeExt.storageGetValue("token");
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

http.interceptors.response.use(async (response: AxiosResponse) => {
    return response;
}, async (err) => {
    if (err.response.status === 401) {
        console.log("Session expired");
        await chromeExt.storageClear();
        console.log("Storage has been cleared");
    }
    return Promise.reject(err.response);
});

export const useSignup = () => {
    return useMutation(["post-signup"], async (params: {
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



export const useGetTopics = () => {
    return useQuery(["get-topics"], async () => {
        const {data} = await http.get<IGetTopicsResponse>("/tpc");
        return data;
    });
};

export const useCreateTopic = () => {
    return useMutation(["post-create-topic"], async (params: {
        url: string,
        title: string,
        comment?: string,
    }) => {
        let body: any = {
            url: params.url,
            title: params.title,
        }
        if (params.comment) {
            body = {
                ...body,
                comment: params.comment,
            };
        }
        const {data} = await http.post<ICreateTopicResponse>("/tpc", body);
        return data;
    });
};

export const useGetComments = (topicId: string, enabled: boolean = true) => {
    return useQuery(["get-comments", topicId], async () => {
        const url = `/tpc/${topicId}/cmt`;
        const {data} = await http.get<IGetCommentsResponse>(url);
        return data;
    }, {
        enabled: enabled,
    });
};

export const useAddComment = (topicId: string) => {
    const queryClient = useQueryClient();
    return useMutation(["post-add-comment", topicId], async (params: {
        comment: string,
    }) => {
        const url = `/tpc/${topicId}/cmt`;
        const {data} = await http.post<IAddCommentResponse>(url, {
            comment: params.comment,
        });
        return data;
    }, {
        onSuccess: () => {
            queryClient.invalidateQueries(["get-comments", topicId]);
        }
    });
};

export const useGetReplies = (commentId: string) => {
    return useQuery(["get-replies", commentId], async () => {
        const url = `/cmt/${commentId}`;
        const {data} = await http.get<IGetRepliesResponse>(url);
        return data;
    });
};

export const useAddReply = (commentId: string) => {
    const queryClient = useQueryClient();
    return useMutation(["post-add-reply", commentId], async (params: {
        comment: string,
    }) => {
        const url = `/cmt/${commentId}`;
        const {data} = await http.post<IAddReplyResponse>(url, {
            comment: params.comment,
        });
        return data;
    }, {
        onSuccess: () => {
            queryClient.invalidateQueries(["get-replies", commentId]);
        }
    });
};

export const useGetRactions = (commentId: string) => {
    return useQuery(["get-reactions", commentId], async () => {
        const url = `/rctn/${commentId}`;
        const {data} = await http.get<IGetReactionsResponse>(url);
        return data;
    });
};

export const usePostReaction = (commentId: string) => {
    const queryClient = useQueryClient();
    return useMutation(["post-reactions", commentId], async (params: {
        reaction: number;
    }) => {
        const url = `/rctn/${commentId}`;
        const {data} = await http.post<IPostReactionResponse>(url, {
            reaction: params.reaction
        });
        return data;
    }, {
        onSuccess: () => {
            queryClient.invalidateQueries(["get-reactions", commentId]);
        }
    });
};

export const useDeleteReaction = (commentId: string) => {
    const queryClient = useQueryClient();
    return useMutation(["delete-reactions", commentId], async (params: {
        reaction: number;
    }) => {
        const url = `/rctn/${commentId}?r=${params.reaction}`;
        await http.delete<null>(url);
        return null;
    }, {
        onSuccess: () => {
            queryClient.invalidateQueries(["get-reactions", commentId]);
        }
    });
};

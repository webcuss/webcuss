export interface IUser {
    id: string;
    uname: string;
}

export interface IC5t {
    comment: string;
    createdOn: string;
    user: IUser;
}

export interface ITopic {
    commentsCount: number;
    hostname: string;
    id: string;
    likes: number;
    path: string;
    query: string;
    title: string;
    user: IUser;
}

export interface IGetTopicsResponse {
    data: ITopic[];
}

export interface ISignupResponse {
    token: string;
}

export interface ICreateTopicResponse {
    id: string;
    commentId?: string;
}

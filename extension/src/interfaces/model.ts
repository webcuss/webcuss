export interface IUser {
    id: string;
    uname: string;
}

export interface IC5t {
    id: string;
    content: string;
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

export interface IGetCommentsResponse {
    data: IC5t[];
}

export interface IAddCommentResponse {
    id: string;
}

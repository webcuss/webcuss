import React, { useEffect, useState } from "react";
import styled from "styled-components";
import C5t from "../../components/c5t/C5t";
import CommentArea from "../../components/comment-area/CommentArea";
import SignupSuggestion from "../../components/signup-suggestion/SignupSuggestion";
import T8y from "../../components/t8y/T8y";
import { useAuth } from "../../hooks/useAuth";
import { useBrowserExtension } from "../../hooks/useBrowserExtension";
import { useCreateTopic, useGetComments } from "../../http";
import { IC5t } from "../../interfaces/model";
import { b } from "../../utils/bool";
import { s } from "../../utils/string";

const Main = () => {
    const {isSignedIn: hIsSignedIn} = useAuth();
    const {chromeExt} = useBrowserExtension();
    const createTopic = useCreateTopic();

    const [isSignedIn, setIsSignedIn] = useState<boolean>(hIsSignedIn);
    const [topicId, setTopicId] = useState<string|undefined>(undefined);
    const [comments, setComments] = useState<IC5t[]>([]);

    const {data: hComments} = useGetComments(s(topicId), b(topicId) && isSignedIn);

    useEffect(() => {
        if (hComments && hComments.data) {
            setComments(hComments.data);
        }
    }, [hComments]);

    useEffect(() => {
        setIsSignedIn(hIsSignedIn);
    }, [hIsSignedIn]);

    useEffect(() => {
        (async () => {
            const pUrl = await chromeExt.getPageUrl();
            const pTitle = await chromeExt.getPageTitle();
            if (pUrl && pTitle) {
                const res = await createTopic.mutateAsync({
                    url: pUrl,
                    title: pTitle,
                });
                setTopicId(res.id);
            }
        })();
    }, [chromeExt]);

    return (
        <Root>
            <CommentArea topicId={s(topicId)} />
            <T8y text={comments.length + " comments"} />

            {isSignedIn && (
                <>
                    {comments.map((c, i) => <C5t key={i} data={c} />)}
                </>
            )}
            {!isSignedIn && (
                <div className="mt-20">
                    <T8y text="Complete your profile to add comments" />
                    <SignupSuggestion className="mt-10" />
                </div>
            )}
        </Root>
    );
};

export default Main;

const Root = styled.div`
    height: calc(var(--html-height) - var(--body-padding-top) - var(--body-padding-bottom));
    overflow-y: auto;
`;

import React, { useState } from "react";
import { P } from "../../interfaces/common";
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styled from "styled-components";
import { useAddComment } from "../../http";

interface ICommentAreaProps extends P {
    topicId: string;
}

const CommentArea = (p: ICommentAreaProps) => {
    const [btnHover, setBtnHover] = useState<boolean>(false);
    const [commentInput, setCommentInput] = useState<string>("");

    const addComment = useAddComment(p.topicId);

    const textChangeHandler = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        const value = e.currentTarget.value;
        setCommentInput(value);
    };

    const sendCommentHandler = async () => {
        if (p.topicId && commentInput) {
            const res = await addComment.mutateAsync({comment: commentInput});
            setCommentInput("");
        }
    };

    return (
        <Root className="w-full flex flex-row ai-center">
            <TextArea value={commentInput} className="flex-1" rows={2} autoFocus onChange={textChangeHandler} />
            <SendButtonWrapper
                onMouseEnter={() => setBtnHover(true)}
                onMouseLeave={() => setBtnHover(false)}
                onClick={sendCommentHandler}>
                <FontAwesomeIcon color={btnHover ? "#2874A6" : "#2E86C1"} size="lg" icon={faPaperPlane} />
            </SendButtonWrapper>
        </Root>
    );
};

export default CommentArea;

const Root = styled.div`
    margin-top: 2px;
`;

const TextArea = styled.textarea`
    min-height: 40px;
    max-height: 80px;
`;

const SendButtonWrapper = styled.div`
    width: 21px;
    padding: 3px 0 3px 3px;
    margin-left: 3px;

    &:hover {
        cursor: pointer;
    }
`;

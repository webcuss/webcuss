import React, { useState } from "react";
import styled from "styled-components";
import { faReply } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { isCtrlEnter } from "../../utils/keyboard";
import { useAddReply } from "../../http";

interface IReplyAreaProps {
    commentId: string;
}

const ReplyArea = (p: IReplyAreaProps) => {
    const [replyInput, setReplyInput] = useState<string>("");

    const addReply = useAddReply(p.commentId);

    const textChangeHandler = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        const value = e.currentTarget.value;
        setReplyInput(value);
    };

    const sendReplyHandler = async () => {
        if (p.commentId && replyInput) {
            await addReply.mutateAsync({comment: replyInput});
            setReplyInput("");
        }
    };

    const keyDownHandler = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (isCtrlEnter(e)) {
            sendReplyHandler();
        }
    };

    return (
        <Root className="w-full flex flex-row ai-center">
            <TextArea
                value={replyInput}
                className="flex-1"
                rows={1}
                autoFocus
                placeholder="Add reply (Ctrl+Enter to send)"
                onChange={textChangeHandler}
                onKeyDown={keyDownHandler} />
            <SendButtonWrapper
                onClick={sendReplyHandler}>
                <FontAwesomeIcon color="#2E86C1" size="lg" icon={faReply} />
            </SendButtonWrapper>
        </Root>
    );
};

export default ReplyArea;

const Root = styled.div``;

const TextArea = styled.textarea`
    min-height: 30px;
    max-height: 50px;
`;

const SendButtonWrapper = styled.div`
    width: 21px;
    padding: 3px 0 3px 3px;
    margin-left: 3px;

    &:hover {
        cursor: pointer;
    }
`;

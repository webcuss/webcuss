import React, { useEffect, useRef, useState } from "react";
import styled from "styled-components";
import { isCtrlEnter } from "../../utils/keyboard";
import { useAddReply } from "../../http";
import { format } from "date-fns";
import { randomNumber } from "../../utils/random";
import { useAuth } from "../../hooks/useAuth";

interface IReplyAreaProps {
    commentId: string;
    onCancel?: () => void;
}

const ReplyArea = (p: IReplyAreaProps) => {
    const {uname} = useAuth();

    const [randId] = useState<string>(`id${randomNumber(99, 9999)}`);
    const [replyInput, setReplyInput] = useState<string>("");

    const addReply = useAddReply(p.commentId);

    useEffect(() => {
        const elem = document.getElementById(randId);
        if (elem) {
            const range = document.createRange();
            const sel = window.getSelection();
            
            range.setStart(elem, 0);
            range.collapse(true);
            
            if (sel) {
                sel.removeAllRanges();
                sel.addRange(range);
            }
        }
    }, []);

    const textChangeHandler = (e: React.ChangeEvent<HTMLSpanElement>) => {
        const value = e.currentTarget.innerHTML;
        setReplyInput(value);
    };

    const sendReplyHandler = async () => {
        if (p.commentId && replyInput) {
            await addReply.mutateAsync({comment: replyInput});
            const elem = document.getElementById(randId);
            if (elem) {
                elem.innerHTML = "";
            }
        }
    };

    const keyDownHandler = (e: React.KeyboardEvent<HTMLSpanElement>) => {
        if (isCtrlEnter(e)) {
            sendReplyHandler();
        }
    };

    const cancelHandler = () => {
        setTimeout(() => {
            if (p.onCancel) {
                p.onCancel();
            }
        }, 200);
    };

    return (
        <Root>
            <div>
                <b>{uname}</b>
                <StyledDate>{format(new Date(), "MM/dd/yyyy")}</StyledDate>
            </div>
            <StyledReplyInput
                id={randId}
                contentEditable
                onInput={textChangeHandler}
                onBlur={cancelHandler}
                onKeyDown={keyDownHandler} />
            <ReplyAssistant>{replyInput ? "" : "Start typing "}<i>(Ctrl+Enter to send)</i></ReplyAssistant>
        </Root>
    );
};

export default ReplyArea;

const Root = styled.div`
    margin-left: 5px;
    border-left: 1px solid #C6C6C6;
    padding-left: 3px;
`;

const StyledDate = styled.span`
    color: var(--gray);
    font-size: 80%;
`;

const StyledReplyInput = styled.span`
    padding: 2px;
    border:0;
    outline:0;

    &:focus {
        outline: none !important;
    }
`;

const ReplyAssistant = styled.span`
    color: var(--gray);
    margin-left: 3px;
`;

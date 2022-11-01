import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { useGetReplies } from "../../http";
import { P } from "../../interfaces/common";
import { IC5t } from "../../interfaces/model";
import Reply from "../reply/Reply";
import T8y from "../t8y/T8y";

interface RepliesProps extends P {
    commentId: string;
}

const Replies = (p: RepliesProps) => {
    const [allReplies, setAllReplies] = useState<IC5t[]>([]);
    const [replies, setReplies] = useState<IC5t[]>([]);
    const [showAllReplies, setShowAllReplies] = useState<boolean>(false);

    const {data: hReplies} = useGetReplies(p.commentId);

    useEffect(() => {
        if (hReplies && hReplies.data) {
            setAllReplies(hReplies.data);
        }
    }, [hReplies]);

    useEffect(() => {
        if (showAllReplies) {
            setReplies([...allReplies]);
        } else {
            if (allReplies.length > 0) {
                setReplies([allReplies[0]]);
            } else {
                setReplies([]);
            }
        }
    }, [showAllReplies, allReplies]);

    const showAllReplyClickHandler = () => {
        setShowAllReplies(true);
    };

    const hideExtraRepliesClickHandler = () => {
        setShowAllReplies(false);
    };

    if (replies.length < 1) {
        return (<></>);
    }

    return (
        <Root>
            {replies.length > 0 && replies.map((r, i) => <Reply key={i} data={r} />)}
            {allReplies.length > 1 && (
                <div>
                    <ShowHideReplies onClick={showAllReplies ? hideExtraRepliesClickHandler : showAllReplyClickHandler}>
                        <T8y text={showAllReplies ? "Hide replies" : "Show all"} />
                    </ShowHideReplies>
                </div>
            )}
        </Root>
    );
};

export default Replies;

const Root = styled.div`
    margin-left: 5px;
    border-left: 1px solid #C6C6C6;
    padding-left: 3px;
`;

const ShowHideReplies = styled.span`
    font-size: 90%;
    color: #2E86C1;

    &:hover {
        cursor: pointer;
    }
`;

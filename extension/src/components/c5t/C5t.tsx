import { format, parseISO } from "date-fns";
import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { useDeleteReaction, useGetRactions, usePostReaction } from "../../http";
import { P } from "../../interfaces/common";
import { IC5t } from "../../interfaces/model";
import ActionLike from "../action-like/ActionLike";
import ActionReply from "../action-reply/ActionReply";
import Replies from "../replies/Replies";
import ReplyArea from "../reply-area/ReplyArea";

interface C5tProps extends P {
    data: IC5t;
}

enum Reaction {
    // Minimum values is 0
    LIKE = 1,
}

const C5t = (p: C5tProps) => {
    const [isReplying, setIsReplying] = useState<boolean>(false);
    const [likeCount, setLikeCount] = useState<number>(0);
    const [liked, setLiked] = useState<boolean>(false);
    const [updatingReaction, setUpdatingReaction] = useState<boolean>(false);

    const {data: hReactions} = useGetRactions(p.data.id);
    const postReaction = usePostReaction(p.data.id);
    const deleteReaction = useDeleteReaction(p.data.id);

    useEffect(() => {
        if (hReactions) {
            const isLiked = hReactions.user.find(v => v === Reaction.LIKE) !== undefined;
            setLiked(isLiked);

            const reactionLike = hReactions.all.find(v => v.reaction === Reaction.LIKE);
            if (reactionLike) {
                setLikeCount(reactionLike.count);
            }
        }
    }, [hReactions]);

    const replyClickHandler = () => {
        setIsReplying(true);
    };

    const cancelReplyHandler = () => {
        setIsReplying(false);
    };

    const likeClickHandler = async () => {
        if (!updatingReaction) {
            try {
                setUpdatingReaction(true);
                if (liked) {
                    await deleteReaction.mutateAsync({reaction: Reaction.LIKE});
                    console.log("reaction deleted");
                } else {
                    const res = await postReaction.mutateAsync({reaction: Reaction.LIKE});
                    console.log({res});
                }
            } finally {
                setUpdatingReaction(false);
            }
        }
    };

    return (
        <Root>
            <div>
                <b>{p.data.user.uname + " "}</b>
                <StyleDate>{format(parseISO(p.data.createdOn), "MM/dd/yyyy")}</StyleDate>
            </div>

            <div>{p.data.content}</div>

            <Actions>
                <ActionLike count={likeCount} liked={liked} onClick={likeClickHandler} />
                <ActionReply onClick={replyClickHandler} />
            </Actions>

            {isReplying && (<ReplyArea commentId={p.data.id} onCancel={cancelReplyHandler} />)}

            <Replies commentId={p.data.id} />
        </Root>
    );
};

export default C5t;

const Root = styled.div`
    margin-bottom: 10px;
`;

const StyleDate = styled.span`
    color: var(--gray);
    font-size: 80%;
`;

const Actions = styled.div`
    margin-top: 3px;
    margin-bottom: 5px;
    display: flex;
    flex-direction: row;
    align-items: center;

    &>*:not(:last-child) {
        margin-right: 10px;
    }
`;

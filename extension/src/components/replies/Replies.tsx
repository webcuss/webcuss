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

const defaultVisibleReplies: number = 2;
const showMoreRepliesIncrement: number = 3;

const Replies = (p: RepliesProps) => {
    const [allReplies, setAllReplies] = useState<IC5t[]>([]);
    const [replies, setReplies] = useState<IC5t[]>([]);
    const [hasMoreReplies, setHasMoreReplies] = useState<boolean>(true);

    const {data: hReplies} = useGetReplies(p.commentId);

    useEffect(() => {
        if (hReplies && hReplies.data) {
            setAllReplies(hReplies.data);
            if (hReplies.data.length > defaultVisibleReplies) {
                const initialReplies = hReplies.data.slice(0, defaultVisibleReplies);
                setReplies(initialReplies);
                setHasMoreReplies(true);
            } else {
                setReplies(hReplies.data);
                setHasMoreReplies(false);
            }
        }
    }, [hReplies]);

    const seeMoreClickHandler = () => {
        if (allReplies.length > replies.length) {
            if (replies.length + showMoreRepliesIncrement <= allReplies.length) {
                const moreReplies = allReplies.slice(replies.length, replies.length + showMoreRepliesIncrement);
                setReplies(pv => [...pv, ...moreReplies]);
                setHasMoreReplies(replies.length + showMoreRepliesIncrement < allReplies.length);
            } else {
                const remainingReplies = allReplies.slice(replies.length, allReplies.length);
                setReplies(pv => [...pv, ...remainingReplies]);
                setHasMoreReplies(false);
            }
        }
    };

    return (
        <Root>
            {replies.map((r, i) => <Reply key={i} data={r} />)}
            {hasMoreReplies && (
                <div>
                    <SeeMore onClick={seeMoreClickHandler}>
                        <T8y text="See more" />
                    </SeeMore>
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

const SeeMore = styled.span`
    font-size: 90%;
    color: #2E86C1;

    &:hover {
        cursor: pointer;
    }
`;

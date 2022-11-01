import React, { useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";
import { IC5t } from "../../interfaces/model";
import Reply from "../reply/Reply";
import T8y from "../t8y/T8y";

const data: IC5t[] = [
    {
        id: "",
        content: "are you dev?",
        createdOn: "2021-10-02T06:04:22.000Z",
        user: {
            uname: "danny",
            id: ""
        }
    },
    {
        id: "",
        content: "The quick brown fox jumps over the lazy dog",
        createdOn: "2021-10-02T06:04:22.000Z",
        user: {
            uname: "jim",
            id: ""
        }
    },
    {
        id: "",
        content: "indeed 🤣",
        createdOn: "2021-10-02T06:04:22.000Z",
        user: {
            uname: "kim",
            id: ""
        }
    },
    {
        id: "",
        content: "agree!",
        createdOn: "2021-10-02T06:04:22.000Z",
        user: {
            uname: "jonnah",
            id: ""
        }
    },
];

interface RepliesProps extends P {
    data: IC5t;
}

const defaultVisibleReplies: number = 2;

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

const Replies = (p: RepliesProps) => {
    const [replies, setReplies] = useState<IC5t[]>(data.slice(0, defaultVisibleReplies));
    const [hasMoreReplies, setHasMoreReplies] = useState<boolean>(true);

    const seeMoreClickHandler = () => {
        setReplies(prev => [...prev, ...data.slice(defaultVisibleReplies, data.length)]);
        setHasMoreReplies(false);
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

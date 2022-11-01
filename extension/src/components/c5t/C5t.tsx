import { format, parseISO } from "date-fns";
import React, { useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";
import { IC5t } from "../../interfaces/model";
import ActionLike from "../action-like/ActionLike";
import ActionReply from "../action-reply/ActionReply";
import Replies from "../replies/Replies";

interface C5tProps extends P {
    data: IC5t;
}

const C5t = (p: C5tProps) => {
    return (
        <Root>
            <div>
                <b>{p.data.user.uname + " "}</b>
                <StyleDate>{format(parseISO(p.data.createdOn), "MM/dd/yyyy")}</StyleDate>
            </div>

            <div>{p.data.content}</div>

            <Actions>
                <ActionLike count={69} />
                <ActionReply />
            </Actions>

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

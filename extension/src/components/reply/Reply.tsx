import { format, parseISO } from "date-fns";
import React from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";
import { IC5t } from "../../interfaces/model";

interface ReplyProps extends P {
    data: IC5t;
}

const Root = styled.div`
    margin-bottom: 3px;
`;

const Date = styled.span`
    color: var(--gray);
    font-size: 80%;
`;

const Reply = (p: ReplyProps) => {
    return (
        <Root>
            <div>
                <b>{p.data.user.uname + " "}</b>
                <Date>{format(parseISO(p.data.createdOn), "MM/dd/yyyy")}</Date>
            </div>
            <div>{p.data.content}</div>
        </Root>
    );
};

export default Reply;

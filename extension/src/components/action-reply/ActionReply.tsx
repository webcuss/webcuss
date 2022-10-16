import { faComment} from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from "react";
import styled from "styled-components";

const Root = styled.div`
    display: flex;
    flex-direction: row;

    &:hover {
        cursor: pointer;
    }
`;

const Txt = styled.span`
    font-size: 80%;
    margin-left: 3px;
`;

const ActionReply = () => {
    return (
        <Root>
            <FontAwesomeIcon icon={faComment} />
            <Txt>reply</Txt>
        </Root>
    );
};

export default ActionReply;

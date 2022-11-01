import { faComment} from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from "react";
import styled from "styled-components";
import { P } from '../../interfaces/common';

interface IActionReplyProps extends P {
    onClick?: () => void;
}

const ActionReply = (p: IActionReplyProps) => {
    const clickHandler = () => {
        if (p.onClick) {
            p.onClick();
        }
    };

    return (
        <Root onClick={clickHandler}>
            <FontAwesomeIcon icon={faComment} />
            <Txt>Reply</Txt>
        </Root>
    );
};

export default ActionReply;

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

import { faThumbsUp } from "@fortawesome/free-regular-svg-icons";
import { faThumbsUp as faThumbsUpSolid} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";

interface ActionLikeProps extends P {
    count: number;
    liked?: boolean;
}

const Root = styled.div`
    display: flex;
    flex-direction: row;

    &:hover {
        cursor: pointer;
    }
`;

const Count = styled.span`
    font-size: 80%;
    margin-left: 3px;
`;

const ActionLike = (p: ActionLikeProps) => {
    return (
        <Root>
            <FontAwesomeIcon icon={p.liked === true ? faThumbsUpSolid : faThumbsUp} />
            <Count>{p.count > 0 ? p.count : ""}</Count>
        </Root>
    );
};

export default ActionLike;

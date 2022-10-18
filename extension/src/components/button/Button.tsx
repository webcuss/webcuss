import React from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";

interface ButtonProps extends P {
    text: string;
    onClick?: () => void;
    enabled?: boolean;
}

const Button = (p: ButtonProps) => {
    const clickHandler = () => {
        if (p.onClick) {
            p.onClick();
        }
    };

    return (
        <Element type="button" value={p.text} onClick={clickHandler} disabled={p.enabled === false} />
    );
};

export default Button;

const Element = styled.input``;

import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";

interface InputProps extends P {
    type: "text" | "password" | "number";
    value?: string;
    onChange?: (v: string) => void;
}

const Input = (p: InputProps) => {
    const {value} = p;

    const changeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (p.onChange) {
            p.onChange(e.currentTarget.value)
        }
    };

    return (
        <Root type={p.type} value={value} onChange={changeHandler} />
    );
};

export default Input;

const Root = styled.input``;

import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";
import { randomNumber } from "../../utils/random";

interface CheckboxProps extends P {
    checked?: boolean;
    label?: string;
    onChange?: (v: boolean) => void;
}

const Checkbox = (p: CheckboxProps) => {
    const {checked: pChecked} = p;
    const [checked, setChecked] = useState<boolean>(pChecked === true);
    const [id] = useState<string>("id" + randomNumber(0, 9999));

    useEffect(() => {
        setChecked(pChecked === true);
    }, [pChecked]);

    const changeHandler = () => {
        setChecked(prev => {
            p.onChange && p.onChange(!prev);
            return !prev;
        });
    };

    return (
        <div className="flex flex-row ai-center">
            <Element id={id} type="checkbox" onChange={changeHandler} checked={checked} />
            {p.label && (
                <label htmlFor={id}>{p.label}</label>
            )}
        </div>
    );
};

export default Checkbox;

const Element = styled.input`
    margin-left: 0;
`;


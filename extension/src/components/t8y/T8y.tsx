import React from "react";
import { P } from "../../interfaces/common";

interface T8yProps extends P {
    text: string;
}

const T8y = (p: T8yProps) => {
    return (
        <div className={[p.className].join(" ")}>
            {p.text}
        </div>
    );
};

export default T8y;

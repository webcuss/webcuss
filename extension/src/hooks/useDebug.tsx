import React, { useEffect, useState } from "react";
import { P } from "../interfaces/common";

interface IDebug {
    isDebugging: boolean;
}

const DebugContext = React.createContext<IDebug>({
    isDebugging: false,
});

export const DebugProvider = (p: P) => {
    const [isDebugging, setIsDebugging] = useState<boolean>(false);

    const cmdListener = (cmd: string) => {
        if (cmd === 'toggle-feature-debug') {
            setIsDebugging(pv => !pv);
        }
    };

    useEffect(() => {
        chrome.commands.onCommand.addListener(cmdListener);
        return () => {
            chrome.commands.onCommand.removeListener(cmdListener);
        }
    }, [cmdListener]);

    return (
        <DebugContext.Provider value={{
            isDebugging: isDebugging,
        }}>
            {p.children}
        </DebugContext.Provider>
    );
};

export const useDebug = () => {
    return React.useContext(DebugContext);
};

import React, { useContext, useEffect, useState } from "react";
import { P } from "../interfaces/common";
import { Buffer } from "buffer";
import { useBrowserExtension } from "./useBrowserExtension";
import { b } from "../utils/bool";

interface IAuth {
    isSignedIn: boolean;
    uname?: string;
}

const AuthContext = React.createContext<IAuth>({
    isSignedIn: true,
});

export const AuthProvider = (p: P) => {
    const {chromeExt} = useBrowserExtension();

    const [token, setToken] = useState<string|undefined>(undefined);
    const [uname, setUname] = useState<string|undefined>(undefined);

    useEffect(() => {
        (async () => {
            const token = await chromeExt.storageGetValue("token");
            setToken(token);
        })();

        chrome.storage.sync.onChanged.addListener(({token}) => {
            if (token) {
                setToken(token.newValue);
            }
        });
    }, []);

    useEffect(() => {
        if (token) {
            const jwtParts = token.split(".");
            if (jwtParts.length === 3) {
                const jwtBodyBase64 = jwtParts[1];
                const jwtBody = Buffer.from(jwtBodyBase64, 'base64').toString('utf-8');
                const bodyObj = JSON.parse(jwtBody);
                setUname(bodyObj["uname"]);
                return;
            }
        }
        setUname(undefined);
    }, [token]);

    return (
        <AuthContext.Provider value={{
            isSignedIn: b(token),
            uname: uname,
        }}>
            {p.children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    return useContext(AuthContext);
};
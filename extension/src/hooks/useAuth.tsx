import React, { useContext, useEffect, useState } from "react";
import { P } from "../interfaces/common";
import { storageGetValue } from "../utils/storage";
import { Buffer } from "buffer";

interface IAuth {
    isSignedIn: boolean;
    uname?: string;
}

const AuthContext = React.createContext<IAuth>({
    isSignedIn: true,
});

export const AuthProvider = (p: P) => {
    const [isSignedIn, setIsSignedIn] = useState<boolean>(true);
    const [uname, setUname] = useState<string|undefined>(undefined);

    useEffect(() => {
        (async () => {
            const token = await storageGetValue("token");
            setIsSignedIn(!!token);
            if (token) {
                const jwtParts = token.split(".");
                if (jwtParts.length === 3) {
                    const jwtBodyBase64 = jwtParts[1];
                    const jwtBody = Buffer.from(jwtBodyBase64, 'base64').toString('utf-8');
                    const bodyObj = JSON.parse(jwtBody);
                    setUname(bodyObj["uname"]);
                }
            }
        })();

        chrome.storage.sync.onChanged.addListener(({token}) => {
            if (token) {
                setIsSignedIn(!!token.newValue);
            }
        });
    }, []);

    return (
        <AuthContext.Provider value={{
            isSignedIn: isSignedIn,
            uname: uname,
        }}>
            {p.children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    return useContext(AuthContext);
};
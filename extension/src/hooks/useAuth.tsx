import React, { useContext, useEffect, useState } from "react";
import { P } from "../interfaces/common";
import { storageGetValue } from "../utils/storage";

interface IAuth {
    isSignedIn: boolean;
}

const AuthContext = React.createContext<IAuth>({
    isSignedIn: true,
});

export const AuthProvider = (p: P) => {
    const [isSignedIn, setIsSignedIn] = useState<boolean>(true);

    useEffect(() => {
        (async () => {
            const token = await storageGetValue("token");
            setIsSignedIn(!!token);
        })();
    }, []);

    return (
        <AuthContext.Provider value={{
            isSignedIn: isSignedIn,
        }}>
            {p.children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    return useContext(AuthContext);
};
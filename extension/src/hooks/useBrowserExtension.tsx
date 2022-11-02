import React from "react";
import { BrowserExtension } from "../data-structure/browser-extension";
import { ChromeExtension } from "../data-structure/chrome-extension";
import { P } from "../interfaces/common";

interface IBrowserExtension {
    chromeExt: BrowserExtension;
}

const defaultChromeExt = new ChromeExtension();

const BrowserExtensionContext = React.createContext<IBrowserExtension>({
    chromeExt: defaultChromeExt,
});

export const BrowserExtensionProvider = (p: P) => {
    return (
        <BrowserExtensionContext.Provider value={{
            chromeExt: defaultChromeExt,
        }}>
            {p.children}
        </BrowserExtensionContext.Provider>
    );
};

export const useBrowserExtension = () => {
    return React.useContext(BrowserExtensionContext);
};

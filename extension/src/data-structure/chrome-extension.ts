import { BrowserExtension } from "./browser-extension";

export class ChromeExtension extends BrowserExtension {
    constructor() {
        super();
    }

    public getPageTitle(): Promise<string | undefined> {
        return new Promise((resolve) => {
            chrome.tabs.query({active: true, lastFocusedWindow: true}, tabs => {
                if (tabs.length < 1) {
                    console.log("tabs are empty");
                    resolve(undefined);
                    return;
                }
                const currentTab = tabs[0];
                if (!currentTab) {
                    console.log("currentTab is null");
                    resolve(undefined);
                    return;
                }
                resolve(currentTab.title);
            });
        });
    }

    public getPageUrl(): Promise<string | undefined> {
        return new Promise((resolve) => {
            chrome.tabs.query({active: true, lastFocusedWindow: true}, tabs => {
                if (tabs.length < 1) {
                    console.log("tabs are empty");
                    resolve(undefined);
                    return;
                }
                const currentTab = tabs[0];
                if (!currentTab) {
                    console.log("currentTab is null");
                    resolve(undefined);
                    return;
                }
                resolve(currentTab.url);
            });
        });
    }

    public storageSetValue(value: { [key: string]: string; }): Promise<void> {
        return new Promise<void>(res => {
            chrome.storage.sync.set(value, () => {
                res();
            });
        });
    }

    public storageGetValue(key: string): Promise<string | undefined> {
        return new Promise<string|undefined>(res => {
            chrome.storage.sync.get([key], (result) => {
                res(result[key]);
            });
        });
    }

    public storageRemoveValue(key: string | string[]): Promise<void> {
        return chrome.storage.sync.remove(key);
    }

    public storageClear(): Promise<void> {
        return chrome.storage.sync.clear();
    }
}

import { BrowserExtension } from "./browser-extension";

export class ChromeExtension extends BrowserExtension {
    constructor() {
        super();
    }

    public getPageTitle(): Promise<string | undefined> {
        return new Promise((resolve, reject) => {
            chrome.tabs.query({active: true, lastFocusedWindow: true}, tabs => {
                if (tabs.length < 1) {
                    reject("tabs are empty");
                    return;
                }
                const currentTab = tabs[0];
                if (!currentTab) {
                    reject("currentTab is null");
                    return;
                }
                resolve(currentTab.title);
            });
        });
    }

    public getPageUrl(): Promise<string | undefined> {
        return new Promise((resolve, reject) => {
            chrome.tabs.query({active: true, lastFocusedWindow: true}, tabs => {
                if (tabs.length < 1) {
                    reject("tabs are empty");
                    return;
                }
                const currentTab = tabs[0];
                if (!currentTab) {
                    reject("currentTab is null");
                    return;
                }
                resolve(currentTab.url);
            });
        });
    }
}

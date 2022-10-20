
export const storageSetValue = async (value: {[key:string]: string}): Promise<void> => {
    return new Promise<void>(res => {
        chrome.storage.sync.set(value, () => {
            res();
        });
    });
};

export const storageGetValue = async (key: string): Promise<string|undefined> => {
    return new Promise<string|undefined>(res => {
        chrome.storage.sync.get([key], (result) => {
            res(result[key]);
        });
    });
};

export const storageRemoveValue = (key: string|string[]): Promise<void> => {
    return chrome.storage.sync.remove(key);
};

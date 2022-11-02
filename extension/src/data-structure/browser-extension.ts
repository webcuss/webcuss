export abstract class BrowserExtension {
    constructor() {}

    public abstract getPageTitle(): Promise<string|undefined>;
    public abstract getPageUrl(): Promise<string|undefined>;
    public abstract storageSetValue(value: {[key:string]: string}): Promise<void>;
    public abstract storageGetValue(key: string): Promise<string|undefined>;
    public abstract storageRemoveValue(key: string|string[]): Promise<void>;
}

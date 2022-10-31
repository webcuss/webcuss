export abstract class BrowserExtension {
    constructor() {}

    public abstract getPageTitle(): Promise<string|undefined>;
    public abstract getPageUrl(): Promise<string|undefined>;
}

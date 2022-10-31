import React, { useEffect, useState } from "react";
import { useBrowserExtension } from "../../hooks/useBrowserExtension";

const PageInfo = () => {
    const {chromeExt} = useBrowserExtension();

    const [url, setUrl] = useState<string>("");
    const [title, setTitle] = useState<string>("");
    const [scheme, setScheme] = useState<string>("");
    const [hostname, setHostname] = useState<string>("");
    const [path, setPath] = useState<string>("");
    const [query, setQuery] = useState<string>("");

    useEffect(() => {
        (async () => {
            try {
                const tabUrl = await chromeExt.getPageUrl();
                if (tabUrl) {
                    setUrl(tabUrl);
                    const u = new URL(tabUrl);
                    setScheme(u.protocol);
                    setHostname(u.hostname);
                    setPath(u.pathname);
                    setQuery(u.search);
                }
            } catch (e) {
                console.log(e);
            }
            
            try {
                const tabTitle = await chromeExt.getPageTitle();
                if (tabTitle) {
                    setTitle(tabTitle);
                }
            } catch (e) {
                console.log(e);
            }
        })();
    }, []);

    return (
        <div>
            <div>URL: {url}</div>
            <div>Title: {title}</div>
            <div>Scheme: {scheme}</div>
            <div>Hostname: {hostname}</div>
            <div>Path: {path}</div>
            <div>Query: {query}</div>
        </div>
    );
};

export default PageInfo;

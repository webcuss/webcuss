import React, { useEffect, useState } from "react";

const Main = () => {
    const [url, setUrl] = useState<string>("");
    const [title, setTitle] = useState<string>("");
    const [scheme, setScheme] = useState<string>("");
    const [hostname, setHostname] = useState<string>("");
    const [path, setPath] = useState<string>("");
    const [query, setQuery] = useState<string>("");

    useEffect(() => {
        chrome.tabs.query({active: true, lastFocusedWindow: true}, tabs => {
            const currentTab = tabs[0];
            const tabUrl = currentTab.url;
            if (tabUrl) {
                setUrl(tabUrl);
                const u = new URL(tabUrl);
                setScheme(u.protocol);
                setHostname(u.hostname);
                setPath(u.pathname);
                setQuery(u.search);
            }
            if (currentTab.title) {
                setTitle(currentTab.title);
            }
        });
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

export default Main;

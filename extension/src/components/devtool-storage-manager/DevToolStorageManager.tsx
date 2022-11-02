import React, { useEffect, useState } from "react";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styled from "styled-components";
import T8y from "../t8y/T8y";
import { useBrowserExtension } from "../../hooks/useBrowserExtension";

interface IEntry {
    entryKey: string;
    entryValue: string;
}

const DevToolStorageManager = () => {
    const {chromeExt} = useBrowserExtension();

    const [entries, setEntries] = useState<IEntry[]>([]);

    useEffect(() => {
        refreshStorage();

        chrome.storage.sync.onChanged.addListener(() => {
            refreshStorage();
        });
    }, []);

    const refreshStorage = () => {
        chrome.storage.sync.get(null, (result) => {
            const res: IEntry[] = [];
            for (const key in result) {
                res.push({
                    entryKey: key,
                    entryValue: result[key]
                });
            }
            setEntries(res);
        });
    };
    
    const deleteHandler = async (key: string) => {
        await chromeExt.storageRemoveValue(key);
        refreshStorage();
    };

    return (
        <Root>
            <T8y text="chrome.storage.sync" />
            <table>
                <tbody>
                    {entries.map((v, i) => (
                        <tr key={i}>
                            <td>{v.entryKey}</td>
                            <td>{v.entryValue}</td>
                            <td>
                                <FontAwesomeIcon className="btn-delete" icon={faTrash} onClick={() => deleteHandler(v.entryKey)} />
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </Root>
    );
};

export default DevToolStorageManager;

const Root = styled.div`
    & table {
        table-layout: fixed;
        width: 100%;
    }

    & table td {
        word-wrap: break-word;         /* All browsers since IE 5.5+ */
        overflow-wrap: break-word;     /* Renamed property in CSS3 draft spec */
    }

    & .btn-delete:hover {
        cursor: pointer;
    }
`;

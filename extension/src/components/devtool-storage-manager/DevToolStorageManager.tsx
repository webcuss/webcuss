import React, { useEffect, useState } from "react";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styled from "styled-components";
import { storageRemoveValue } from "../../utils/storage";
import T8y from "../t8y/T8y";
import { randomString } from "../../utils/random";

interface IEntry {
    entryKey: string;
    entryValue: string;
}

const DevToolStorageManager = () => {
    const [entries, setEntries] = useState<IEntry[]>([]);

    useEffect(() => {
        refreshStorage();
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
        await storageRemoveValue(key);
        refreshStorage();
    };

    return (
        <Root>
            <T8y text="chrome.storage.sync:" />
            <table>
                <tbody>
                    {entries.map((v, i) => (
                        <tr key={i}>
                            <td>{v.entryKey}{": "}</td>
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
    & .btn-delete:hover {
        cursor: pointer;
    }
`;

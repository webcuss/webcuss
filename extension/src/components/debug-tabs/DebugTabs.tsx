import React, { useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";

interface IDebugTabsProps extends P {
    onActiveTabChanged?: (_: number) => void;
}

const DebugTabs = (p: IDebugTabsProps) => {
    const [activeTab, setActiveTab] = useState<number>(0);

    const tabClickHandler = (tabIndex: number) => {
        setActiveTab(tabIndex);
        if (p.onActiveTabChanged) {
            p.onActiveTabChanged(tabIndex);
        }
    };

    return (
        <Root className="flex flex-row w-full">
            <div
                className={["tab flex-1", activeTab === 0 ? "active" : ""].join(" ")}
                onClick={() => tabClickHandler(0)}>
                main
            </div>
            <div
                className={["tab flex-1", activeTab === 1 ? "active" : ""].join(" ")}
                onClick={() => tabClickHandler(1)}>
                info
            </div>
            <div
                className={["tab flex-1", activeTab === 2 ? "active" : ""].join(" ")}
                onClick={() => tabClickHandler(2)}>
                storage
            </div>
        </Root>
    );
};

export default DebugTabs;

const Root = styled.div`
    & .tab {
        text-align: center;
        padding: 2px 0;
    }

    & .tab.active {
        border-bottom: 1px solid #D5D8DC;
    }

    & .tab:hover {
        cursor: pointer;
        background-color: #D5D8DC;
    }
`;

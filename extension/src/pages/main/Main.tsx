import React, { useState } from "react";
import styled from "styled-components";
import C5t from "../../components/c5t/C5t";
import PageInfo from "../../components/page-info/PageInfo";
import SignupSuggestion from "../../components/signup-suggestion/SignupSuggestion";
import T8y from "../../components/t8y/T8y";
import { IC5t } from "../../interfaces/model";

const data: IC5t[] = [
    {
        comment: "this is great! I like birds",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "dan"
        }
    },
    {
        comment: "cute birds",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "john"
        }
    },
    {
        comment: "the second bird is adorable",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "benten"
        },
    },
    {
        comment: "the crow is thug life haha",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "jim"
        },
    },
    {
        comment: "the crow 😂",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "nathan"
        }
    },
    {
        comment: "eagleeeee",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "shim"
        }
    },
    {
        comment: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi deleniti sint fuga, voluptatibus impedit qui ipsam id magnam esse debitis nobis maxime iure quo consequatur assumenda minima hic at voluptatem?",
        createdOn: "2021-09-29T13:08:26.000Z",
        user: {
            uname: "felix"
        }
    },
];

const Main = () => {
    const [isSignedIn, setIsSignedIn] = useState<boolean>(false);
    const [comments] = useState<IC5t[]>(data);

    return (
        <Root>
            <T8y text={comments.length + " comments"} />
            {isSignedIn && (
                <>
                    {comments.map((c, i) => <C5t key={i} data={c} />)}
        
                    {"--------------------------------"}
                    <PageInfo />
                </>
            )}
            {!isSignedIn && (
                <div className="mt-20">
                    <T8y text="Complete your profile to add comments" />
                    <SignupSuggestion className="mt-10" />
                </div>
            )}
        </Root>
    );
};

export default Main;

const Root = styled.div`
    height: calc(var(--html-height) - var(--body-padding-top) - var(--body-padding-bottom));
    overflow-y: auto;
`;

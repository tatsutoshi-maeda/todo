import { useEffect } from "react";
import { handleClick } from "./service/Book";
import styled from 'styled-components'
import { TabBodyContainer } from "./components/tab-body-container";

const ListItem = styled.div`
padding: 8px 16px;

&:nth-child(n+2) {
    border-top: 1px solid #D9DBDE;
}
`

const books = [];

export const List = ({ langs }) => {
    // export const List = () => {
    
    useEffect(() => {
        console.log('List.jsのuseEffectが走った')
        // books = handleClick();
        return () => {
            console.log('List.jsのuseEffectのunmountが走った')
        }
    })
    return (
        <TabBodyContainer title = "取ってきたデータリスト（API未実装）">
            {
                langs.map((lang, index) => {
                return <ListItem key={index}>{ lang }</ListItem>
                } )
            }
            {/* {
                books.map((book, index) => {
                return <ListItem key={index}>{ book }</ListItem>
                } )
            } */}
            <div>
                <button onClick={() => handleClick()}>Get Data</button>
            </div>
        </TabBodyContainer>
    )
}
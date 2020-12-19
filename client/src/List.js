import { useEffect } from "react";
import { handleClick } from "./service/Book";
import styled from 'styled-components'

const Container = styled.div`
  padding: 12px 64px;
`

const ListItem = styled.div`
padding: 8px 16px;

&:nth-child(n+2) {
    border-top: 1px solid #D9DBDE;
}
`

export const List = ({ langs }) => {
    
    useEffect(() => {
        console.log('List.jsのuseEffectが走った')

        return () => {
            console.log('List.jsのuseEffectのunmountが走った')
        }
    })
    return (
        <Container>
            {
                langs.map((lang, index) => {
                return <ListItem key={index}>{ lang }</ListItem>
                } )
            }
            <div>
                <button onClick={() => handleClick()}>Get Data</button>
            </div>
        </Container>
    )
}
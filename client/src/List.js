import { useEffect } from "react";
import { handleClick } from "./service/Book";

export const List = ({ langs }) => {
    
    useEffect(() => {

    })
    return (
        <div>
            {
                langs.map((lang, index) => {
                return <div key={index}>{ lang }</div>
                } )
            }
            <div>
                <button onClick={() => handleClick()}>Get Data</button>
            </div>
        </div>
    )
}
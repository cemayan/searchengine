import {Cards} from "@/components/cards.tsx";
import React from "react";

export const Results = (props: any    ) =>  {
    let [counter] = React.useState(0);

    return (
        <div className="gap-2 grid grid-cols-2 sm:grid-cols-4">
            {props.records && props.records.map((item: any, _: any) => {
                counter++
               return (
                   <Cards key={counter} title={item.title} url={item.url}/>
               )
            })}
        </div>
    )
}
import {SearchIcon} from "@/components/icons.tsx";
import {Autocomplete, AutocompleteItem} from "@nextui-org/react";
import {useAsyncList} from "@react-stately/data";


export const SearchBar = (props: any    ) =>  {


    let list: any  = useAsyncList({
        async load({signal, filterText}) {

            var finalArr = []

            if(filterText != "" && props.selectedKey != filterText ){
                let res = await fetch(`${import.meta.env.VITE_READAPI}/v1/query?q=${filterText}`, {signal});
                let json: string[] = await res.json();

                for (let item of json) {
                    var obj = {
                        name: item
                    }
                    finalArr.push(obj);
                }
            }


            return {
                items: finalArr,
            };
        },
    });



    let sendSelectionRequest = async (query:string, selectedKey:string) => {
        const rawResponse = await fetch(`${import.meta.env.VITE_WRITEAPI}/v1/selection`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({query: query, selectedKey: selectedKey})
        });
        return  await rawResponse.json();
    };


    let getResultsRequest = async (_selectedKey: string) => {
        const rawResponse = await fetch(`${import.meta.env.VITE_READAPI}/v1/results`, {
            method: 'GET',
            headers: {
                'Origin': 'http://localhost:5173',
                'X-SearchEngine-Query': _selectedKey,
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        });
        return  await rawResponse.json();
    };



    return (
        <Autocomplete
            autoFocus={true}
            inputValue={list.filterText}
            isLoading={list.isLoading}
            items={list.items}
            fullWidth={true}
            placeholder="Type to search..."
            variant="bordered"
            onInputChange={list.setFilterText}
            startContent={
                <SearchIcon className="text-base text-default-400 pointer-events-none flex-shrink-0"/>
            }
            onSelectionChange={(_selectedKey: any) => {

                if (_selectedKey) {
                    props.setSelectedKey(_selectedKey);
                    sendSelectionRequest(list.filterText, _selectedKey).then(() => {
                        props.setSelectedKey("");
                        getResultsRequest(_selectedKey).then((value: any) => {
                            props.setRecords(value.items);
                        });
                    });
                }
            }}
        >

            {(item: { name: any; }) => (
                <AutocompleteItem key={item.name} >
                    {item.name}
                </AutocompleteItem>
            )}
        </Autocomplete>
    )
}
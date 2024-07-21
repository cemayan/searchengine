import DefaultLayout from "@/layouts/default";
import React from "react";
import {SearchBar} from "@/components/autocomplete.tsx";
import {Results} from "@/components/results.tsx";

export default function IndexPage() {

    const [selectedKey, setSelectedKey] = React.useState('');
    const [records, setRecords] = React.useState([]);


  return (
    <DefaultLayout>
      <section className="flex flex-col  justify-center">
        <div className="mt-8">
          <SearchBar setSelectedKey={setSelectedKey}  setRecords={setRecords} selectedKey={selectedKey} ></SearchBar>
        </div>
        <br/>

        <Results records={records}></Results>

      </section>
    </DefaultLayout>
);
}

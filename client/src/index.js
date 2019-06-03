import React from "react";
import ReactDOM from "react-dom";
import FolderView from "./files/FolderView";

const Index = () => {
  return <div>
    Hello React! mod
      <FolderView path="/"></FolderView>
    </div>;
};

ReactDOM.render(<Index />, document.getElementById("index"));

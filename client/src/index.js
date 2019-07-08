import React, { Component } from 'react';
import ReactDOM from "react-dom";

import Axios from "axios";

import FolderView from "./files/FolderView";
import FolderUp from './files/FolderUp';

import "./styles/main.scss"
import "./styles/icons.scss"

class Index extends Component {

    constructor(props) {
        super(props);
        this.state = { 
            currentPath: '/',
            folder: [],
            isLoaded: false
        }

        this.updatePath = this.updatePath.bind(this);
    }

    componentDidMount() {
        this.updatePath(this.state.currentPath);
    }

    updatePath(newPath) {
        Axios.get('/api/files?path=' + newPath)
            .then(result => {
                this.setState({
                    folder: result.data,
                    isLoaded: true,
                    currentPath: result.data.fullPath
                });
            });
    }

    render() {
        return <div>
            <h1>Loaded path: <b>{this.state.currentPath}</b></h1>
            <FolderUp 
                path={this.state.currentPath}
                updatePath={this.updatePath}></FolderUp>
            <FolderView 
                path={this.state.currentPath}
                folder={this.state.folder}
                isLoaded={this.state.isLoaded}
                updatePath={this.updatePath}
                ></FolderView>
        </div>;
    }
};


ReactDOM.render(<Index />, document.getElementById("index"));

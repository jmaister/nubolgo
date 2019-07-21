import React, { Component } from 'react';
import ReactDOM from "react-dom";

import Axios from "axios";

import FolderView from "./files/FolderView";
import Uploader from './files/Uploader';
import Actions from './actions/Actions';

import "./styles/reset.scss"
import "./styles/main.scss"
import "./styles/icons.scss"

class Index extends Component {

    constructor(props) {
        super(props);
        this.state = { 
            currentPath: '/',
            folder: [],
            isLoaded: false,
            selected: {}
        }

        this.updatePath = this.updatePath.bind(this);
        this.updateSelected = this.updateSelected.bind(this);
    }

    componentDidMount() {
        this.updatePath(this.state.currentPath);
    }

    updatePath(newPath) {
        Axios.get('/api/list?path=' + newPath)
            .then(result => {
                this.setState({
                    folder: result.data,
                    isLoaded: true,
                    currentPath: result.data.fullPath,
                    selected: {}
                });
            });
    }

    updateSelected(key, value) {
        this.setState({
            selected:{...this.state.selected, [key]: value }
        });
    }

    render() {
        return <div>
            <h1>Loaded path: <b>{this.state.currentPath}</b></h1>
            <Uploader path={this.state.currentPath}></Uploader>
            <Actions
                folder={this.state.folder}
                selected={this.state.selected}
                ></Actions>
            <FolderView 
                path={this.state.currentPath}
                folder={this.state.folder}
                isLoaded={this.state.isLoaded}
                selected={this.state.selected}
                updateSelected={this.updateSelected}
                updatePath={this.updatePath}
                ></FolderView>
        </div>;
    }
};

ReactDOM.render(<Index />, document.getElementById("main"));

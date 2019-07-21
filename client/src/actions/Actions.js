import React, { Component } from 'react';
import ReactModal from 'react-modal';

import Axios from "axios";

class Actions extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isDeleteDialogOpen: false
        };

        this.delete = this.delete.bind(this);
        this.cancelDelete = this.cancelDelete.bind(this);
        this.confirmDelete = this.confirmDelete.bind(this);
        this.closeResultDelete = this.closeResultDelete.bind(this);

        ReactModal.setAppElement('#main');
    }

    delete() {
        this.setState({
            isDeleteConfirmOpen: true
        })
    }

    cancelDelete() {
        this.setState({
            isDeleteConfirmOpen: false
        });
    }

    confirmDelete() {
        const promises = Object.getOwnPropertyNames(this.props.selected).map(file => {
            return Axios.delete('/api/file?path=' + file);
        });
        Promise.all(promises).then(responses => {
            this.setState({
                isDeleteResultOpen: true
            });
            console.log("responses", responses);
        });
    }

    closeResultDelete() {
        this.setState({
            isDeleteResultOpen: false
        });
    }

    render() { 
        const fileList = <ul>
            {this.props.folder.files ? this.props.folder.files
            .filter(e => this.props.selected[e.name])
            .map(e => {
                return <li key={e.name}>{e.name}</li>;
            }): null}
            </ul>;
        return <div>
            <button onClick={this.delete}>Delete</button>
            <ReactModal isOpen={this.state.isDeleteConfirmOpen}>
                <div><h3>Do you want to delete these elements?</h3></div>
                {fileList}
                <div>
                    <button onClick={this.confirmDelete}>Yes</button>
                    <button onClick={this.cancelDelete}>No</button>
                </div>
            </ReactModal>
            <ReactModal isOpen={this.state.isDeleteResultOpen}>
                asdf
                <button onClick={this.closeResultDelete}>No</button>
            </ReactModal>
        </div>;
    }
}
 
export default Actions;
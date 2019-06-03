import React, { Component } from 'react';
import Axios from "axios";


class FolderView extends Component {

    constructor(props) {
        super(props);
        this.state = { 
            isLoaded: false
         }
    }

    componentDidMount() {
        Axios.get('/api/files?path=' + this.props.path)
            .then(result => {
                this.setState({
                    folder: result.data,
                    isLoaded: true
                });
            })
    }

    render() {
        const { isLoaded, folder } = this.state;
        if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
            return <div>Loaded path: <b>{this.props.path}</b>
                <div>{JSON.stringify(folder)}</div>
            </div>;
        }
        
    }
}
 
FolderView.defaultProps = {
    path: '/'
};


export default FolderView;
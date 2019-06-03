import React, { Component } from 'react';


class FolderView extends Component {

    constructor(props) {
        super(props);
        this.state = { 
            isLoaded: false
        }

        this.folderClick = this.folderClick.bind(this);
    }

    folderClick(folder) {
        this.props.updatePath(this.props.path + folder.name);
    }

    render() {
        console.log("state", this.state);
        const { isLoaded, folder } = this.props;
        if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
            const items = folder.files;
            console.log("items", items);
            let folders = items.filter(f => f.isFolder).map(folder => {
                return <div key={folder.name} onClick={this.folderClick.bind(this, folder)}>📁 {folder.name}</div>
            });
            let files = items.filter(f => !f.isFolder).map(file => {
                return <div key={file.name}>📄 {file.name}</div>
            });

            return <div>
                {folders}
                {files}
            </div>;
        }
        
    }
}
 
FolderView.defaultProps = {
    path: '/'
};


export default FolderView;
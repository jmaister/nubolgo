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
            return <div className="item"><span className="icon loading"></span>Loading...</div>;
        } else {
            const items = folder.files;
            console.log("items", items);
            let folders = items.filter(f => f.isFolder).map(folder => {
                return <div className="item folder" key={folder.name} onClick={this.folderClick.bind(this, folder)}><span className="icon folder"></span> {folder.name}</div>
            });
            let files = items.filter(f => !f.isFolder).map(file => {
                return <div className="item file" key={file.name}><span className="icon file"></span> {file.name}</div>
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
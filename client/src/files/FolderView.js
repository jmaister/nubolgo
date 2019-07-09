import React, { Component } from 'react';


class FolderView extends Component {

    constructor(props) {
        super(props);
        this.state = { 
            isLoaded: false
        }
    }

    folderClick(folder) {
        console.log("folder:", folder);
        this.props.updatePath(folder.fullPath);
    }

    fileClick(file) {
        console.log("file:", file);
        this.props.downloadPath(file.fullPath);
    }

    render() {
        console.log("state", this.state);
        const { isLoaded, folder } = this.props;
        if (!isLoaded) {
            return <div className="item"><span className="icon loading"></span>Loading...</div>;
        } else {
            let sum = 0;
            const dfOptions = {
                year: 'numeric', month: '2-digit', day: '2-digit',
                hour: '2-digit', minute: '2-digit', second: '2-digit',
                timeZoneName: 'short'
            };
            const df = Intl.DateTimeFormat('default', dfOptions);
            const nf = Intl.NumberFormat('default');

            const items = folder.files.map(i => {
                const iconName = i.isFolder ? 'folder' : 'file';
                let clickFn = i.isFolder 
                    ? this.folderClick.bind(this, i)
                    : this.fileClick.bind(this, i);

                sum = sum + (i.isFolder ? 0 : i.size);

                return <tr className={iconName} key={i.name}>
                    <td>
                        <input type="checkbox"></input>
                    </td>
                    <td className="clickable" onClick={clickFn}>
                        <span className={"icon " + iconName}></span> {i.name}
                    </td>
                    <td className="size">
                        {i.isFolder ? '-' : nf.format(i.size)}
                    </td>
                    <td>
                        {df.format(i.time)}
                    </td>
                </tr>;
            });

            return <table className="files-table">
                <thead>
                    <tr>
                        <th><input type="checkbox"></input></th>
                        <th>Name</th>
                        <th>Size</th>
                        <th>Updated</th>
                    </tr>
                </thead>
                <tbody>
                    {items}
                </tbody>
                <tfoot>
                    <tr>
                        <td></td>
                        <td></td>
                        <td className="size">{nf.format(sum)}</td>
                        <td></td>
                    </tr>
                </tfoot>
            </table>;
        }
        
    }
}
 
FolderView.defaultProps = {
    path: '/'
};


export default FolderView;
import React, { Component } from 'react';

class FolderView extends Component {

    constructor(props) {
        super(props);
        this.state = {
            selected: {}
        };
    }

    folderClick(folder) {
        this.props.updatePath(folder.fullPath);
    }

    folderUp() {
        let newPath = this.props.path + '/../';
        this.props.updatePath(newPath);
    }

    onChange(key, value) {
        this.props.updateSelected(key, value);
    }
    
    render() {
        const { isLoaded, folder } = this.props;
        if (!isLoaded) {
            return <div className="item"><span className="icon loading"></span>Loading...</div>;
        } else {
            const isRootFolder = folder.fullPath === '/';
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
                let link = i.isFolder 
                    ? <span className="link clickable" onClick={this.folderClick.bind(this, i)}><span className={"icon " + iconName}></span> {i.name}</span>
                    : <a className="link" href={"/api/file?path=" + i.fullPath}><span className={"icon " + iconName}></span> {i.name}</a>;

                sum = sum + (i.isFolder ? 0 : i.size);

                return <tr className={iconName + " " + (this.props.selected[i.name] ? 'selected' : '')} key={i.name}>
                    <td>
                        <input type="checkbox"
                            onChange={(e) => this.onChange(i.name, e.target.checked)}
                            value={this.props.selected[i.name]}></input>
                    </td>
                    <td className="td-link">
                        {link}
                    </td>
                    <td className="size">
                        {i.isFolder ? '-' : nf.format(i.size)}
                    </td>
                    <td>
                        {df.format(i.time)}
                    </td>
                </tr>;
            });

            const goUpRow = <tr key="up">
                <td></td>
                <td className="td-link">
                    <span className="link clickable" onClick={this.folderUp.bind(this)}><span className={"icon back"}></span></span>
                </td>
                <td className="size">
                    -
                </td>
                <td>
                    -
                </td>
            </tr>;

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
                    {isRootFolder ? null : goUpRow}
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
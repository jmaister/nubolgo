import React, { Component } from 'react';

class FolderUp extends Component {
    constructor(props) {
        super(props);
        this.state = {  }
    }

    folderUp() {
        let newPath = this.props.path + '/../';

        this.props.updatePath(newPath);
    }

    render() { 
        const isRoot = this.props.path === '/';

        return ( <div onClick={this.folderUp.bind(this)}>"⬑"</div> );
    }
}
 
export default FolderUp;
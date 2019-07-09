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
        const isRoot = (this.props.path === '/');



        return isRoot ? (<div className="item"></div>) :
         ( <div className="item clickable" onClick={this.folderUp.bind(this)}><span className="icon back"></span></div> );
    }
}
 
export default FolderUp;
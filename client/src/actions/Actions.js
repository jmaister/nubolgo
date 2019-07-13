import React, { Component } from 'react';

class Actions extends Component {
    constructor(props) {
        super(props);
        this.state = {  };

        this.delete = this.delete.bind(this);
    }

    delete() {
        let msg = "Do you want to delete these elements?\n";
        this.props.folder.files.forEach(e => {
            if (this.props.selected[e.name]) {
                msg = msg + e.name + "\n";
            }
        });
        const answer = confirm(msg);
        console.log("answer", answer);
    }

    render() { 
        return <div>
            <button onClick={this.delete}>Delete</button>
        </div>;
    }
}
 
export default Actions;
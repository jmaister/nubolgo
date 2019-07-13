import React, { Component } from 'react';

import Dropzone from 'react-dropzone'

class Uploader extends Component {
    constructor(props) {
        super(props);
        this.state = {};

        this.onDrop = this.onDrop.bind(this);
    }

    onDrop(acceptedFiles) {
        console.log(acceptedFiles);
        console.log(this.props);

        const formData = new FormData();
  
        formData.append("path", this.props.path);
        acceptedFiles.map((file, index) => {
            formData.append("files[]", file);
        });
        
        fetch('/api/upload', {
            // content-type header should not be specified!
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(success => {
            // Do something with the successful response
            console.log(success);
        })
        .catch(error => console.log(error));        
    }

    render() { 
        return <Dropzone onDrop={this.onDrop}>
            {({getRootProps, getInputProps}) => (
                <section className="dropzone">
                <div {...getRootProps()}>
                    <input {...getInputProps()} />
                    <p>Drag 'n' drop some files here, or click to select files</p>
                </div>
                </section>
            )}
        </Dropzone>;
    }
}
 
export default Uploader;
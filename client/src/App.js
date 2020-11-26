import React from 'react'
import { v4 as uuidv4 } from 'uuid';
import axios from 'axios';
import {apiDeleteFile, apiUploadFile} from './api';

export default class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      uuid: uuidv4(),
    }

    this.uploadFile = this.uploadFile.bind(this);
    this.refreshUUID = this.refreshUUID.bind(this);
    this.deleteObject = this.deleteObject.bind(this);

    this.fileInput = React.createRef();

    this.backend = axios.create({
      baseURL: 'http://localhost:9000/api',
      timeout: 1000,
      
    });

    this.minio = axios.create({
      baseURL: 'http://localhost:9000',
      timeout: 1000,
    });
  }

  refreshUUID(event) {
    event.preventDefault();
    console.log('uuid ->', this.state.uuid);
    this.setState({uuid: uuidv4()});
  }

  deleteObject(event) {
    event.preventDefault();

    this.backend.delete(apiDeleteFile(
      this.state.uuid,
      this.fileInput.current.files[0].name,
    ));
  }

  uploadFile(event) {
    event.preventDefault();

    this.backend.get(apiUploadFile(
      this.fileInput.current.files[0].size,
      this.state.uuid,
      this.fileInput.current.files[0].name,
    ))
    .then((resp) => {
      if (resp.status === 200) {
        const url = new URL(resp.data);
        console.log(url.pathname);
        this.minio.put(url.pathname, this.fileInput.current.files[0]);
      }
    });
  }

  render() {
    return (
      <div className="wrapper">
        <div className="control-panel">
          <input type="file" ref={this.fileInput}/>
          <button onClick={this.uploadFile} type="button">Upload</button>
          <button onClick={this.deleteObject} type="button">Delete</button>
          <button onClick={this.refreshUUID} type="button">New uuid</button>
        </div>
      </div>
    );
  }
}

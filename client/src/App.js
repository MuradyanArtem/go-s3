import React from 'react'
import { v4 as uuidv4 } from 'uuid';
import axios from 'axios';
import {apiUploadFile} from './api';

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
        console.log(resp.data);
        axios.put(resp.data, this.fileInput.current.files[0])
        .then((resp) => {
          console.log('s3', resp.status);
          console.log('s3', resp.data);
          }
        )
        .catch((resp) => {
          console.log('s3', resp.status);
          console.log('s3', resp.data);
          });
      }
    })
    .catch((resp) => {
      console.log('backend', resp.status);
      console.log('backend', resp.data);
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

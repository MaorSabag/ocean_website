import axios from 'axios'
import { METHOD } from '../Models/index'

const sendRequest = async (method: METHOD, routing: string, body?: any) => {
    let response: any = null;
  
    switch (method) {
      case METHOD.get:
        response = await axios.get(
          routing,
          {
            params:body,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        );
        
        break;
      case METHOD.post:
        response = await axios.post(
          routing,
          { body: body },
          {headers : {
            'Content-Type': 'applicaiton/json',
          }}
        );
        break;
    }
  
    return response.data;
  }
  


export const getRepos = () => {
    console.log("Sending getDatabase api request..")
    return sendRequest(
        METHOD.get,
        '/api/repositories'
    )
}

export const getHome = () => {
    console.log("Sending / api request..")
    return sendRequest(
        METHOD.get,
        '/api/'
    )
}
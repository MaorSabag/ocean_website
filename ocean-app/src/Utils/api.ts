import axios from 'axios'
import { METHOD } from '../Models/index'

const BACKEND_SERVER = !process.env.NODE_ENV || process.env.NODE_ENV === 'development' ? process.env.REACT_APP_BACKEND_DEV : process.env.REACT_APP_BACKEND_PROD

const sendRequest = async (method: METHOD, routing: string, body?: any) => {
    const URL = `${BACKEND_SERVER}${routing}`;
    let response: any = null;
  
    switch (method) {
      case METHOD.get:
        response = await axios.get(
          URL,
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
          URL,
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
        '/repositories'
    )
}

export const getHome = () => {
    console.log("Sending / api request..")
    return sendRequest(
        METHOD.get,
        '/'
    )
}
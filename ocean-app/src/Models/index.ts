
export enum METHOD {
    get= 'GET',
    post= 'POST'
}

export type Repository = {
    Name: String,
    Language: String,
    Stars: Number,
    Description: String,
    Link: String
}

export type errorMessage = {
    Error: String,
    Status: String
}

export type Database = Array<Repository>

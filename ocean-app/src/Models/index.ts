
export enum METHOD {
    get= 'GET',
    post= 'POST'
}

export type Repository = {
    Name: String,
    Language: String,
    Stars: Number,
    Description: String,
    Link: String,
    ReleaseDate: Date
}

export type Repositories = Array<Repository>

export type errorMessage = {
    Error: String,
    Status: String
}

import { useState, useEffect } from 'react';
import {Database, errorMessage} from '../../Models/index';
import { getRepos } from '../../Utils/api';
import { RepoCard } from './RepoCard';
import { Background } from './Background';
import { AlertPopUp } from './AlertPopUp';
import { Alert, CircularProgress } from '@mui/material';



export const Repos = () => {
    const [database, setDatabase] = useState<Database>([]);
    const [filteredDatabase, setFilteredDatabase] = useState<Database>([]);
    const [error, setError] = useState<errorMessage>();
    const [isOpenAlert, setIsOpenAlert] = useState(false);
    const [searchText, setSearchText] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleError = (alertText: errorMessage) => {
        setError(alertText);
        setIsOpenAlert(true);
    }
  
    const updateDatabase = async () => {
      try {
        setIsLoading(true)  
        const data = await getRepos();
          console.log(data)
          if (Array.isArray(data)) {
            setDatabase(data)
            setFilteredDatabase(data)
            setIsLoading(false)
            return
          }
          setIsLoading(false)
          handleError(data)

        } catch(error) {
            setIsLoading(false)
            const alertText = {
                Error: "Error fetching the repositories",
                Status: "No Ok"
            }
            handleError(alertText)
        }
    }
  
    useEffect(() => {
      const fetchData = async () => {
        await updateDatabase();
      };
  
      fetchData();
    }, []);
  
    const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        const searchText = e.target.value
        setSearchText(searchText)
  
        const filteredData = database.filter(repo => {
          const repoName = repo.Name.toString().toLowerCase() + repo.Description.toString().toLowerCase() + repo.Language.toString().toLowerCase()
          return repoName.includes(searchText.toLowerCase())
        })
  
        setFilteredDatabase(filteredData)
    }
  
    return (
      <div>
        <div>
          <Background
            handleSearch={handleSearch}
          />
        </div>
        <div className="repoContainer">
          {isLoading && 
            <div className="centered-container">
                <CircularProgress 
                        className="processBar"
                        size={'10rem'}
                />
            </div>
            }
          { !isLoading && isOpenAlert ? (
              console.log(`Got In isOpenAlert ${error}`),
              <div>
                  <AlertPopUp
                        errorMessage={error}
                        isOpenAlert={isOpenAlert}
                        onClose={async () => {
                            setIsOpenAlert(false)
                            await updateDatabase()
                        }}
                    />
              </div>
            ) : (
              filteredDatabase.map((repo) => (
                <div>
                  <RepoCard
                      key={repo.Name.toString()}
                      Description={repo.Description}
                      Language={repo.Language}
                      Link={repo.Link}
                      Name={repo.Name}
                      NumOfLikes={repo.NumOfLikes}
                      Stars={repo.Stars}
                    />
                  <br/>
                </div>
                
              ))
            )
          }
        </div>
      </div>
    );
  }
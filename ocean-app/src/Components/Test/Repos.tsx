import { useState, useEffect } from 'react';
import {Repositories, errorMessage} from '../../Models/index';
import { getRepos } from '../../Utils/api';
import { RepoCard } from './RepoCard';
import { Background } from './Background';
import { SortButton } from './sortButton';
import { AlertPopUp } from './AlertPopUp';
import {  CircularProgress } from '@mui/material';



export const Repos = () => {
    const [repositories, setRepositories] = useState<Repositories>([]);
    const [filteredRepositories, setFilteredRepositories] = useState<Repositories>([]);
    const [error, setError] = useState<errorMessage>();
    const [isOpenAlert, setIsOpenAlert] = useState(false);
    const [searchText, setSearchText] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [isFilteredByDate, setIsFilteredByDate] = useState(false);
    const [isFilteredByStars, setIsFilteredByStars] = useState(false);

    const handleError = (alertText: errorMessage) => {
      setError(alertText);
      setIsOpenAlert(true);
    }

    const filterByDate = () => {
      if(isFilteredByDate) {
        const sortedRepos = [...repositories].sort((a, b) => 
          new Date(b.ReleaseDate).valueOf() - new Date(a.ReleaseDate).valueOf()
        )
        console.log(sortedRepos)
        setRepositories(sortedRepos)
        setFilteredRepositories(sortedRepos)
        setIsFilteredByDate(false)
      } else {
        const sortedRepos = [...repositories].sort((a, b) => 
          new Date(b.ReleaseDate).valueOf() - new Date(a.ReleaseDate).valueOf()
        ).reverse()
        console.log(sortedRepos)
        setRepositories(sortedRepos)
        setFilteredRepositories(sortedRepos)
        setIsFilteredByDate(true)
      }
    }
    
    const filterByStars = () => {
      if(isFilteredByStars) {
        const sortedRepos = [...repositories].sort((a, b) => 
          b.Stars.valueOf() - a.Stars.valueOf() 
        )
        setRepositories(sortedRepos)
        setFilteredRepositories(sortedRepos)
        setIsFilteredByStars(false)
      } else {
        const sortedRepos = [...repositories].sort((a, b) => 
          b.Stars.valueOf() - a.Stars.valueOf() 
        ).reverse()
        setRepositories(sortedRepos)
        setFilteredRepositories(sortedRepos)
        setIsFilteredByStars(true)
      }
    }
  
    const updateDatabase = async () => {
      try {
        setIsLoading(true)  
        const data = await getRepos();
        console.log("Got in updateDateabase ",data)
        if (Array.isArray(data)) {
          setRepositories(data)
          setFilteredRepositories(data)
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

        if (searchText.includes("<script>alert(1)</script>")) {
            alert(1); // for the lols
        }
  
        const filteredData = repositories.filter(repo => {
          const repoName = repo.Name.toString().toLowerCase() + repo.Description.toString().toLowerCase() + repo.Language.toString().toLowerCase()
          return repoName.includes(searchText.toLowerCase())
        })
  
        const sortedRepos = [...filteredData].sort((a, b) => 
          b.Stars.valueOf() - a.Stars.valueOf() 
        )
        setFilteredRepositories(sortedRepos)
    }
  
    return (
      <div>
        <div>
            <Background
              handleSearch={handleSearch}
            />
            <SortButton
              sortByDate={filterByDate}
              sortByStars={filterByStars}
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
              
              filteredRepositories.map((repo) => (
                <div>
                  <RepoCard
                      key={repo.Name.toString()}
                      Description={repo.Description}
                      Language={repo.Language}
                      Link={repo.Link}
                      Name={repo.Name}
                      Stars={repo.Stars}
                      ReleaseDate={repo.ReleaseDate}
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
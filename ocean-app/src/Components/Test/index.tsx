// import React from 'react';
import { useState, useEffect } from 'react';
import './Test.css';
import backgroundImg from '../../Images/background1.png'
import {Database, errorMessage, Repository} from '../../Models/index'
import { getRepos } from '../../Utils/api'
import AspectRatio from '@mui/joy/AspectRatio';
import Card from '@mui/joy/Card';
import CardOverflow from '@mui/joy/CardOverflow';
import Divider from '@mui/joy/Divider';
import Typography from '@mui/joy/Typography';
import Link from '@mui/joy/Link';

const RepoCard = (props: Repository) => {

    return (
      <div className="repoData">
        <Card variant="outlined" sx={{ width: 320, height: 290, maxHeight: '100%' }}>
          <CardOverflow>
            <AspectRatio ratio="2">
              <img
                src={backgroundImg}
                loading="lazy"
                alt=""
              />
            </AspectRatio>
          </CardOverflow>

          <CardOverflow sx={{display: 'flex', overflow: 'auto', overflowY: 'auto'}}>
            <CardOverflow>
              <Typography level="h2" sx={{ fontSize: 'md', mt: 2 }}>
                <Link href={props.Link.toString()} overlay underline="none">
                  {props.Name}
                </Link>
              </Typography>
              <Typography level="body2" sx={{ mt: 0.5, mb: 2 }}>
                {props.Language}
              </Typography>
            </CardOverflow>
            <Typography level="h2" sx={{ fontSize: 'md', mt: 2}}>
                {props.Description}
            </Typography>
          </CardOverflow>

          <Divider inset="context" />

          <CardOverflow
          variant="soft"
          sx={{
            display: 'flex',
            gap: 1.5,
            py: 1.5,
            px: 'var(--Card-padding)',
            bgcolor: 'rgba(255, 255, 255, 0.8)',
          }}
        >
          <Typography level="body3" sx={{ fontWeight: 'md', color: 'text.secondary' }}>
            Numer of ⭐'s {props.Stars.toString()}
          </Typography>
          <Divider orientation="vertical" />
          <Typography level="body3" sx={{ fontWeight: 'md', color: 'text.secondary' }}>
            Num of Likes {props.NumOfLikes.toString()}
          </Typography>
        </CardOverflow>
        </Card>
      </div>
      
    );
}


const Repos = () => {
  const [database, setDatabase] = useState<Database>([]);
  const [filteredDatabase, setFilteredDatabase] = useState<Database>([]);
  const [error, setError] = useState<errorMessage>();
  const [searchText, setSearchText] = useState('');

  const updateDatabase = async () => {
    try {
        const data = await getRepos();
        console.log(data)
        if (Array.isArray(data)) {
          setDatabase(data)
          setFilteredDatabase(data)
          return
        }
        setError(data)
      } catch(error) {
        console.log(error)    
        setError({Error: "Error fetching the database", Status: "No Ok"})
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
        {error ? (
            console.log(error),
            <div>
              <p>Error: {error.Error}</p>
              <p>Error: {error.Status}</p>
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

const Background = ( { handleSearch }: any) => {  
  return (
    <div className="background">
      <div className="searchBar">
        <input type="text" placeholder="Search in repos" onChange={handleSearch}></input>
      </div>
      <div>
        <div id="social" className="social">
          <ul className="footer-social-links list-reset">
            <li>
              <a href="https://www.linkedin.com/in/maor-sabag/">
                <span className="screen-reader-text">LinkedIn</span>
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M19 0h-14c-2.761 0-5 2.239-5 5v14c0 2.761 2.239 5 5 5h14c2.762 0 5-2.239 5-5v-14c0-2.761-2.238-5-5-5zm-11 19h-3v-11h3v11zm-1.5-12.268c-.966 0-1.75-.79-1.75-1.764s.784-1.764 1.75-1.764 1.75.79 1.75 1.764-.783 1.764-1.75 1.764zm13.5 12.268h-3v-5.604c0-3.368-4-3.113-4 0v5.604h-3v-11h3v1.765c1.396-2.586 7-2.777 7 2.476v6.759z"/></svg>
              </a>
            </li>
            <li>
              <a href="https://github.com/MaorSabag">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div className="innerDiv">
        <img src={backgroundImg} alt="Background" />
      </div>
    </div>
  );
};

const AllRepos = () => {
  return (
    <div className="pageContainer">
      <div>
        <Repos/>
      </div>
        
    </div>
  );
};

export default AllRepos;

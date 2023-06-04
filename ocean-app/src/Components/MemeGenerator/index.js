import React, { useEffect, useState } from 'react';
import axios from 'axios';

const MemeGenerator = () => {
    const [memeImg, setMemeImg] = useState('');

    useEffect(() => {
        fetchRandomMeme();
    }, []);

    const fetchRandomMeme = async () => {
        try {
            const responses = await axios.get("https://api.imgflip.com/get_memes");
            const memeImages = responses.data.data.memes;
            const memeIndex = Math.floor(Math.random() * memeImages.length);
            setMemeImg(memeImages[memeIndex].url);

        } catch (error) {
            console.log('Erroring fetching meme: ', error);
        }
    }
    
    return (
        <div>
            <h1>Random Meme Generator</h1>
            <img
                id="memeImg"
                src={memeImg}
                alt="Meme"
                style={{
                    maxWidth: '800px',
                    maxHeight: '800px',
                    marginTop: '20px',
                    boxShadow: '0 2px 5px rgba(0, 0, 0, 0.3)',
                    borderRadius: '5px',
                }}
            />
        </div>
    )
}

export default MemeGenerator
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MemeGenerator from './Components/MemeGenerator';
import Home from './Components/Home';
import AllRepos from './Components/AllRepos';
import NotFoundPage from './Components/NotFound';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<AllRepos />} />
        <Route path="/meme" element={<MemeGenerator />} />
        <Route path="/test" element={<Home />} />
        <Route path="/NotFound" element={<NotFoundPage />} />
      </Routes>
    </Router>
  );
}

export default App;

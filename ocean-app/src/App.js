import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MemeGenerator from './Components/MemeGenerator';
import Home from './Components/Home';
import AllRepos from './Components/Test';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/meme" element={<MemeGenerator />} />
        <Route path="/" element={<AllRepos />} />
        <Route path="/test" element={<Home />} />
        <Route path="*" element={<Home />} />
      </Routes>
    </Router>
  );
}

export default App;

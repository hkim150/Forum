import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Main from './main/main'
import PostContent from './postContent/postContent'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/posts/:id" element={<PostContent />} />
      </Routes>
    </Router>
  )
}

export default App

import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Main from './main/main'
import PostContentDetail from './postContentDetail/postContentDetail'
import CreatePost from './createPost/createPost'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/posts/:id" element={<PostContentDetail />} />
        <Route path="/create-post" element={<CreatePost />} />
      </Routes>
    </Router>
  )
}

export default App

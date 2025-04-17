import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Main from './main/main'
import PostContentDetail from './postContentDetail/postContentDetail'
import CreatePost from './createPost/createPost'
import EditPost from './editPost/editPost'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/posts/:id" element={<PostContentDetail />} />
        <Route path="/create-post" element={<CreatePost />} />
        <Route path="/edit-post/:id" element={<EditPost />} />
      </Routes>
    </Router>
  )
}

export default App

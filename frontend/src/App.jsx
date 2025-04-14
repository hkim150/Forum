import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Table from './components/Table'
import PostContent from './postContent/postContent'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Table />} />
        <Route path="/posts/:id" element={<PostContent />} />
      </Routes>
    </Router>
  )
}

export default App

import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import MainLayout from './components/layout/MainLayout/MainLayout'
import Login from './pages/Login/Login'
import Register from './pages/Register/Register'
import WorksList from './pages/WorksList/WorksList'
import WorkEditor from './pages/WorkEditor/WorkEditor'
import Optimization from './pages/Optimization/Optimization'
import Publish from './pages/Publish/Publish'

function App() {
  const isAuthenticated = !!localStorage.getItem('token')

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={isAuthenticated ? <MainLayout /> : <Navigate to="/login" />}>
          <Route index element={<Navigate to="/works" />} />
          <Route path="works" element={<WorksList />} />
          <Route path="works/:workId/*" element={<WorkEditor />} />
          <Route path="optimization" element={<Optimization />} />
          <Route path="publish" element={<Publish />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App

import { Outlet } from 'react-router-dom'
import Header from '../Header/Header'
import Sidebar from '../Sidebar/Sidebar'
import './MainLayout.less'

const MainLayout = () => {
  return (
    <div className="main-layout">
      <Header />
      <div className="main-layout-body">
        <Sidebar />
        <main className="main-layout-content">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default MainLayout

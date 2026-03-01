import { NavLink } from 'react-router-dom'
import './Sidebar.less'

const Sidebar = () => {
  const menuItems = [
    { path: '/works', label: '作品列表', icon: 'book' },
    { path: '/optimization', label: 'AI优化', icon: 'ai' },
    { path: '/publish', label: '一键发布', icon: 'publish' },
  ]

  return (
    <aside className="sidebar">
      <nav className="sidebar-nav">
        {menuItems.map((item) => (
          <NavLink
            key={item.path}
            to={item.path}
            className={({ isActive }) => 
              `sidebar-item ${isActive ? 'sidebar-item-active' : ''}`
            }
          >
            <span className="sidebar-item-icon">
              {item.icon === 'book' && '📚'}
              {item.icon === 'ai' && '🤖'}
              {item.icon === 'publish' && '📤'}
            </span>
            <span className="sidebar-item-label">{item.label}</span>
          </NavLink>
        ))}
      </nav>
    </aside>
  )
}

export default Sidebar

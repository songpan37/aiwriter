import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useUserStore } from '@/store'
import './Header.less'

const Header = () => {
  const navigate = useNavigate()
  const { user, logout } = useUserStore()
  const [showDropdown, setShowDropdown] = useState(false)

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  const handleProfile = () => {
    navigate('/profile')
  }

  return (
    <header className="header">
      <div className="header-left">
        <h1 className="header-logo">AI Writer</h1>
      </div>
      <div className="header-center">
        <select className="header-work-select" defaultValue="">
          <option value="" disabled>选择作品</option>
        </select>
      </div>
      <div className="header-right">
        <span className="header-notification">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
            <path d="M13.73 21a2 2 0 0 1-3.46 0" />
          </svg>
        </span>
        <div className="header-user" onClick={() => setShowDropdown(!showDropdown)}>
          <img 
            src={user?.avatar || 'https://api.dicebear.com/7.x/avataaars/svg?seed=default'} 
            alt="avatar" 
            className="header-avatar"
          />
          <span className="header-username">{user?.username || '用户'}</span>
          {showDropdown && (
            <div className="header-dropdown">
              <div className="header-dropdown-item" onClick={handleProfile}>个人设置</div>
              <div className="header-dropdown-item" onClick={handleLogout}>退出登录</div>
            </div>
          )}
        </div>
      </div>
    </header>
  )
}

export default Header

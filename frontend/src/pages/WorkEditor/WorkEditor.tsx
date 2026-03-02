import { Outlet, useParams, useNavigate } from 'react-router-dom'
import './WorkEditor.less'

const menuItems = [
  { path: 'basic', label: '基本信息' },
  { path: 'volumes', label: '分卷信息' },
  { path: 'chapters', label: '章节信息' },
]

const WorkEditor = () => {
  const { workId } = useParams()
  const navigate = useNavigate()

  return (
    <div className="work-editor">
      <div className="work-editor-nav">
        <h3 className="work-editor-title">作品编辑 {workId ? `#${workId}` : ''}</h3>
        <nav className="work-editor-menu">
          {menuItems.map((item) => (
            <button
              key={item.path}
              className="menu-item"
              onClick={() => navigate(item.path)}
            >
              {item.label}
            </button>
          ))}
        </nav>
      </div>
      <div className="work-editor-content">
        <Outlet />
      </div>
    </div>
  )
}

export default WorkEditor

import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '@/api'
import './WorksList.less'

interface Work {
  id: number
  title: string
  cover?: string
  chapterCount: number
  wordCount: number
  updatedAt: string
  category?: string
}

const WorksList = () => {
  const navigate = useNavigate()
  const [works, setWorks] = useState<Work[]>([])
  const [loading, setLoading] = useState(true)
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid')

  useEffect(() => {
    loadWorks()
  }, [])

  const loadWorks = async () => {
    try {
      const response = await api.get('/works')
      if (response.code === 0) {
        setWorks(response.data.list || [])
      }
    } catch (error) {
      console.error('Failed to load works:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async (id: number, e: React.MouseEvent) => {
    e.stopPropagation()
    if (confirm('确定要删除这个作品吗？')) {
      try {
        await api.delete(`/works/${id}`)
        setWorks(works.filter(w => w.id !== id))
      } catch (error) {
        console.error('Failed to delete work:', error)
      }
    }
  }

  const formatDate = (dateStr: string) => {
    return new Date(dateStr).toLocaleDateString('zh-CN')
  }

  const formatWordCount = (count: number) => {
    if (count >= 10000) {
      return `${(count / 10000).toFixed(1)}万字`
    }
    return `${count}字`
  }

  if (loading) {
    return <div className="works-loading">加载中...</div>
  }

  return (
    <div className="works-list">
      <div className="works-header">
        <h2 className="works-title">作品列表</h2>
        <div className="works-actions">
          <div className="works-view-toggle">
            <button
              className={`view-btn ${viewMode === 'grid' ? 'active' : ''}`}
              onClick={() => setViewMode('grid')}
            >
              网格
            </button>
            <button
              className={`view-btn ${viewMode === 'list' ? 'active' : ''}`}
              onClick={() => setViewMode('list')}
            >
              列表
            </button>
          </div>
          <button className="works-add-btn">+ 新增作品</button>
        </div>
      </div>

      {works.length === 0 ? (
        <div className="works-empty">
          <p>暂无作品</p>
          <button className="works-add-btn">创建第一个作品</button>
        </div>
      ) : (
        <div className={`works-content works-content-${viewMode}`}>
          {works.map((work) => (
            <div
              key={work.id}
              className="work-card"
              onClick={() => navigate(`/works/${work.id}`)}
            >
              <div className="work-cover">
                {work.cover ? (
                  <img src={work.cover} alt={work.title} />
                ) : (
                  <div className="work-cover-placeholder">📚</div>
                )}
              </div>
              <div className="work-info">
                <h3 className="work-title">{work.title}</h3>
                <div className="work-meta">
                  <span>📖 {work.chapterCount}章</span>
                  <span>✍️ {formatWordCount(work.wordCount)}</span>
                </div>
                <div className="work-time">{formatDate(work.updatedAt)}</div>
              </div>
              <button
                className="work-delete"
                onClick={(e) => handleDelete(work.id, e)}
              >
                🗑️
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default WorksList

import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '@/api'
import './WorkWizard.less'

interface WorkWizardProps {
  visible: boolean
  onClose: () => void
}

interface FormData {
  category: string
  referenceWorks: number[]
  title: string
  authorName: string
  synopsis: string
  chaptersPerVolume: number
  wordsPerChapter: number
  cover?: string
}

const CATEGORIES = [
  { value: 'xuanhuan', label: '玄幻' },
  { value: 'dushi', label: '都市' },
  { value: 'lishi', label: '历史' },
  { value: 'kehuan', label: '科幻' },
  { value: 'youxi', label: '游戏' },
  { value: 'qita', label: '其他' },
]

const WorkWizard = ({ visible, onClose }: WorkWizardProps) => {
  const navigate = useNavigate()
  const [step, setStep] = useState(1)
  const [loading, setLoading] = useState(false)
  const [formData, setFormData] = useState<FormData>({
    category: '',
    referenceWorks: [],
    title: '',
    authorName: '',
    synopsis: '',
    chaptersPerVolume: 10,
    wordsPerChapter: 3000,
    cover: '',
  })

  if (!visible) return null

  const handleNext = () => {
    if (step === 1 && !formData.category) {
      alert('请选择作品类别')
      return
    }
    if (step === 3) {
      if (!formData.title.trim()) {
        alert('请输入书名')
        return
      }
      if (!formData.synopsis.trim()) {
        alert('请输入简介')
        return
      }
      handleSubmit()
      return
    }
    setStep(step + 1)
  }

  const handleBack = () => {
    setStep(step - 1)
  }

  const handleSubmit = async () => {
    setLoading(true)
    try {
      const categoryId = formData.category ? parseInt(formData.category) : 0
      const response = await api.post('/works', {
        title: formData.title,
        intro: formData.synopsis,
        category_id: categoryId > 0 ? categoryId : 1,
        chapters_per_volume: formData.chaptersPerVolume,
        words_per_chapter: formData.wordsPerChapter,
        cover: formData.cover,
      }) as { code: number; data: { id: number } }

      if (response.code === 0) {
        onClose()
        navigate(`/works/${response.data.id}`)
      } else {
        alert('创建作品失败')
      }
    } catch (error) {
      console.error('Failed to create work:', error)
      alert('创建作品失败')
    } finally {
      setLoading(false)
    }
  }

  const handleCategorySelect = (category: string) => {
    setFormData({ ...formData, category })
  }

  const renderStep1 = () => (
    <div className="wizard-step">
      <h3>选择作品类别</h3>
      <p className="wizard-step-desc">请选择您的作品所属类别</p>
      <div className="wizard-categories">
        {CATEGORIES.map((cat) => (
          <div
            key={cat.value}
            className={`wizard-category ${formData.category === cat.value ? 'selected' : ''}`}
            onClick={() => handleCategorySelect(cat.value)}
          >
            {cat.label}
          </div>
        ))}
      </div>
    </div>
  )

  const renderStep2 = () => (
    <div className="wizard-step">
      <h3>选择同类作品参考</h3>
      <p className="wizard-step-desc">选择您希望参考的同类作品（可选）</p>
      <div className="wizard-reference">
        <p className="wizard-reference-hint">暂无参考作品</p>
      </div>
    </div>
  )

  const renderStep3 = () => (
    <div className="wizard-step">
      <h3>填写基本信息</h3>
      <div className="wizard-form">
        <div className="wizard-field">
          <label>书名 *</label>
          <input
            type="text"
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
            placeholder="请输入作品名称"
          />
        </div>
        <div className="wizard-field">
          <label>作者署名</label>
          <input
            type="text"
            value={formData.authorName}
            onChange={(e) => setFormData({ ...formData, authorName: e.target.value })}
            placeholder="请输入作者名称"
          />
        </div>
        <div className="wizard-field">
          <label>简介 *</label>
          <textarea
            value={formData.synopsis}
            onChange={(e) => setFormData({ ...formData, synopsis: e.target.value })}
            placeholder="请输入作品简介"
            rows={4}
          />
        </div>
        <div className="wizard-field-row">
          <div className="wizard-field">
            <label>每卷章节数</label>
            <input
              type="number"
              value={formData.chaptersPerVolume}
              onChange={(e) => setFormData({ ...formData, chaptersPerVolume: parseInt(e.target.value) || 10 })}
              min={1}
              max={100}
            />
          </div>
          <div className="wizard-field">
            <label>每章字数</label>
            <input
              type="number"
              value={formData.wordsPerChapter}
              onChange={(e) => setFormData({ ...formData, wordsPerChapter: parseInt(e.target.value) || 3000 })}
              min={100}
              max={100000}
              step={100}
            />
          </div>
        </div>
      </div>
    </div>
  )

  return (
    <div className="wizard-overlay">
      <div className="wizard-modal">
        <div className="wizard-header">
          <h2>新增作品</h2>
          <button className="wizard-close" onClick={onClose}>×</button>
        </div>
        <div className="wizard-progress">
          <div className={`wizard-progress-step ${step >= 1 ? 'active' : ''}`}>
            <span className="wizard-progress-num">1</span>
            <span className="wizard-progress-label">选择类别</span>
          </div>
          <div className="wizard-progress-line" />
          <div className={`wizard-progress-step ${step >= 2 ? 'active' : ''}`}>
            <span className="wizard-progress-num">2</span>
            <span className="wizard-progress-label">选择参考</span>
          </div>
          <div className="wizard-progress-line" />
          <div className={`wizard-progress-step ${step >= 3 ? 'active' : ''}`}>
            <span className="wizard-progress-num">3</span>
            <span className="wizard-progress-label">填写信息</span>
          </div>
        </div>
        <div className="wizard-content">
          {step === 1 && renderStep1()}
          {step === 2 && renderStep2()}
          {step === 3 && renderStep3()}
        </div>
        <div className="wizard-footer">
          {step > 1 && (
            <button className="wizard-btn wizard-btn-secondary" onClick={handleBack}>
              上一步
            </button>
          )}
          <button
            className="wizard-btn wizard-btn-primary"
            onClick={handleNext}
            disabled={loading}
          >
            {loading ? '创建中...' : step === 3 ? '确认创建' : '下一步'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default WorkWizard

import { useState } from 'react'
import './Publish.less'

const Publish = () => {
  const [selectedWork, setSelectedWork] = useState('')
  const [selectedChapters, setSelectedChapters] = useState<number[]>([])
  const [targetWordCount, setTargetWordCount] = useState(2000)
  const [selectedPlatform, setSelectedPlatform] = useState('')

  const handleSelectAll = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.checked) {
      setSelectedChapters([1, 2, 3])
    } else {
      setSelectedChapters([])
    }
  }

  return (
    <div className="publish">
      <h2 className="publish-title">一键发布</h2>
      
      <div className="publish-form">
        <div className="form-field">
          <label>选择作品</label>
          <select 
            value={selectedWork} 
            onChange={(e) => setSelectedWork(e.target.value)}
          >
            <option value="">请选择作品</option>
          </select>
        </div>

        <div className="form-field">
          <label>选择章节 ({selectedChapters.length}章)</label>
          <div className="chapter-select">
            <label className="chapter-checkbox">
              <input type="checkbox" onChange={handleSelectAll} /> 全选
            </label>
          </div>
        </div>

        <div className="form-field">
          <label>目标字数（每章）</label>
          <input 
            type="number" 
            value={targetWordCount}
            onChange={(e) => setTargetWordCount(Number(e.target.value))}
          />
          <button className="resplit-btn">重新划分</button>
        </div>

        <div className="form-field">
          <label>发布平台</label>
          <select 
            value={selectedPlatform}
            onChange={(e) => setSelectedPlatform(e.target.value)}
          >
            <option value="">请选择平台</option>
            <option value="fanqie">番茄小说</option>
            <option value="qidian">起点中文网</option>
          </select>
        </div>

        <div className="publish-actions">
          <button className="publish-btn">发布</button>
        </div>
      </div>
    </div>
  )
}

export default Publish

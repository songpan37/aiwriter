import { useState } from 'react'
import './Optimization.less'

const defaultSteps = [
  { id: 1, name: '信息认知时序优化' },
  { id: 2, name: '人物对话优化' },
  { id: 3, name: '滥用表达优化' },
  { id: 4, name: '抽象设定说明优化' },
]

const Optimization = () => {
  const [steps] = useState(defaultSteps)
  const [activeStep, setActiveStep] = useState(1)

  return (
    <div className="optimization">
      <div className="optimization-sidebar">
        <h2 className="optimization-title">AI优化</h2>
        <div className="optimization-steps">
          {steps.map((step, index) => (
            <button
              key={step.id}
              className={`step-item ${activeStep === step.id ? 'active' : ''}`}
              onClick={() => setActiveStep(step.id)}
            >
              <span className="step-number">{index + 1}</span>
              <span className="step-name">{step.name}</span>
            </button>
          ))}
        </div>
        <button className="optimization-add-btn">+ 添加步骤</button>
      </div>
      <div className="optimization-content">
        <div className="optimization-review">
          <h3>审阅结论</h3>
          <p>AI分析结果将显示在这里</p>
        </div>
        <div className="optimization-diff">
          <div className="diff-original">
            <h3>原文</h3>
            <textarea placeholder="输入待优化的文本..." />
          </div>
          <div className="diff-optimized">
            <h3>优化后</h3>
            <textarea placeholder="优化后的文本将显示在这里" readOnly />
          </div>
        </div>
        <div className="optimization-actions">
          <button className="optimization-execute-btn">执行优化</button>
        </div>
      </div>
    </div>
  )
}

export default Optimization

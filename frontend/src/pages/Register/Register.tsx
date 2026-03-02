import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '@/api'
import './Register.less'

const Register = () => {
  const navigate = useNavigate()
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
  })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')

    if (formData.password !== formData.confirmPassword) {
      setError('两次输入的密码不一致')
      return
    }

    setLoading(true)

    try {
      const { confirmPassword, ...data } = formData
      const response = await api.post('/auth/register', data) as {code: number, message: string}
      if (response.code === 0) {
        navigate('/login')
      } else {
        setError(response.message || '注册失败')
      }
    } catch (err: unknown) {
      setError('注册失败，请稍后重试')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="register-page">
      <div className="register-card">
        <h1 className="register-title">注册</h1>
        <p className="register-subtitle">创建您的AI Writer账号</p>
        
        <form onSubmit={handleSubmit} className="register-form">
          {error && <div className="register-error">{error}</div>}
          
          <div className="register-field">
            <input
              type="text"
              placeholder="用户名"
              value={formData.username}
              onChange={(e) => setFormData({ ...formData, username: e.target.value })}
              required
            />
          </div>

          <div className="register-field">
            <input
              type="email"
              placeholder="邮箱"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              required
            />
          </div>
          
          <div className="register-field">
            <input
              type="password"
              placeholder="密码"
              value={formData.password}
              onChange={(e) => setFormData({ ...formData, password: e.target.value })}
              required
            />
          </div>

          <div className="register-field">
            <input
              type="password"
              placeholder="确认密码"
              value={formData.confirmPassword}
              onChange={(e) => setFormData({ ...formData, confirmPassword: e.target.value })}
              required
            />
          </div>
          
          <button type="submit" className="register-button" disabled={loading}>
            {loading ? '注册中...' : '注册'}
          </button>
        </form>
        
        <div className="register-footer">
          <span>已有账号？</span>
          <a href="/login">立即登录</a>
        </div>
      </div>
    </div>
  )
}

export default Register

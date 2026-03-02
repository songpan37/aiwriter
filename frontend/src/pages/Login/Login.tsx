import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useUserStore } from '@/store'
import api from '@/api'
import './Login.less'

const Login = () => {
  const navigate = useNavigate()
  const { setToken, setUser } = useUserStore()
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    try {
      const response = await api.post('/auth/login', formData) as {code: number, message: string, data: {token: string, user: any}}
      if (response.code === 0) {
        setToken(response.data.token)
        setUser(response.data.user)
        navigate('/works')
      } else {
        setError(response.message || '登录失败')
      }
    } catch (err: unknown) {
      setError('用户名或密码错误')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="login-page">
      <div className="login-card">
        <h1 className="login-title">AI Writer</h1>
        <p className="login-subtitle">AI写作辅助系统</p>
        
        <form onSubmit={handleSubmit} className="login-form">
          {error && <div className="login-error">{error}</div>}
          
          <div className="login-field">
            <input
              type="text"
              placeholder="用户名"
              value={formData.username}
              onChange={(e) => setFormData({ ...formData, username: e.target.value })}
              required
            />
          </div>
          
          <div className="login-field">
            <input
              type="password"
              placeholder="密码"
              value={formData.password}
              onChange={(e) => setFormData({ ...formData, password: e.target.value })}
              required
            />
          </div>
          
          <button type="submit" className="login-button" disabled={loading}>
            {loading ? '登录中...' : '登录'}
          </button>
        </form>
        
        <div className="login-footer">
          <span>还没有账号？</span>
          <a href="/register">立即注册</a>
        </div>
      </div>
    </div>
  )
}

export default Login

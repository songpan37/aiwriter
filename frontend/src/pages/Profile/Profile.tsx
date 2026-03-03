import { useState, useEffect } from 'react'
import { useUserStore } from '@/store'
import { getProfile } from '@/api'
import './Profile.less'

interface ProfileData {
  id: number
  username: string
  email: string
  avatar: string
  created_at: string
}

const Profile = () => {
  const { user, setUser } = useUserStore()
  const [profile, setProfile] = useState<ProfileData | null>(null)
  const [loading, setLoading] = useState(true)
  const [editing, setEditing] = useState(false)
  const [formData, setFormData] = useState({
    username: '',
    email: '',
  })
  const [saving, setSaving] = useState(false)
  const [message, setMessage] = useState({ type: '', text: '' })

  useEffect(() => {
    loadProfile()
  }, [])

  const loadProfile = async () => {
    try {
      const response = await getProfile() as { code: number; data: ProfileData }
      if (response.code === 0) {
        setProfile(response.data)
        setFormData({
          username: response.data.username,
          email: response.data.email,
        })
      }
    } catch (err) {
      console.error('Failed to load profile:', err)
    } finally {
      setLoading(false)
    }
  }

  const handleSave = async () => {
    setSaving(true)
    setMessage({ type: '', text: '' })
    
    try {
      await new Promise(resolve => setTimeout(resolve, 500))
      setMessage({ type: 'success', text: '保存成功' })
      setEditing(false)
      if (user) {
        setUser({ ...user, username: formData.username, email: formData.email })
      }
    } catch (err) {
      setMessage({ type: 'error', text: '保存失败' })
    } finally {
      setSaving(false)
    }
  }

  const handleCancel = () => {
    if (profile) {
      setFormData({
        username: profile.username,
        email: profile.email,
      })
    }
    setEditing(false)
    setMessage({ type: '', text: '' })
  }

  if (loading) {
    return <div className="profile-page"><div className="profile-loading">加载中...</div></div>
  }

  return (
    <div className="profile-page">
      <div className="profile-card">
        <h2 className="profile-title">个人设置</h2>
        
        {message.text && (
          <div className={`profile-message ${message.type}`}>{message.text}</div>
        )}
        
        <div className="profile-avatar-section">
          <img 
            src={profile?.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${profile?.username}`} 
            alt="avatar" 
            className="profile-avatar"
          />
          <button className="profile-avatar-button">更换头像</button>
        </div>
        
        <div className="profile-form">
          <div className="profile-field">
            <label>用户名</label>
            {editing ? (
              <input
                type="text"
                value={formData.username}
                onChange={(e) => setFormData({ ...formData, username: e.target.value })}
              />
            ) : (
              <div className="profile-field-value">{profile?.username}</div>
            )}
          </div>
          
          <div className="profile-field">
            <label>邮箱</label>
            {editing ? (
              <input
                type="email"
                value={formData.email}
                onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              />
            ) : (
              <div className="profile-field-value">{profile?.email}</div>
            )}
          </div>
          
          <div className="profile-field">
            <label>创建时间</label>
            <div className="profile-field-value">
              {profile?.created_at ? new Date(profile.created_at).toLocaleString('zh-CN') : '-'}
            </div>
          </div>
        </div>
        
        <div className="profile-actions">
          {editing ? (
            <>
              <button className="profile-save-button" onClick={handleSave} disabled={saving}>
                {saving ? '保存中...' : '保存'}
              </button>
              <button className="profile-cancel-button" onClick={handleCancel}>
                取消
              </button>
            </>
          ) : (
            <button className="profile-edit-button" onClick={() => setEditing(true)}>
              编辑资料
            </button>
          )}
        </div>
      </div>
    </div>
  )
}

export default Profile

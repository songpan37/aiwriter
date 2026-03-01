package response

import "time"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type WorkResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Cover        string    `json:"cover"`
	Intro        string    `json:"intro"`
	ChapterCount int       `json:"chapter_count"`
	WordCount    int       `json:"word_count"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type WorkListResponse struct {
	List       []WorkResponse `json:"list"`
	Pagination Pagination     `json:"pagination"`
}

type VolumeResponse struct {
	ID           uint      `json:"id"`
	WorkID       uint      `json:"work_id"`
	VolumeIndex  int       `json:"volume_index"`
	Title        string    `json:"title"`
	ChapterCount int       `json:"chapter_count"`
	WordCount    int       `json:"word_count"`
	Summary      string    `json:"summary"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ChapterResponse struct {
	ID           uint      `json:"id"`
	WorkID       uint      `json:"work_id"`
	VolumeID     uint      `json:"volume_id"`
	ChapterIndex int       `json:"chapter_index"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	WordCount    int       `json:"word_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

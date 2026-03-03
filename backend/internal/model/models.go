package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	AvatarKey string    `json:"avatar_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Work struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	CategoryID   uint      `json:"category_id"`
	Title        string    `json:"title"`
	CoverKey     string    `json:"cover_key"`
	ContentKey   string    `json:"content_key"`
	Intro        string    `json:"intro"`
	ChapterCount int       `json:"chapter_count"`
	WordCount    int       `json:"word_count"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type WorkCategory struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Volume struct {
	ID            uint      `json:"id"`
	WorkID        uint      `json:"work_id"`
	VolumeIndex   int       `json:"volume_index"`
	Title         string    `json:"title"`
	ChapterCount  int       `json:"chapter_count"`
	WordCount     int       `json:"word_count"`
	Summary       string    `json:"summary"`
	Characters    string    `json:"characters"`
	PlotUnits     string    `json:"plot_units"`
	Relationships string    `json:"relationships"`
	Goals         string    `json:"goals"`
	Conflicts     string    `json:"conflicts"`
	Hooks         string    `json:"hooks"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Chapter struct {
	ID           uint      `json:"id"`
	WorkID       uint      `json:"work_id"`
	VolumeID     uint      `json:"volume_id"`
	ChapterIndex int       `json:"chapter_index"`
	Title        string    `json:"title"`
	ContentKey   string    `json:"content_key"`
	Summary      string    `json:"summary"`
	TimeSpace    string    `json:"time_space"`
	Characters   string    `json:"characters"`
	Scenes       string    `json:"scenes"`
	WordCount    int       `json:"word_count"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Scene struct {
	ID         uint      `json:"id"`
	ChapterID  uint      `json:"chapter_id"`
	ContentKey string    `json:"content_key"`
	Begin      string    `json:"begin"`
	Conflict   string    `json:"conflict"`
	Turn       string    `json:"turn"`
	Result     string    `json:"result"`
	Style      string    `json:"style"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OptimizationStep struct {
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	Name           string    `json:"name"`
	ReviewPrompt   string    `json:"review_prompt"`
	OptimizePrompt string    `json:"optimize_prompt"`
	StepOrder      int       `json:"step_order"`
	IsDefault      bool      `json:"is_default"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type OptimizationRecord struct {
	ID               uint      `json:"id"`
	WorkID           uint      `json:"work_id"`
	ChapterID        uint      `json:"chapter_id"`
	StepID           uint      `json:"step_id"`
	OriginalKey      string    `json:"original_key"`
	OptimizedKey     string    `json:"optimized_key"`
	ReviewConclusion string    `json:"review_conclusion"`
	CreatedAt        time.Time `json:"created_at"`
}

type PublishTask struct {
	ID              uint      `json:"id"`
	WorkID          uint      `json:"work_id"`
	Platform        string    `json:"platform"`
	ChapterIDs      string    `json:"chapter_ids"`
	SplitWordCount  int       `json:"split_word_count"`
	NewChapterNames string    `json:"new_chapter_names"`
	Status          int       `json:"status"`
	OutputKey       string    `json:"output_key"`
	Result          string    `json:"result"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Notification struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type MetaStore struct {
	Users      map[string]User         `json:"users"`
	Works      map[string]Work         `json:"works"`
	Volumes    map[string]Volume       `json:"volumes"`
	Chapters   map[string]Chapter      `json:"chapters"`
	Categories map[string]WorkCategory `json:"categories"`
	NextIDs    map[string]uint         `json:"next_ids"`
}

func NewMetaStore() *MetaStore {
	return &MetaStore{
		Users:      make(map[string]User),
		Works:      make(map[string]Work),
		Volumes:    make(map[string]Volume),
		Chapters:   make(map[string]Chapter),
		Categories: make(map[string]WorkCategory),
		NextIDs: map[string]uint{
			"user":         1,
			"work":         1,
			"volume":       1,
			"chapter":      1,
			"category":     1,
			"optimization": 1,
			"publish_task": 1,
		},
	}
}

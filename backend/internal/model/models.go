package model

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password  string     `json:"-" gorm:"size:255;not null"`
	Email     string     `json:"email" gorm:"uniqueIndex;size:255;not null"`
	Avatar    string     `json:"avatar" gorm:"size:500"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

type Work struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	UserID       uint       `json:"user_id" gorm:"index;not null"`
	CategoryID   uint       `json:"category_id" gorm:"index"`
	Title        string     `json:"title" gorm:"size:255;not null"`
	Cover        string     `json:"cover" gorm:"size:500"`
	Intro        string     `json:"intro" gorm:"type:text"`
	ChapterCount int        `json:"chapter_count" gorm:"default:0"`
	WordCount    int        `json:"word_count" gorm:"default:0"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index"`
}

type WorkCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:50;not null"`
	CreatedAt time.Time `json:"created_at"`
}

type Volume struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	WorkID        uint       `json:"work_id" gorm:"index;not null"`
	VolumeIndex   int        `json:"volume_index" gorm:"not null"`
	Title         string     `json:"title" gorm:"size:255;not null"`
	ChapterCount  int        `json:"chapter_count" gorm:"default:0"`
	WordCount     int        `json:"word_count" gorm:"default:0"`
	Summary       string     `json:"summary" gorm:"type:text"`
	Characters    string     `json:"characters" gorm:"type:json"`
	PlotUnits     string     `json:"plot_units" gorm:"type:json"`
	Relationships string     `json:"relationships" gorm:"type:text"`
	Goals         string     `json:"goals" gorm:"type:text"`
	Conflicts     string     `json:"conflicts" gorm:"type:text"`
	Hooks         string     `json:"hooks" gorm:"type:text"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"index"`
}

type Chapter struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	WorkID       uint       `json:"work_id" gorm:"index;not null"`
	VolumeID     uint       `json:"volume_id" gorm:"index"`
	ChapterIndex int        `json:"chapter_index" gorm:"not null"`
	Title        string     `json:"title" gorm:"size:255;not null"`
	Summary      string     `json:"summary" gorm:"type:text"`
	TimeSpace    string     `json:"time_space" gorm:"size:255"`
	Characters   string     `json:"characters" gorm:"type:json"`
	Scenes       string     `json:"scenes" gorm:"type:json"`
	WordCount    int        `json:"word_count" gorm:"default:0"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index"`
}

type Scene struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ChapterID uint      `json:"chapter_id" gorm:"index;not null"`
	Begin     string    `json:"begin" gorm:"type:text"`
	Conflict  string    `json:"conflict" gorm:"type:text"`
	Turn      string    `json:"turn" gorm:"type:text"`
	Result    string    `json:"result" gorm:"type:text"`
	Style     string    `json:"style" gorm:"size:50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OptimizationStep struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"index;not null"`
	Name           string    `json:"name" gorm:"size:100;not null"`
	ReviewPrompt   string    `json:"review_prompt" gorm:"type:text"`
	OptimizePrompt string    `json:"optimize_prompt" gorm:"type:text"`
	StepOrder      int       `json:"step_order" gorm:"not null"`
	IsDefault      bool      `json:"is_default" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type OptimizationRecord struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	WorkID           uint      `json:"work_id" gorm:"index;not null"`
	ChapterID        uint      `json:"chapter_id" gorm:"index"`
	StepID           uint      `json:"step_id" gorm:"index"`
	OriginalText     string    `json:"original_text" gorm:"type:text"`
	OptimizedText    string    `json:"optimized_text" gorm:"type:text"`
	ReviewConclusion string    `json:"review_conclusion" gorm:"type:text"`
	CreatedAt        time.Time `json:"created_at"`
}

type PublishTask struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	WorkID          uint      `json:"work_id" gorm:"index;not null"`
	Platform        string    `json:"platform" gorm:"size:50;not null"`
	ChapterIDs      string    `json:"chapter_ids" gorm:"type:json"`
	SplitWordCount  int       `json:"split_word_count"`
	NewChapterNames string    `json:"new_chapter_names" gorm:"type:json"`
	Status          int       `json:"status" gorm:"default:0"`
	Result          string    `json:"result" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	Content   string    `json:"content" gorm:"type:text"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type CreateWorkRequest struct {
	Title             string `json:"title" binding:"required"`
	Cover             string `json:"cover"`
	Intro             string `json:"intro"`
	CategoryID        uint   `json:"category_id"`
	ChaptersPerVolume int    `json:"chapters_per_volume"`
	WordsPerChapter   int    `json:"words_per_chapter"`
}

type UpdateWorkRequest struct {
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Intro      string `json:"intro"`
	CategoryID uint   `json:"category_id"`
}

type CreateVolumeRequest struct {
	VolumeIndex int    `json:"volume_index"`
	Title       string `json:"title" binding:"required"`
	Summary     string `json:"summary"`
}

type UpdateVolumeRequest struct {
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Characters string `json:"characters"`
	PlotUnits  string `json:"plot_units"`
}

type CreateChapterRequest struct {
	VolumeID     uint   `json:"volume_id"`
	ChapterIndex int    `json:"chapter_index"`
	Title        string `json:"title" binding:"required"`
	Summary      string `json:"summary"`
	TimeSpace    string `json:"time_space"`
	Characters   string `json:"characters"`
	Scenes       string `json:"scenes"`
}

type UpdateChapterRequest struct {
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	TimeSpace  string `json:"time_space"`
	Characters string `json:"characters"`
	Scenes     string `json:"scenes"`
}

type OptimizationExecuteRequest struct {
	StepID       uint   `json:"step_id" binding:"required"`
	OriginalText string `json:"original_text" binding:"required"`
}

type PublishPreviewRequest struct {
	ChapterIDs     []uint `json:"chapter_ids" binding:"required"`
	SplitWordCount int    `json:"split_word_count" binding:"required"`
}

type PublishExecuteRequest struct {
	WorkID         uint   `json:"work_id" binding:"required"`
	ChapterIDs     []uint `json:"chapter_ids" binding:"required"`
	SplitWordCount int    `json:"split_word_count"`
	Platform       string `json:"platform" binding:"required"`
}

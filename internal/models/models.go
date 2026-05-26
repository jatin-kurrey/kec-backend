package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Admin struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"default:'EDITOR'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Application struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FormType  string         `gorm:"not null" json:"form_type"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"index;not null" json:"email"`
	Phone     string         `gorm:"not null" json:"phone"`
	Data      datatypes.JSON `json:"data"`
	Status    string         `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Notice struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Type      string    `gorm:"not null" json:"type"` // academic, placement, event, general, student
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	Link      string    `json:"link"`
	Important bool      `gorm:"default:false" json:"important"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Department struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	ShortName     string         `json:"short_name"`
	Description   string         `json:"description"`
	Image         string         `json:"image"`
	StudentCount  int            `json:"student_count"`
	CourseCount   int            `json:"course_count"`
	PlacementRate string         `json:"placement_rate"`
	FacultyCount  int            `json:"faculty_count"`
	LabCount      int            `json:"lab_count"`
	Established   int            `json:"established"`
	Highlights    datatypes.JSON `json:"highlights"` // Array of strings
	Vision        string         `json:"vision"`
	Mission       datatypes.JSON `json:"mission"`      // Array of strings
	Programs      datatypes.JSON `json:"programs"`     // Array of objects {name, duration, seats}
	Facilities    datatypes.JSON `json:"facilities"`   // Array of strings
	Achievements  datatypes.JSON `json:"achievements"` // Array of strings
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Link          string         `json:"link"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Gallery struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"` // Cultural, Academic, Religious, National
	Description string    `json:"description"`
	ImageURL    string    `gorm:"not null" json:"image_url"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	Location    string    `json:"location"`
	Likes       int       `json:"likes"`
	Downloads   int       `json:"downloads"`
	Attendees   int       `json:"attendees"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
}

type Exam struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Duration               int            `json:"duration"` // in minutes
	NegativeMarking        float64        `gorm:"default:0" json:"negative_marking"`
	ShuffleQuestions       bool           `gorm:"default:false" json:"shuffle_questions"`
	BrowserLockdown        bool           `gorm:"default:false" json:"browser_lockdown"`
	ShowResultImmediately  bool           `gorm:"default:false" json:"show_result_immediately"`
	StartTime              time.Time      `json:"start_time"`
	EndTime                time.Time      `json:"end_time"`
	IsActive               bool           `gorm:"default:false" json:"is_active"`
	Questions              []Question     `gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE" json:"questions"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
}

type Event struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title            string    `gorm:"not null" json:"title"`
	Date             time.Time `json:"date"`
	Time             string    `json:"time"`
	Location         string    `json:"location"`
	Category         string    `json:"category"`
	Description      string    `json:"description"`
	ImageURL         string    `json:"image"`
	Attendees        int       `json:"attendees"`
	Status           string    `json:"status"` // upcoming, past
	RegistrationLink string    `json:"registration_link"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Question struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ExamID        uuid.UUID      `gorm:"type:uuid;not null" json:"exam_id"`
	Text          string         `gorm:"not null" json:"text"`
	Type          string         `gorm:"not null" json:"type"` // MCQ, INTEGER
	Options       datatypes.JSON `json:"options"`             // Array of strings
	CorrectAnswer string         `gorm:"not null" json:"correct_answer"`
	Points        int            `gorm:"default:1" json:"points"`
	ImageURL      string         `json:"image_url"`
	Subject       string         `json:"subject"`   // e.g. Mathematics, Physics
	Difficulty    string         `json:"difficulty"` // e.g. Easy, Medium, Hard
	CreatedAt     time.Time      `json:"created_at"`
}

type ExamResponse struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ExamID    uuid.UUID      `gorm:"type:uuid;not null" json:"exam_id"`
	StudentID string         `gorm:"index;not null" json:"student_id"`
	Name      string         `json:"name"`
	Responses datatypes.JSON `json:"responses"` // question_id -> answer
	Score     int            `json:"score"`
	Submitted bool           `gorm:"default:false" json:"submitted"`
	CreatedAt time.Time      `json:"created_at"`
}

type Leadership struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	Role           string         `gorm:"not null" json:"role"` // e.g. HOD, Principal, Director
	Department     string         `json:"department"`
	Image          string         `json:"image"`
	Qualification  string         `json:"qualification"`
	Experience     string         `json:"experience"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	Specialization string         `json:"specialization"`
	Achievements   datatypes.JSON `json:"achievements"` // Array of strings
	Bio            string         `json:"bio"`
	Priority       int            `gorm:"default:0" json:"priority"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type Faculty struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	SNo            int       `json:"sno"`
	Name           string    `gorm:"not null" json:"name"`
	Role           string    `gorm:"not null" json:"role"`
	Department     string    `json:"dept"`
	Image          string    `json:"image"`
	Qualification  string    `json:"qualification"`
	Experience     string    `json:"experience"`
	Specialization string    `json:"specialization"`
	Publications   int       `gorm:"default:0" json:"publications"`
	Projects       int       `gorm:"default:0" json:"projects"`
	Bio            string    `json:"bio"`
	Email          string    `json:"email"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Course struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string         `gorm:"not null" json:"title"`
	ShortName     string         `json:"short_name"`
	Icon          string         `json:"icon"` // Name of the lucide icon
	Duration      string         `json:"duration"`
	Seats         int            `json:"seats"`
	Credits       int            `json:"credits"`
	Description   string         `json:"description"`
	Highlights    datatypes.JSON `json:"highlights"` // Array of strings
	Career        string         `json:"career"`
	Color         string         `json:"color"`      // Gradient string
	IconColor     string         `json:"icon_color"` // CSS class
	BgColor       string         `json:"bg_color"`   // CSS class
	Department    string         `json:"department"`
	Eligibility   string         `json:"eligibility"`
	Fees          string         `json:"fees"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type ResearchArea struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title        string         `gorm:"not null" json:"title"`
	Description  string         `json:"description"`
	Icon         string         `json:"icon"` // Lucide icon name
	Projects     int            `json:"projects"`
	Publications int            `json:"publications"`
	Color        string         `json:"color"` // Tailwind gradient classes
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ResearchProject struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Department  string         `json:"department"`
	Funding     string         `json:"funding"`
	Duration    string         `json:"duration"`
	Status      string         `json:"status"` // Ongoing, Completed
	Description string         `json:"description"`
	Team        datatypes.JSON `json:"team"`     // Array of strings
	Outcomes    string         `json:"outcomes"`
	Image       string         `json:"image"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ResearchFacility struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Features    datatypes.JSON `json:"features"` // Array of strings
	Capacity    string         `json:"capacity"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ResearchStat struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Label     string    `gorm:"not null" json:"label"`
	Value     string    `json:"value"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// --- Admission Models ---

type AdmissionGuide struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Position      string `json:"position"`
	Qualification string `json:"qualification"`
	Experience    string `json:"experience"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Image         string `json:"image"`
	Message       string `json:"message" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AdmissionStep struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Icon        string `json:"icon"` // lucide icon name
	SortOrder   int    `json:"sort_order"`
}

type AdmissionEligibility struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Criteria  string `json:"criteria" gorm:"type:text"`
	SortOrder int    `json:"sort_order"`
}

type AdmissionDocument struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type AdmissionFee struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Program        string `json:"program"`
	TuitionFee     string `json:"tuitionFee"`
	DevelopmentFee string `json:"developmentFee"`
	Total          string `json:"total"`
	SortOrder      int    `json:"sort_order"`
}

type CampusFacility struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Category    string         `json:"category"` // Essential, Desirable, Infrastructure, Lab
	Icon        string         `json:"icon"`     // Lucide icon name
	Available   bool           `gorm:"default:true" json:"available"`
	Department  string         `json:"department"`  // For labs
	Area        int            `json:"area"`        // For labs
	Description string         `json:"description"`
	Features    datatypes.JSON `json:"features"`    // Array of strings for infrastructure
	SortOrder   int            `json:"sort_order"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type CampusStat struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Label     string    `gorm:"not null" json:"label"`
	Value     string    `json:"value"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Alumni struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name            string         `gorm:"not null" json:"name"`
	Batch           string         `json:"batch"`
	Branch          string         `json:"branch"`
	Degree          string         `json:"degree"`
	Location        string         `json:"location"`
	Company         string         `json:"company"`
	CurrentPosition string         `json:"currentPosition"`
	Story           string         `json:"story"`
	ImageURL        string         `json:"image_url"`
	Photo           string         `gorm:"-" json:"photo"` // Virtual field for frontend compatibility
	Achievements    datatypes.JSON `json:"achievements"`    // Array of strings
	IsFeatured      bool           `gorm:"default:false" json:"is_featured"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type AlumniStat struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Label     string    `gorm:"not null" json:"label"`
	Value     string    `json:"value"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

type PlacementStat struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Label     string    `gorm:"not null" json:"label"`
	Value     string    `json:"value"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

type Recruiter struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	LogoURL   string    `json:"logo_url"`
	Logo      string    `gorm:"-" json:"logo"` // Virtual field for frontend compatibility
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"created_at"`
}

type PlacementTestimonial struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Role       string    `json:"role"`
	Company    string    `json:"company"`
	Department string    `json:"department"`
	Quote      string    `json:"quote"`
	ImageURL   string    `json:"image_url"`
	Image      string    `gorm:"-" json:"image"` // Virtual field for frontend compatibility
	CreatedAt  time.Time `json:"created_at"`
}

type PressMedia struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Source      string    `json:"source"`
	Date        time.Time `json:"date"`
	Link        string    `json:"link"`
	Type        string    `json:"type"` // "article" or "video"
	ImageURL    string    `json:"image_url"`
	Image       string    `gorm:"-" json:"image"` // Virtual field for frontend compatibility
	CreatedAt   time.Time `json:"created_at"`
}

type SystemSetting struct {
	Key       string    `gorm:"primaryKey" json:"key"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

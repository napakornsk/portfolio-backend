package models

// Intro struct corresponding to the Intro message
type Intro struct {
	ID          int64  `gorm:"primaryKey"`
	Pos         string `gorm:"type:varchar(100)"`
	Email       string `gorm:"type:varchar(100)"`
	MobilePhone string `gorm:"type:varchar(20)"`
	Address     string `gorm:"type:text"`
	Summary     string `gorm:"type:text"`
	CvID        int64  `gorm:"not null"` // Foreign key for CV
}

func (Intro) TableName() string {
	return "intro"
}

// WorkExperience struct corresponding to the WorkExperience message
type WorkExperience struct {
	ID          int64      `gorm:"primaryKey"`
	CompanyName string     `gorm:"type:varchar(100)"`
	Role        string     `gorm:"type:varchar(100)"`
	Description string     `gorm:"type:text"`
	StartDate   string     `gorm:"type:varchar(7)"` // Format YYYY-MM
	EndDate     string     `gorm:"type:varchar(7)"`
	WorkDesc    []WorkDesc `gorm:"foreignKey:WorkID"` // One-to-many relationship with WorkDesc
	CvID        int64      `gorm:"not null"`          // Foreign key for CV
}

func (WorkExperience) TableName() string {
	return "work_experience"
}

// WorkDesc struct corresponding to the WorkDesc message
type WorkDesc struct {
	ID          int64  `gorm:"primaryKey"`
	Description string `gorm:"type:text"`
	WorkID      int64  `gorm:"not null"` // Foreign key for WorkExperience
}

func (WorkDesc) TableName() string {
	return "work_desc"
}

// Education struct corresponding to the Education message
type Education struct {
	ID         int64  `gorm:"primaryKey"`
	SchoolName string `gorm:"type:varchar(100)"`
	Degree     string `gorm:"type:varchar(100)"`
	Field      string `gorm:"type:varchar(100)"`
	StartDate  string `gorm:"type:varchar(7)"` // Format YYYY-MM
	EndDate    string `gorm:"type:varchar(7)"`
	CvID       int64  `gorm:"not null"` // Foreign key for CV
}

func (Education) TableName() string {
	return "education"
}

// Skill struct corresponding to the Skill message
type Skill struct {
	ID        int64  `gorm:"primaryKey"`
	SkillName string `gorm:"type:varchar(100)"`
	CvID      int64  `gorm:"not null"` // Foreign key for CV
}

func (Skill) TableName() string {
	return "skill"
}

// CV struct corresponding to the CV message
type CV struct {
	ID          int64            `gorm:"primaryKey"`
	Intro       Intro            `gorm:"foreignKey:CvID"` // One-to-one relationship with Intro
	Experiences []WorkExperience `gorm:"foreignKey:CvID"` // One-to-many relationship with WorkExperience
	Education   []Education      `gorm:"foreignKey:CvID"` // One-to-many relationship with Education
	Skills      []Skill          `gorm:"foreignKey:CvID"` // One-to-many relationship with Skill
	UID         int64            `gorm:"not null"`        // Foreign key for user ID or another unique identifier
}

func (CV) TableName() string {
	return "cv"
}

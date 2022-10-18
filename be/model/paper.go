package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// 论文

type Paper struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title   string `gorm:"not null;"`
	Authors string

	Chapter int    // 卷
	Phase   int    // 期
	Doi     string // DOI
	Year    string // 年
	Issue   string // 发表时间

	Url          string // 文献详情页链接
	Introduction string // 摘要

	Keywords []Keyword `gorm:"many2many:paper_keywords;" json:"-"`

	// belongs to a magazine
	MagazineName string   `json:"-"`
	Magazine     Magazine `gorm:"foreignKey:MagazineName;references:Name;" json:"-"`
}

func (p *Paper) TableUnique() [][]string {
	return [][]string{
		{"Title", "Authors"},
	}
}

func (p *Paper) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

type PaperModel struct {
	db *gorm.DB
}

func GetPaperModel() *PaperModel {
	return &PaperModel{db_gorm}
}

// CreatePaper 添加文章
func (m *PaperModel) CreatePaper(paper *Paper) error {
	result := m.db.Create(paper)
	return result.Error
}

// CreatePapers 添加文章
func (m *PaperModel) CreatePapers(papers []Paper) error {
	result := m.db.Create(papers)
	return result.Error
}

// GetPaperById 通过Id获取论文
func (m *PaperModel) GetPaperById(id string) (Paper, error) {
	paper := Paper{}
	result := m.db.Preload("Keywords").Where("id = ?", id).First(&paper)
	return paper, result.Error
}

//// GetPaperByIdWithMagazine 通过Id获取论文以及期刊信息
//func (m *PaperModel) GetPaperByIdWithMagazine(id string) (Paper, error) {
//
//}

// DeletePaperByTitle 通过标题删除文章
func (m *PaperModel) DeletePaperByTitle(title string) error {
	result := m.db.Where("title = ?", title).Delete(&Paper{})

	return result.Error
}

// DeletePaperById 通过ID删除文章
func (m *PaperModel) DeletePaperById(id string) error {
	result := m.db.Where("id = ?", id).Delete(&Paper{})

	return result.Error
}

// DeletePaperByDoi 通过DOI删除文章
func (m *PaperModel) DeletePaperByDoi(doi string) error {
	result := m.db.Where("doi = ?", doi).Delete(&Paper{})

	return result.Error
}

// GetPaperByPaperDoi 通过DOI获取文章
func (m *PaperModel) GetPaperByPaperDoi(doi string) (Paper, error) {
	paper := Paper{}

	result := m.db.Where("doi = ?", doi).First(&paper)

	return paper, result.Error
}

// GetPaperByPaperTitle 通过Title获取文章
func (m *PaperModel) GetPaperByPaperTitle(title string) (Paper, error) {
	paper := Paper{}

	result := m.db.Where("title = ?", title).First(&paper)

	return paper, result.Error
}

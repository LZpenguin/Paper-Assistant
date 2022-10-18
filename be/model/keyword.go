package model

import "gorm.io/gorm"

type Keyword struct {
	Explain string `gorm:"primarykey;uniqueIndex;not null;"`
}

type KeywordModel struct {
	db *gorm.DB
}

func GetKeywordModel() *KeywordModel {
	return &KeywordModel{db_gorm}
}

// CreateKeyword Add a keyword
func (m *KeywordModel) CreateKeyword(keyword Keyword) error {
	result := m.db.Create(&keyword)
	return result.Error
}

// CreateManyKeyword Add many keywords
func (m *KeywordModel) CreateManyKeyword(keywords []Keyword) error {
	result := m.db.Create(keywords)
	return result.Error
}

// GetKeywordById Get keyword by ID
func (m *KeywordModel) GetKeywordById(id uint) (Keyword, error) {
	keyword := Keyword{}
	result := m.db.First(&keyword, id)
	return keyword, result.Error
}

// GetAllKeywords get all keywords
func (m *KeywordModel) GetAllKeywords() ([]Keyword, error) {
	keywords := []Keyword{}
	result := m.db.Find(&keywords)
	return keywords, result.Error
}

// DeleteKeywordById Delete keyword by ID
func (m *KeywordModel) DeleteKeywordById(id uint) error {
	result := m.db.Delete(&Keyword{}, id)

	return result.Error
}

func (m *KeywordModel) SearchAll() ([]Keyword, error) {
	var kws []Keyword
	result := m.db.Find(&kws)
	return kws, result.Error
}

func (m *KeywordModel) SearchSome(explains []string) ([]Keyword, error) {
	var kws []Keyword
	result := m.db.Where("explain IN ?", explains).Find(&kws)
	return kws, result.Error
}

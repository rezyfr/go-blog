package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Post struct {
	Guid        string     `gorm:"primary_key;not null;unique" json:"guid"`
	Link        string     `gorm:"not null;unique" json:"link"`
	Title       string     `gorm:"size:255;not null;unique" json:"title"`
	Content     string     `gorm:"not null;" json:"content"`
	Description string     `gorm:"size:255" json:"description"`
	Author      User       `json:"author"`
	Categories  []Category `gorm:"not null"`
	Thumbnail   string     `gorm:"not null" json:"thumbnail"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Guid = html.EscapeString(strings.TrimSpace(p.Guid))
	p.Link = html.EscapeString(strings.TrimSpace(p.Link))
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Author = User{}
	p.Thumbnail = html.EscapeString(strings.TrimSpace(p.Thumbnail))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {
	if p.Guid == "" {
		return errors.New("Required Guid")
	}
	if p.Link == "" {
		return errors.New("Required Link")
	}
	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.Thumbnail == "" {
		return errors.New("Required Thumbnail")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.Guid != "" {
		err = db.Debug().Model(&User{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	var posts []Post
	err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].Author.ID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

func (p *Post) FindPostByID(db *gorm.DB, guid string) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("guid = ?", guid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.Guid == "" {
		err = db.Debug().Model(&User{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) UpdateAPost(db *gorm.DB) (*Post, error) {
	var err error

	err = db.Debug().Model(&Post{}).Where("guid = ?", p.Guid).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Post{}, err
	}
	if p.Guid == "" {
		err = db.Debug().Model(&User{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) DeleteAPost(db *gorm.DB, guid string, uid uint32) (int64, error) {
	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", guid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

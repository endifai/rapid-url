package repositories

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"rapidurl/internal/database"
	"time"
)

var (
	ErrNotFound = errors.New("Not found")
)

type Link struct {
	ShortLink string    `json:"shortLink"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

func (l *Link) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Link) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}

type LinkRepository struct {
	db *database.Database
}

func NewLinkRepository(db *database.Database) *LinkRepository {
	return &LinkRepository{
		db,
	}
}

func (repo *LinkRepository) CreateLink(link string) (*Link, error) {
	hash := sha1.New()
	hash.Write([]byte(link))
	shortLink := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8]

	entity := &Link{
		Link:      link,
		ShortLink: shortLink,
		CreatedAt: time.Now(),
	}

	_, saveError := repo.db.Client.SetNX(database.Ctx, shortLink, entity, time.Hour*24).Result()

	if saveError != nil {
		return nil, saveError
	}

	return entity, nil
}

func (repo *LinkRepository) GetByShortLink(shortLink string) (*Link, error) {
	val, err := repo.db.Client.Get(database.Ctx, shortLink).Result()

	if err != nil {
		return nil, ErrNotFound
	}

	entity := &Link{}

	if err := json.Unmarshal([]byte(val), &entity); err != nil {
		return nil, err
	}

	return entity, nil
}

package repository

import (
	"fmt"
	"my_link_shortener/internal/domain"
	"time"

	"github.com/sqids/sqids-go"
)

var s, _ = sqids.New()

type URLRepository interface {
	DoShort(original string) (*domain.URL, error)
	GetByShort(short string) (*domain.URL, error)
}

type InMemoryUrlRepository struct {
	urls   map[string]*domain.URL
	nextID int
}

func NewInMemoryUrlRepository() *InMemoryUrlRepository {
	return &InMemoryUrlRepository{
		urls:   make(map[string]*domain.URL),
		nextID: 1,
	}
}

func (r *InMemoryUrlRepository) DoShort(original string) (*domain.URL, error) {
	for _, url := range r.urls {
		if url.Original == original {
			return url, nil
		}
	}

	short, err := s.Encode([]uint64{uint64(r.nextID)})
	if err != nil {
		return nil, err
	}
	r.urls[short] = &domain.URL{
		ID:        r.nextID,
		Original:  original,
		Short:     short,
		CreatedAt: time.Now(),
	}
	r.nextID++
	return r.urls[short], nil
}

func (r *InMemoryUrlRepository) GetByShort(short string) (*domain.URL, error) {
	if url, exists := r.urls[short]; exists {
		return url, nil
	}
	return nil, fmt.Errorf("URL not found")
}

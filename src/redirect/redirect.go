package redirect

import "ggf/src/model"

type Store interface {
	GetByShortURL(shortURL string) (*model.Redirect, error)
	GetByOriginalURL(originalURL string) (*model.Redirect, error)
	Create(redirect *model.Redirect) (err error)
	GetAnalytics(shortURL string) (err error)
}

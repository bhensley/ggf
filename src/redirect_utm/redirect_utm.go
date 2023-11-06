package redirect_utm

import "ggf/src/model"

type Store interface {
	GetByShortURL(shortURL string) ([]*model.RedirectUTM, error)
	Create(redirect *model.RedirectUTM) (err error)
}

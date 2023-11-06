package store

import (
	"errors"
	"ggf/src/model"

	"gorm.io/gorm"
)

type RedirectStore struct {
	db *gorm.DB
}

func NewReidrectStore(db *gorm.DB) *RedirectStore {
	return &RedirectStore{db: db}
}

/**
 * @description: GetByShortURL returns a redirect model by its short url
 * @param {string} shortURL
 * @return {*model.Redirect}
 * @return {error}
 */
func (redirectStore *RedirectStore) GetByShortURL(shortURL string) (*model.Redirect, error) {
	var redirectModel model.Redirect

	if err := redirectStore.db.First(&redirectModel, "short_url = ?", shortURL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &redirectModel, nil
}

/**
 * @description: GetByOriginalURL returns a slice of redirect models by original url
 * @param {string} origURL
 * @return {[]*model.Redirect}
 * @return {error}
 * @note: This function is used to check if a redirect already exists- not entirely sure I want to keep it
 */
func (redirectStore *RedirectStore) GetByOriginalURL(origURL string) ([]*model.Redirect, error) {
	var redirectModels []*model.Redirect

	if err := redirectStore.db.Find(&redirectModels, "original_url = ?", origURL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return redirectModels, nil
}

/**
 * @description: Create creates a new redirect model
 * @param {*model.Redirect} redirect
 * @return {error}
 */
func (redirectStore *RedirectStore) Create(redirect *model.Redirect) (err error) {
	return redirectStore.db.Create(redirect).Error
}

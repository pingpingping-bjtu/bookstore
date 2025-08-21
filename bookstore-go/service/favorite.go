package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type FavoriteService struct {
	FavoriteDAO *repository.FavoriteDAO
}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{FavoriteDAO: repository.NewFavoriteDAO()}
}

func (f *FavoriteService) AddFavorite(userID, bookID int) error {
	return f.FavoriteDAO.AddFavorite(userID, bookID)
}

func (f *FavoriteService) RemoveFavorite(userID, bookID int) error {
	return f.FavoriteDAO.RemoveFavorite(userID, bookID)
}

func (f *FavoriteService) GetFavoriteList(userID, page, pageSize int) ([]*model.Favorite, int64, error) {
	favorites, err := f.FavoriteDAO.GetFavoriteList(userID)
	if err != nil {
		return nil, 0, err
	}
	// 简单的分页实现
	total := int64(len(favorites))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(favorites) {
		return []*model.Favorite{}, total, nil
	}

	if end > len(favorites) {
		end = len(favorites)
	}

	return favorites[start:end], total, nil
}

func (f *FavoriteService) CheckFavorite(userID, bookID int) (bool, error) {
	return f.FavoriteDAO.CheckFavorite(userID, bookID)
}

func (f *FavoriteService) GetFavoriteCount(userID int) (int64, error) {
	return f.FavoriteDAO.GetFavoriteCount(userID)
}

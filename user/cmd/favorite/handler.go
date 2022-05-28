package main

import (
	"context"
	"userser/cmd/favorite/kitex_gen/favorite"
)

// FavoritesSerciceImpl implements the last service interface defined in the IDL.
type FavoritesSerciceImpl struct{}

// AddFavorite implements the FavoritesSerciceImpl interface.
func (s *FavoritesSerciceImpl) AddFavorite(ctx context.Context, req *favorite.AddFavoriteRequest) (resp *favorite.AddFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// DelFavorite implements the FavoritesSerciceImpl interface.
func (s *FavoritesSerciceImpl) DelFavorite(ctx context.Context, req *favorite.DelFavoriteRequest) (resp *favorite.DelFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// RenameFavorite implements the FavoritesSerciceImpl interface.
func (s *FavoritesSerciceImpl) RenameFavorite(ctx context.Context, req *favorite.RenameFavoriteRequest) (resp *favorite.RenameFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// AddResult implements the FavoritesSerciceImpl interface.
func (s *FavoritesSerciceImpl) AddResult(ctx context.Context, req *favorite.AddResultRequest) (resp *favorite.AddResultResponse, err error) {
	// TODO: Your code here...
	return
}

// DelResult implements the FavoritesSerciceImpl interface.
func (s *FavoritesSerciceImpl) DelResult(ctx context.Context, req *favorite.DelResultRequest) (resp *favorite.DelResultResponse, err error) {
	// TODO: Your code here...
	return
}

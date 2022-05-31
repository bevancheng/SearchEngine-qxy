package db

import (
	"context"
	"encoding/json"
	"userser/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

type IndexFavorite struct {
	UserName     string `json:"username"`
	FavoriteName string `json:"FavoriteName"`
}

type DocFavorite struct {
	//FavoriteId int64   `json:"favoriteid"`
	IndexId []int64 `json:"indexid"`
}

func IdxFavToBytes(idxF *IndexFavorite) []byte {

	buf, err := json.Marshal(*idxF)
	if err != nil {
		klog.Fatal(err)
		return nil
	}
	return buf
}

func DocFavToBytes(docF *DocFavorite) []byte {

	buf, err := json.Marshal(*docF)
	if err != nil {
		klog.Fatal(err)
		return nil
	}
	return buf
}

func AddFavorite(c *context.Context, idxF *IndexFavorite) error {
	key := IdxFavToBytes(idxF)
	if ldb.Has(key) {
		err := errno.FavAlreadyExistErr
		return err
	}
	emptyDoc := &DocFavorite{IndexId: nil}
	val := DocFavToBytes(emptyDoc)
	err := ldb.Set(key, val)
	if err != nil {
		return err
	}
	return nil
}

func DelFavorite(c *context.Context, idxF *IndexFavorite) error {
	key := IdxFavToBytes(idxF)
	if !ldb.Has(key) {
		err := errno.FavNotExistErr
		return err
	}

	err := ldb.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

func RenameFavorite(c *context.Context, idxF *IndexFavorite, newname string) error {
	key := IdxFavToBytes(idxF)
	if !ldb.Has(key) {
		err := errno.FavNotExistErr
		return err
	}

	val, _ := ldb.Get(key)
	newidx := &IndexFavorite{UserName: idxF.UserName, FavoriteName: newname}
	err := ldb.Delete(key)
	if err != nil {
		return err
	}
	buf := IdxFavToBytes(newidx)
	err = ldb.Set(buf, val)
	if err != nil {
		return err
	}
	return nil
}

func AddResult(c *context.Context, idxF *IndexFavorite, resultid int64) error {
	key := IdxFavToBytes(idxF)
	val, ok := ldb.Get(key)
	if !ok {
		err := errno.FavNotExistErr
		return err
	}
	results := &DocFavorite{}
	err := json.Unmarshal(val, results)
	if err != nil {
		return err
	}

	for _, id := range results.IndexId {
		if id == resultid {
			err := errno.ResAlreadyExistErr
			return err
		}
	}
	results.IndexId = append(results.IndexId, resultid)
	buf, err := json.Marshal(results)
	if err != nil {
		return err
	}
	err = ldb.Set(key, buf)
	if err != nil {
		return err
	}
	return nil
}
func DelResult(c *context.Context, idxF *IndexFavorite, resultid int64) error {
	key := IdxFavToBytes(idxF)
	val, ok := ldb.Get(key)
	if !ok {
		err := errno.FavNotExistErr
		return err
	}
	results := &DocFavorite{}
	err := json.Unmarshal(val, results)
	if err != nil {
		return err
	}

	for i, id := range results.IndexId {
		if id == resultid {
			results.IndexId = append(results.IndexId[:i], results.IndexId[i+1:]...)
			buf, err := json.Marshal(results)
			if err != nil {
				return err
			}
			err = ldb.Set(key, buf)
			if err != nil {
				return err
			}
			return nil
		}
	}
	err = errno.ResNotExistErr
	return err

}

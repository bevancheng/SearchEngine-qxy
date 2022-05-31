package db

import (
	"os"
	"userser/pkg/constants"
)

var ldb *LeveldbStorage

func getFilePath(filename string) string {
	return constants.FavoriteDBPath + string(os.PathSeparator) + filename
}
func Init() {
	s, err := Open(getFilePath(constants.FavoriteTable)) //modify!
	if err != nil {
		panic(err)
	}
	ldb = s
}

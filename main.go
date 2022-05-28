package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"searchengineqxy/searcher"
	"searchengineqxy/searcher/words"
)

type Args struct {
	Addr           string
	DataDir        string
	DictionaryPath string
	GOMAXPROCS     int
	Shared         int
}

func parseArgs() Args {
	var addr = flag.String("addr", "0.0.0.0:5678", "设置监听地址和端口")

	dir := fmt.Sprintf(".%sdata", string(os.PathSeparator))

	var dataDir = flag.String("data", dir, "设置数据存储目录")

	var dictionaryPath = flag.String("dictionary", "./data/dictionary.txt", "设置词典路径")

	var gomaxprocs = flag.Int("gomaxprocs", runtime.NumCPU(), "设置GOMAXPROCS")

	var shared = flag.Int("shared", 10, "文件分片数量")

	flag.Parse()

	return Args{
		Addr:           *addr,
		DataDir:        *dataDir,
		DictionaryPath: *dictionaryPath,
		GOMAXPROCS:     *gomaxprocs,
		Shared:         *shared,
	}
}

func initTokenizer(dictionaryPath string) *words.Tokenizer {
	return words.NewTokenizer(dictionaryPath)
}

func initContainer(args Args, tokenizer *words.Tokenizer) *searcher.Container {
	return nil
}

func main() {
	fmt.Println("hello engine")

}

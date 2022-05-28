package searcher

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"searchengineqxy/searcher/model"
	"searchengineqxy/searcher/storage"
	"searchengineqxy/searcher/utils"
	"searchengineqxy/searcher/words"
	"sync"
	"time"
)

type Engine struct {
	IndexPath string  //索引文件存储目录
	Option    *Option //配置

	invertedIndexStorages []*storage.LeveldbStorage //关键字和Id映射，倒排索引,key=id,value=[]words
	positiveIndexStorages []*storage.LeveldbStorage //ID和key映射，用于计算相关度，一个id 对应多个key，正排索引
	docStorages           []*storage.LeveldbStorage //文档仓

	sync.Mutex                                   //锁
	sync.WaitGroup                               //等待
	addDocumentWorkerChan []chan *model.IndexDoc //添加索引的通道
	IsDebug               bool                   //是否调试模式
	Tokenizer             *words.Tokenizer       //分词器
	DatabaseName          string                 //数据库名

	Shard int //分片数
}

type Option struct {
	InvertedIndexName string //倒排索引
	PositiveIndexName string //正排索引
	DocIndexName      string //文档存储
}

func (e *Engine) Init() {
	e.Add(1)
	defer e.Done()

	if e.Option == nil {
		e.Option = e.GetOptions()
	}
	log.Println("数据存储目录：", e.IndexPath)

	//初始化索引channel
	e.addDocumentWorkerChan = make([]chan *model.IndexDoc, e.Shard)

	for share := 0; share < e.Shard; share++ {
		worker := make(chan *model.IndexDoc, 1000)
		e.addDocumentWorkerChan[share] = worker

		//
		go e.DocumentWorkerExec(worker)

		//初始化docStorage
		s, err := storage.Open(e.getFilePath(fmt.Sprintf("%s_%d", e.Option.DocIndexName, share)))
		if err != nil {
			panic(err)
		}
		e.docStorages = append(e.docStorages, s)

		//初始化invertedIndexStorages
		ks, kerr := storage.Open(e.getFilePath(fmt.Sprintf("%s_%d", e.Option.InvertedIndexName, share)))
		if kerr != nil {
			panic(kerr)
		}
		e.invertedIndexStorages = append(e.invertedIndexStorages, ks)

		//初始化positiveIndexStorages
		idks, ierr := storage.Open(e.getFilePath(fmt.Sprintf("%s_%d", e.Option.PositiveIndexName, share)))
		if ierr != nil {
			panic(ierr)
		}
		e.positiveIndexStorages = append(e.positiveIndexStorages, idks)
	}
	go e.automaticGC()
	log.Println("初始化完成")
}

func (e *Engine) getFilePath(filename string) string {
	return e.IndexPath + string(os.PathSeparator) + filename
}

func (e *Engine) automaticGC() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		<-ticker.C
		runtime.GC()
	}
}

func (e *Engine) InitOption(option *Option) {

	if option == nil {
		//默认值
		option = e.GetOptions()
	}
	e.Option = option
	//shard默认值
	if e.Shard <= 0 {
		e.Shard = 10
	}
	//初始化其他的
	e.Init()

}

func (e *Engine) GetOptions() *Option {
	return &Option{
		DocIndexName:      "docs",
		InvertedIndexName: "inverted_index",
		PositiveIndexName: "postive_index",
	}
}

func (e *Engine) DocumentWorkerExec(worker chan *model.IndexDoc) {
	for {
		doc := <-worker
		e.AddDocument(doc)
	}
}

//获取索引数量
func (e *Engine) GetIndexCount() int64 {
	var count int64
	for i := 0; i < e.Shard; i++ {
		count += e.docStorages[i].Count()
	}
	return count
}

func (e *Engine) GetDocumentCount() int64 {
	var count int64
	for i := 0; i < e.Shard; i++ {
		count += e.docStorages[i].Count()
	}
	return count
}

//分词索引
func (e *Engine) AddDocument(index *model.IndexDoc) {
	e.Wait()
	text := index.Text

	splitWords := e.Tokenizer.Cut(text)

	id := index.Id

	//判断id是否存在，如果存在需要计算差值，并更新
	isUpdate := e.optimizeIndex(id, splitWords)

	if !isUpdate {
		return
	}

	for _, word := range splitWords {
		e.addInvertedIndex(word, id)
	}
	e.addPostiveIndex(index, splitWords)

}

func (e *Engine) optimizeIndex(id uint64, splitWords []string) bool {
	e.Lock()
	defer e.Unlock()

	removes, found := e.getDifference(id, splitWords)
	if found && len(removes) > 0 {
		for _, word := range removes {
			e.removeIdInWordIndex(id, word)
		}
	}

	return !found || len(removes) > 0
}
func (e *Engine) addPostiveIndex(index *model.IndexDoc, keys []string) {
	e.Lock()
	defer e.Unlock()

	key := utils.Uint64ToBytes(index.Id)
	shared := e.getShared(index.Id)
	docStorage := e.docStorages[shared]

	postiveIndexStorage := e.positiveIndexStorages[shared]

	doc := &model.StorageIndexDoc{
		IndexDoc: index,
		Keys:     keys,
	}

	//Id: document
	docStorage.Set(key, utils.Encoder(doc)) //todo Leveldb and Encoder

	//Id: keys(Splited words)
	postiveIndexStorage.Set(key, utils.Encoder(keys))
}

func (e *Engine) addInvertedIndex(word string, id uint64) {
	e.Lock()
	defer e.Unlock()

	shared := e.getSharedByword(word)

	s := e.invertedIndexStorages[shared]

	key := []byte(word)

	buf, find := s.Get(key) //todo
	ids := make([]uint64, 0)
	if find {
		utils.Decoder(buf, &ids)
	}

	//在得到的ids中二分查找，没找到直接append
	if !arrays.BinarySearch(ids, id) {
		ids = append(ids, id)
	}

	s.Set(key, utils.Encoder(ids))
}

func (e *Engine) getSharedByword(word string) int {
	return int(utils.StringToInt(word) % uint32(e.Shard))
}

func (e *Engine) getShared(id uint64) int {
	return int(id % uint64(e.Shard))
}

func (e *Engine) Drop() error {
	return nil
}

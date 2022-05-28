package searcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"searchengineqxy/searcher/words"
)

type Container struct {
	Dir       string
	engines   map[string]*Engine
	Tokenizer *words.Tokenizer
	Debug     bool
	Shared    int
}

func (c *Container) Init() error {
	c.engines = make(map[string]*Engine)

	//读取路径下所有目录，即数据库名称
	dirs, err := ioutil.ReadDir(c.Dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(c.Dir, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			c.engines[dir.Name()] = c.GetDatabase(dir.Name())
		}
	}
	return nil
}

func (c *Container) NewEngine(name string) *Engine {
	var engine = &Engine{
		IndexPath:    fmt.Sprintf("%s%c%s", c.Dir, os.PathSeparator, name),
		DatabaseName: name,
		Tokenizer:    c.Tokenizer,
		Shard:        c.Shared,
	}
	option := engine.GetOptions()

	engine.InitOption(option)
	engine.IsDebug = c.Debug
	return engine
}

func (c *Container) GetDatabase(name string) *Engine {
	if name == "" {
		name = "default"
	}

	log.Println("Get DataBase:", name)
	engine, ok := c.engines[name]
	if !ok {
		engine = c.NewEngine(name)
		c.engines[name] = engine
	}

	return engine
}

func (c *Container) GetDatabases() map[string]*Engine {
	return c.engines
}

func (c *Container) GetDataBaseNumber() int {
	return len(c.engines)
}

func (c *Container) GetIndexCount() int64 {
	var count int64
	for _, engine := range c.engines {
		count += engine.GetIndexCount()
	}
	return count
}

func (c *Container) GetDocumentCount() int64 {
	var count int64
	for _, engine := range c.engines {
		count += engine.GetDocumentCount()
	}
	return count
}

//删除数据库
func (c *Container) DropDataBase(name string) error {
	err := c.engines[name].Drop()
	if err != nil {
		return err
	}
	delete(c.engines, name)
	runtime.GC()
	return nil
}

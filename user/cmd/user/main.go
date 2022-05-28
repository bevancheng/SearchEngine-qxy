package main

import (
	"log"
	"net"
	user "userser/kitex_gen/user/userservice"
	"userser/pkg/constants"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}

	Init()

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

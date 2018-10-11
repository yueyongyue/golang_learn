package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"context"
	"fmt"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	fmt.Println("存储值")
	if _, err := cli.Put(context.Background(), "api", `{api:{name:"k8s"}}`); err != nil {
		fmt.Println("存储值错误：", err)
	}
	fmt.Println("获取值")
	if resp, err := cli.Get(context.Background(), "api"); err != nil {
		fmt.Println("获取值错误：", err)
	} else {
		for _, v := range resp.Kvs {
			fmt.Println(string(v.Value))
		}
	}
	fmt.Println("事务&超时")
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	_, err = cli.Txn(ctx).
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")).
		Then(clientv3.OpPut("key", "XYZ")).
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println("监视")
	rch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

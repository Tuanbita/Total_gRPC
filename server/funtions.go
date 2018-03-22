package main

import (
//	"thriftpool"
	bs "openstars/core/bigset/generic"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
	"fmt"
	"thriftpool"

)

func BigSetClientCreator(host, port string, connTimeout time.Duration, forPool* thriftpool.ThriftPool) (*thriftpool.ThriftSocketClient, error){
	socket, err := thrift.NewTSocketTimeout(fmt.Sprintf("%s:%s", host, port), connTimeout)
	if err != nil {
		return nil, err
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := bs.NewTStringBigSetKVServiceClientFactory(transportFactory.GetTransport(socket), protocolFactory)

	err = client.Transport.Open()
	if err != nil {
		return nil, err
	}

	return &thriftpool.ThriftSocketClient{
		Client: client,
		Socket: socket,
		Parent: forPool,
	}, nil
}

func close(c *thriftpool.ThriftSocketClient) error {
	err := c.Socket.Close()
	//err = c.Client.(*tutorial.PlusServiceClient).Transport.Close()
	return err
}


//GlobalRpcPool = thriftPool.NewThriftPool("10.5.20.3", "23455", 100, 32, 600, Dial, Close)

var (mp=thriftpool.NewMapPool(100, 3600, 3600, BigSetClientCreator, close))
//ham ghi du lieu, vao cong 18407
func testPoolWrite(str string, key string, val string){
	client, _ := mp.Get("127.0.0.1", "18407").Get()
	if (client == nil ) {
		return
	}
	defer client.BackToPool()
	//ghi du lieu vao trong 1 Document la: testBigSet
	client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(str, &bs.TItem{[]byte(str), []byte(val)  } )

	//fmt.Println(mp.Get("127.0.0.1", "18407").GetConnCount(), mp.Get("127.0.0.1", "18407").GetIdleCount() )
}
//18407 la port de minh doc du lieu
func testPoolRead(str string, key string)(val string) {
	client, _ := mp.Get("127.0.0.1", "18407").Get()
	if (client == nil ) {
		return;
	}
	defer client.BackToPool()
	//co the thay the testBigSet = UserInfo cung duoc
	//luc nay phai them cac truong vao trong UserInfo: ID, Name, Age, Address

	//Doc Du Lieu Tu Document testBigSet

	res, _ := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(str, []byte(key))
	val = string (res.Item.Value[:])
	if (res != nil && res.Item != nil && res.Item.Value != nil) {
		return val
	}
	return ""
}

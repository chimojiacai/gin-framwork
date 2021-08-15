// @Author: yongzhen5
// @Description: nacos 配种中心
// @File: nacos
// @Date: 2021/8/15 21:18

package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewNacosConfig() string {
	clientConfig := constant.ClientConfig{
		TimeoutMs:            10 * 1000, //http请求超时时间，单位毫秒
		ListenInterval:       30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
		BeatInterval:         5 * 1000,  //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
		NamespaceId:          "",        //nacos命名空间
		Endpoint:             "",        //获取nacos节点ip的服务地址
		CacheDir:             "",        //缓存目录
		LogDir:               "",        //日志目录
		UpdateThreadNum:      20,        //更新服务的线程数
		NotLoadCacheAtStart:  true,      //在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,      //当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	// 至少一个(集群可以多个)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
		return ""
	}

	//success, err := configClient.PublishConfig(vo.ConfigParam{
	//	DataId:  "testId",
	//	Group:   "testGroup",
	//	Content: "hello world!222222"})
	//if err!=nil ||!success {
	//	fmt.Println("publish failed",err)
	//}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "ceshi",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
		return ""
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "ceshi",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil || content == "" {
		panic(err)
		return ""
	}
	fmt.Println("content:" + content)
	return content
}

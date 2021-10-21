# Redis设计原理
https://www.cnblogs.com/orange-CC/p/12517201.html


## 高可用

```
  Sentinel
  
    选举模式
            Sentinel哨兵模式  
            如果以前的主服务器再次上线时，自动变成从服务器，
            将从新的主服务器复制数据保持同步
    健康检查
            每秒钟发送一次  判断当前节点是否存在
  
    选举协议
           https://zhuanlan.zhihu.com/p/27207160
           raft协议 一致性分布式协议
```

## 分布式

```
Redis Cluster

    每个主节点还可以再加入自己的从节点，主节点兼有Sentinel特性
    集群即主主之间通过分片（sharding）来分配和共享数据
    集群的整个数据库被分为16384个槽（slot）
    集群的各个节点通过Gossip协议来交换各自信息
        一个节点状态发生变化，并向临近节点发送更新副本信息
        对于节点状态变化的信息随机发送给b个节点
        随着时间推移，信息能够传达到所有的节点
```




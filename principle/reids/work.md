# Redis工作原理

## 工作方式

1. 多样的数据模型 
   ``` 
   key-value string hash list set sort_set HyperLogLog 订阅 事务
   ```
2. 持久化
   ```
        ROF 每次操作都卸乳数据库
        RDB 类似于每小时归档您的RDB文件   
   ```
3. 主从同步
   ``` 
        redis 哨兵模式
   ```

## 分布式可扩展性

   ```
      后来Redis Cluster是一个实现了分布式且允许单点故障的Redis高级版本
      它没有中心节点，各个节点地位一致，具有线性可伸缩的功能。
      Redis Cluster的分布式存储结构，其中节点与节点之间通过二进制协议进行
      通讯，节点与客户端之间通过ascii协议进行通信， 在数据的放置策略上，
      Redis Cluster将整个key的数值域分成16384个哈希槽。
      每个节点上可以存储一个或多个哈希槽，
      也就是说当前Redis Cluster支持的最大节点就是16384
   ```


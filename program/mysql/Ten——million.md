# 一次MySQL两千万数据大表的优化过程

## 方案一：优化现有mysql数据库。

1. 优点：不影响现有业务，源程序不需要修改代码，成本最低。
2. 缺点：有优化瓶颈，数据量过亿就玩完了。
3. 数据库设计和表创建时就要考虑性能
4. sql的编写需要注意优化
5. 分区
6. 分表
7. 分库

## 方案二：升级数据库类型，换一种100%兼容mysql的数据库。

1. 优点：不影响现有业务，源程序不需要修改代码，你几乎不需要做任何操作就能提升数据库性能，
2. 缺点：多花钱
3. tiDB https://github.com/pingcap/tidb
4. Cubrid https://www.cubrid.org/
5. 阿里云POLARDB https://www.aliyun.com/product/polardb
6. 阿里云OcenanBase
7. 阿里云HybridDB for MySQL
8. 腾讯云DCDB https://cloud.tencent.com/product/dcdb_for_tdsql

## 方案三：一步到位，大数据解决方案，更换newsql/nosql数据库。
1. 优点：扩展性强，成本低，没有数据容量瓶颈
2. 缺点：需要修改源程序代码
3. hadoop家族。hbase/hive怼上就是了
4. 阿里云的MaxCompute配合DataWorks

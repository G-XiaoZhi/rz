## 这里记录整体进展、思考、过程

### 2021-08-20
- 在今天之前已经整理好arch、judge文件
- 今天针对代码的文件夹框架进行搭建

- 下一步考虑：
  - 考虑对于db读取性能的验证
  - 对于yaml解析过程的验证
  - 对于yaml解析后judge逻辑的处理
  - 对于core内容的完善
  - 对于models、dao的完善
  - 对于日志、权限等的处理
  
  
### 20210822
- 测试数据库读写性能
    - go test -v -benchmem -bench="BenchmarkGetDailyQfqByInterval"
    - 之前只加了uniq（ts_code、trade_date），所以基于trade_date查找的时候是没有索引的，效率比较低, 其实后续考虑给trade_date加索引，
    因为考虑按日期去db读取数据，然后加载到内存中再按ts_code去处理
    > ALTER  TABLE  `daily_qfq`  ADD  KEY `idx_trade_date` (`trade_date`);
    - 测试结果如下：一年的数据大概2G，一次读取23s左右
    > [46132.928ms] [rows:1822155] SELECT * FROM `daily_qfq` WHERE trade_date >= '20190101' and trade_date <= '20210101'
    - 数据采用select后内存和效率提升不少
    > [16724.554ms] [rows:1822155] SELECT `ts_code`,`trade_date`,`open`,`close`,`ma5` FROM `daily_qfq` WHERE trade_date >= '20190101' and trade_date <= '20210101'
      2018413
      BenchmarkGetDailyQfqByInterval-4   	       1	16724666664 ns/op	1281178888 B/op	49198890 allocs/op
      PASS
      ok  	rz/dao	16.847s

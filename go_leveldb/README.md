
- [sstable](#sstable)
  - [data block结构](#data-block结构)
  - [filter block结构](#filter-block结构)
  - [读写操作](#读写操作)
- [Reference](#reference)

# sstable

## data block结构

* key值前缀压缩
    * Restart point: 每间隔若干个keyvalue对，将为该条记录重新存储一个完整的key => 前缀索引
    * 默认间隔值为16: 过大/过小 都有问题，权衡？

```
第一部分用来存储keyvalue数据。由于sstable中所有的keyvalue对都是严格按序存储的，为了节省存储空间，leveldb并不会为每一对keyvalue对都存储完整的key值，而是存储与上一个key非共享的部分，避免了key重复内容的存储。
```

## filter block结构

* 布隆过滤器
  * 加快sstable中数据查询的效率
  * 若判断不存在，则无需对这个datablock进行数据查找
* 其索引：Meta Index Block

## 读写操作

* 对sstable进行写操作的数据结构为tWriter

```go
// tWriter wraps the table writer. It keep track of file descriptor
// and added key range.
type tWriter struct {
	t *tOps

	fd storage.FileDesc
	w  storage.Writer
	tw *table.Writer

	first, last []byte
}
```

# Reference

* leveldb-handbook https://leveldb-handbook.readthedocs.io/zh/latest/index.html



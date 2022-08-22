---
null
---
# 同步topic

## commodity-srv: spu

### Topic:

```go
commodity.internal.spu.sync.create

commodity.internal.spu.sync.update

commodity.internal.spu.sync.issue   // 发布中心发布商品

commodity.internal.spu.sync.cancel  // 发布中心取消发布商品

commodity.internal.spu.sync.audit   // 审核通过

commodity.internal.spu.sync.reject  // 驳回

commodity.internal.spu.sync.delete  // 删除

commodity.internal.spu.sync.move    // 迁移
```

### Msg:

```go
type SpuSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Action SpuSyncInfo_Action `protobuf:"varint,2,opt,name=action,proto3,enum=spu.SpuSyncInfo_Action" json:"action"`
}
```

## thing-srv: demand  tag: [micro4/v0.0.3-dev](https://git.uino.com/thingyouwe-public-proto/thing-pb/-/tags/micro4%2Fv0.0.3-dev)

### Topic:

```go
thing.internal.demand.sync.delete  // 删除
thing.internal.demand.sync.cancel  // 取消发布
thing.internal.demand.sync.audit  // 审核通过
thing.internal.demand.sync.reject  // 驳回
```

### Msg:

```go
type DemandSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Action DemandSyncInfo_Action `protobuf:"varint,2,opt,name=action,proto3,enum=demand.DemandSyncInfo_Action" json:"action"`
}
```

## thing-srv: article  tag: [micro4/v0.0.3-dev](https://git.uino.com/thingyouwe-public-proto/thing-pb/-/tags/micro4%2Fv0.0.3-dev)

### Topic:

```go
thing.internal.article.sync.cancel // 取消发布
thing.internal.article.sync.delete // 删除文章
thing.internal.article.sync.issue // 发布文章
```

### Msg:

```go
type ArticleSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Action ArticleSyncInfo_Action `protobuf:"varint,2,opt,name=action,proto3,enum=article.ArticleSyncInfo_Action" json:"action"`
}

```

## tender-srv tag: [proto/micro4/v0.0.1-dev](https://git.uino.com/thingyouwe-service-app/tender-srv/-/tags/proto%2Fmicro4%2Fv0.0.1-dev)

### Topic:

```go
tender.internal.tender.sync.issue //发布
tender.internal.tender.sync.cancel //取消发布
tender.internal.tender.sync.delete //删除
```

### Msg:

```go
type TenderSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Action TenderSyncInfo_Action `protobuf:"varint,2,opt,name=action,proto3,enum=tender.TenderSyncInfo_Action" json:"action"`
}
```

## metadata-srv tag: [micro4/v0.0.1-dev](https://git.uino.com/thingyouwe-public-proto/metadata-srv-pb/-/tags/micro4%2Fv0.0.1-dev)

### Topic:

```go
metadata.internal.tag.sync.create //创建
metadata.internal.tag.sync.update //更新
metadata.internal.tag.sync.update_alias //更新别名
metadata.internal.tag.sync.delete //删除
```

### Msg:

```go
type TagSyncInfo struct {
   state         protoimpl.MessageState
   sizeCache     protoimpl.SizeCache
   unknownFields protoimpl.UnknownFields

   Id     string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
   Action TagSyncInfo_Action `protobuf:"varint,2,opt,name=action,proto3,enum=tag.TagSyncInfo_Action" json:"action"`
}
```

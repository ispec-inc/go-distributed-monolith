# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [admin/view/article.proto](#admin/view/article.proto)
    - [Article](#admin.view.Article)
  
- [admin/api/rest/v1/article/article.proto](#admin/api/rest/v1/article/article.proto)
    - [CreateRequest](#admin.api.article.CreateRequest)
    - [CreateResponse](#admin.api.article.CreateResponse)
    - [DeleteRequest](#admin.api.article.DeleteRequest)
    - [DeleteResponse](#admin.api.article.DeleteResponse)
    - [GetRequest](#admin.api.article.GetRequest)
    - [GetResponse](#admin.api.article.GetResponse)
    - [ListRequest](#admin.api.article.ListRequest)
    - [ListResponse](#admin.api.article.ListResponse)
    - [UpdateRequest](#admin.api.article.UpdateRequest)
    - [UpdateResponse](#admin.api.article.UpdateResponse)
  
- [article/view/user.proto](#article/view/user.proto)
    - [User](#article.view.User)
  
- [article/api/rest/v1/user/user.proto](#article/api/rest/v1/user/user.proto)
    - [CreateRequest](#article.api.user.CreateRequest)
    - [CreateResponse](#article.api.user.CreateResponse)
    - [GetRequest](#article.api.user.GetRequest)
    - [GetResponse](#article.api.user.GetResponse)
  
- [media/api/rest/v1/image/image.proto](#media/api/rest/v1/image/image.proto)
    - [CreateRequest](#media.api.image.CreateRequest)
    - [CreateResponse](#media.api.image.CreateResponse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="admin/view/article.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## admin/view/article.proto



<a name="admin.view.Article"></a>

### Article



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| body | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="admin/api/rest/v1/article/article.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## admin/api/rest/v1/article/article.proto



<a name="admin.api.article.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| body | [string](#string) |  |  |






<a name="admin.api.article.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| articles | [admin.view.Article](#admin.view.Article) | repeated |  |






<a name="admin.api.article.DeleteRequest"></a>

### DeleteRequest







<a name="admin.api.article.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| articles | [admin.view.Article](#admin.view.Article) | repeated |  |






<a name="admin.api.article.GetRequest"></a>

### GetRequest







<a name="admin.api.article.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article | [admin.view.Article](#admin.view.Article) |  |  |






<a name="admin.api.article.ListRequest"></a>

### ListRequest







<a name="admin.api.article.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| articles | [admin.view.Article](#admin.view.Article) | repeated |  |






<a name="admin.api.article.UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| body | [string](#string) |  |  |






<a name="admin.api.article.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article | [admin.view.Article](#admin.view.Article) |  |  |





 

 

 

 



<a name="article/view/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## article/view/user.proto



<a name="article.view.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| email | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="article/api/rest/v1/user/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## article/api/rest/v1/user/user.proto



<a name="article.api.user.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="article.api.user.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [article.view.User](#article.view.User) | repeated |  |






<a name="article.api.user.GetRequest"></a>

### GetRequest







<a name="article.api.user.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [article.view.User](#article.view.User) |  |  |





 

 

 

 



<a name="media/api/rest/v1/image/image.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## media/api/rest/v1/image/image.proto



<a name="media.api.image.CreateRequest"></a>

### CreateRequest
Request should be `multipart/form-data`, not json.






<a name="media.api.image.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | The place stored requested image. |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |


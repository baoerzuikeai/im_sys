### user_basic
```json
{
    "_id": ObjectId("6688f4cf2f5a0000d60066a2"),
    "user_identity": "自定义用户uid",
    "account": "账号",
    "password": "密码",
    "nickname": "昵称",
    "sex": 1,
    "email": "邮箱",
    "avatar": "头像",
    "created_at": 1,
    "updated_at": 1
}
```
### user_room
```json
{
    "_id": ObjectId("6688f5e72f5a0000d60066a5"),
    "user_identity": "用户的唯一标识",
    "room_identity": ObjectId("000000000000000000000000"),
    "created_at": 1,
    "updated_at": 1
}

```

### room_basic
```json
{
    "_id": ObjectId("6688f5a22f5a0000d60066a4"),
    "number": "房间号",
    "name": "房间名称",
    "info": "房间简介",
    "user_identity": "用户唯一标识",
    "created_at": 1,
    "updated_at": 1
}

```
### private_message_basic

```json
{
    "_id": ObjectId("669f6c83bc5600005e00365a"),
    "user_identity": "用户的唯一标识",
    "receive_user_identity": "接受方用户唯一标识",
    "data": "发送的数据",
    "created_at": 1,
    "updated_at": 1
}

```
### public_message_basic
```json
{
    "_id": ObjectId("6688f5182f5a0000d60066a3"),
    "user_identity": "用户的唯一标识",
    "room_identity": ObjectId("000000000000000000000000"),
    "data": "发送的数据",
    "created_at": 1,
    "updated_at": 1
}
```

### TestMdoel
- User1 
    - account(2634174807)
- User2 
    - account(486642549) 
- User1 
    - account(baoerzuikeai)


- Roomid()
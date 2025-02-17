# IM System - 基于WebSocket的即时通讯系统

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-1.9%2B-blue.svg)](https://gin-gonic.com/)
[![MongoDB](https://img.shields.io/badge/MongoDB-6.0%2B-green.svg)](https://www.mongodb.com/)

## 项目简介

基于Golang的高性能即时通讯系统，采用WebSocket协议实现实时消息推送，支持高并发场景。结合MongoDB进行数据持久化，Redis实现验证码缓存，提供安全可靠的即时通讯解决方案。

## 功能特性

✅ **用户管理**  
- 用户注册/登录（邮件验证码）
- JWT身份验证
- 密码找回（邮件验证码）

✅ **即时通讯**  
- WebSocket长连接
- 单聊/群聊支持
- 消息实时推送
- 消息已读状态

✅ **数据管理**  
- 聊天记录存储（MongoDB）
- 历史消息查询
- 消息分页加载

✅ **安全体系**  
- 验证码机制（Redis缓存）
- API访问权限控制
  
✅ **高性能设计**  
- Goroutine连接池
- MongoDB查询优化
- Redis缓存热点数据

## 技术栈

| 技术组件       | 说明                              |
|----------------|----------------------------------|
| **Golang**     | 后端核心开发语言                  |
| **Gin**        | RESTful API框架                  |
| **WebSocket**  | 实时双向通信协议                  |
| **MongoDB**    | 主数据库（存储用户/消息数据）     |
| **Redis**      | 验证码缓存/会话状态管理           |
| **JWT**        | 安全的身份验证机制                |

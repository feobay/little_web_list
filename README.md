# Gin web



本项目来源于 

[基于gin框架和gorm的web开发实战 (七米出品)](https://www.bilibili.com/video/BV1gJ411p7xC)

重构了连接数据库部分。



运行方式

- 将代码下载到本地
- 运行docker设置好数据库配置（配置见 `configs/config.yml` 文件）
- 将代码下载到本地，并在Goland新建项目



疑惑点解答

- 任务完成状态变化由前端完成
  - 前端点击完成，则自动将状态取反，并向后端发送PUT请求更新数据库
- 后端数据库部分连接的是dokcer容器，容器内部mysql服务为3306端口，映射到主机中3307端口








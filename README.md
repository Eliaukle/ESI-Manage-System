# ESI Manage System

​		本项目为一个数据库课程设计——ESI高校管理系统，包括大学信息、大学排名、学科信息、学科排名、论文数据、用户信息的增删改查，通过对上述数据的管理评估高校实力与教学水平

## 1. 环境要求

在开始之前，请保证您安装好以下环境（本项目所使用的版本仅供参考）：

|         | 版本查看语句      | 本项目使用的版本 |
| ------- | ----------------- | ---------------- |
| Go      | `go version`      | 1.17.6           |
| Node.js | `node --version`  | 16.14.0          |
| Vue.js  | `vue --version`   | 2.9.6            |
| MySQL   | `mysql --version` | 8.0.28           |

## 2. 数据库准备

在mysql上新建数据库esi_db并在数据库中创建一个初始用户

```sql
INSERT INTO esi_db.users(user_name, password, identify) VALUES ('admin', '$2a$10$NtFIBoDruaKkNsAupMryQ.532STnLj9kERd8bzoGaZq5N9dhYY5R.', '系统管理员')
```

![](https://s3.uuu.ovh/imgs/2022/12/05/8e70cae9e0b8eaf6.png)

打开esi_server/common/database.go，在InitDB函数中修改mysql用户名（root）及密码（root password）：

![](https://s3.uuu.ovh/imgs/2022/12/05/5fa5aced28b748d4.png)

## 3. 启动项目

从终端进入esi_server，输入以下语句启动后端：

```
go run main.go
```

从前端进入esi_client，输入以下语句启动前端：

```
npm run dev
```


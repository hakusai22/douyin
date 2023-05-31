# douyin

基于 MySql + Gorm + Gin HTTP服务完成的第三届字节跳动青训营--极简抖音后端项目
<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
![GitHub top language](https://img.shields.io/github/languages/top/hakusai22/douyin?style=for-the-badge)

<!-- PROJECT LOGO -->
<br />



<p align="center">
    <a href="https://github.com/hakusai22/douyin/">
    <img src="https://fastly.jsdelivr.net/gh/hakusai22/image/qq.jpg" alt="Logo" width="200" height="200">
    </a>
    <h3 align="center">字节青训营抖音项目</h3>
  <p align="center">
    ·
    <a href="https://github.com/hakusai22/douyin/issues">报告Bug</a>
    ·
    <a href="https://github.com/hakusai22/douyin/issues">提出新特性</a>
  </p>

<!-- links -->
[your-project-path]:hakusai22/douyin
[contributors-shield]: https://img.shields.io/github/contributors/hakusai22/douyin.svg?style=for-the-badge
[contributors-url]: https://github.com/hakusai22/douyin/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/hakusai22/douyin.svg?style=for-the-badge
[forks-url]: https://github.com/hakusai22/douyin/network/members
[stars-shield]: https://img.shields.io/github/stars/hakusai22/douyin.svg?style=for-the-badge
[stars-url]: https://github.com/hakusai22/douyin/stargazers
[issues-shield]: https://img.shields.io/github/issues/hakusai22/douyin.svg?style=for-the-badge
[issues-url]: https://img.shields.io/github/issues/hakusai22/douyin.svg
[license-shield]: https://img.shields.io/github/license/hakusai22/douyin.svg?style=for-the-badge
[license-url]: https://github.com/hakusai22/douyin/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/xxxx

![img.png](https://fastly.jsdelivr.net/gh/hakusai22/douyin/image/img3.png)


## 目录

- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
    - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [项目特点](#项目特点)
- [功能介绍](#功能介绍)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)
- [成果演示](#成果演示)


### 上手指南

###### 开发前的配置要求

###### **安装步骤**


### 项目特点

1. 基于《[接口文档在线分享](https://www.apifox.cn/apidoc/shared-7f20ed46-edeb-4dff-a35d-5b899855b8bf)[- Apifox](https://www.apifox.cn/apidoc/shared-7f20ed46-edeb-4dff-a35d-5b899855b8bf)》提供的接口进行开发，使用《[极简抖音](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)[App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行Demo测试， **功能完整实现** ，前端接口匹配良好。

2. 代码结构采用 (HTTP API 层 + RPC Service 层+Dao 层) 项目 **结构清晰** ，代码 **符合规范**

3. 使用 **JWT** 进行用户token的校验

4. 使用 **Gorm** 对 MySQL 进行 ORM 操作；

5. 数据库表建立了索引和外键约束，对于具有关联性的操作一旦出错立刻回滚，保证数据一致性和安全性

### 功能介绍

- 视频：视频推送、视频投稿、发布列表
- 用户：用户注册、用户登录、用户信息
- 点赞：点赞操作、点赞列表
- 评论：评论操作、评论列表
- 关注：关注操作、关注列表、粉丝列表

```sh
git clone https://github.com/haksuai22/douyin.git
go mod tidy
./main.go 
```

### 文件目录说明

eg:

```shell
filetree
├─douyin（公共基础库，封装一些通用的逻辑）
│  ├─redis_cache（缓存）
│  ├─config（配置文件）
│  ├─controller（于handlers层）
│  ├─middlewares（中间件）
│  ├─models（实体类）
│  └─router（接口path访问路口）
│  └─service（操作数据库）
│  └─static（存储视频图片静态目录）
│  └─util（工具类）
│  └─image（图床）
```

### 开发的架构

![img.png](https://fastly.jsdelivr.net/gh/hakusai22/douyin/image/framework.png)

本项目采用 MVC 分层设计模型分离模型层、视图层和控制层，从而降低代码的耦合度，提高项目的可维护性。

使用 Gin 作为 Web 框架，Redis 作为缓存框架，MySQL 作为持久层框架。
```
         ┌─────────┐      ┌─────────┐      ┌─────────┐
  ──req──►         ├──────►         ├──────►         │
         │   Gin   │      │  Redis  │      │  MySQL  │
  ◄─resp─┤         ◄──────┤         ◄──────┤         │
```

### 部署

### 使用到的框架
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- [MySQL](https://www.mysql.com/)
- [Gorm](https://gorm.io/)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [ffmpeg](https://github.com/FFmpeg/FFmpeg)
- [toml](https://github.com/BurntSushi/toml)
- ......
### 贡献者

请阅读 **CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在 repository 参看当前可用版本。



### 作者

hakusai22@qq.com

博客:[Hakusai](https://hakusai.cn)  

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/mrxuexi/tiktok/LICENSE.txt)

### 鸣谢


### 成果演示

![img.png](https://fastly.jsdelivr.net/gh/hakusai22/douyin/image/img1.png)

![img.png](https://fastly.jsdelivr.net/gh/hakusai22/douyin/image/img2.png)
![img.png](https://fastly.jsdelivr.net/gh/hakusai22/douyin/image/img3.png)

## star 趋势图

![Stargazers over time](https://starchart.cc/hakusai22/douyin.svg)

## 贡献者
<a href="https://github.com/hakusai22/douyin/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=hakusai22/douyin" />
</a>
# douyin

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
│  ├─cache（缓存）
│  ├─config（配置文件）
│  ├─handlers（于handlers层）
│  ├─middleware（中间件）
│  ├─models（实体类）
│  └─router（接口path访问路口）
│  └─service（操作数据库）
│  └─static（存储视频图片静态目录）
│  └─util（工具类）
```

### 开发的架构

### 部署

### 使用到的技术

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- [MySQL](https://www.mysql.com/)
- - ......
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
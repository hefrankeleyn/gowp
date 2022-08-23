# 使用Go和Gin开发一个RESTFulAPI

[toc]

## 一、说明

这篇教程介绍基础的使用Go和Gin Web服务框架写一个RESTful web服务API。

如果您对 Go 及其工具有基本的了解，您将充分利用本教程。

Gin简化了很多与构建web应用，包括web服务的代码任务。在这个教程中，你将使用Gin去路由请求，获取请求详情，并编组 JSON 以获取响应。

在这个教程中，你将使用两个断点构建RESTful API。您的示例项目将是有关老式爵士乐唱片的数据存储库。

这个教程包含下面部分：

- 设计API端点；
- 为你的代码创建一个文件夹；
- 创建数据；
- 写一个处理器去返回所有的项；
- 写一个处理器去添加新的项目；
- 写一个处理器去返回一个特定的项目；

要尝试将此作为您在 Google Cloud Shell 中完成的交互式教程，请单击下面的按钮。

## 二、准备

- 安装了Go1.16 及以上版本；
- 有用于编写代码的工具；
- 一个命令行工具；
- 一个`curl`工具；

## 三、设计一个API端点

您将构建一个 API，该 API 可让您访问销售黑胶唱片的商店。因此，您需要提供端点，客户端可以通过这些端点为用户获取和添加相册。

当开发一个API，您通常从设计端点开始。如果端点易于理解，您的 API 用户将获得更大的成功。

这儿是你将在本教程中创建的断点：

- `/albums`

  GET请求，获取所有的专辑，作为JSON返回；

  POST请求，根据JSON格式的请求数据，添加一个转接；

- `albums/:id`

  GET请求，根据它的ID获取一个专辑，以JSON格式返回请求数据。

接下来，你将为你的代码创建一个目录。

### 3.1 为你的代码创建一个目录

开始，为你将写的代码创建一个项目：

1. 打开命令行，进入主目录：

   ```shell
   $ cd 
   $
   ```

2. 使用命令行，创建一个叫`web-service-gin`的目录

   ```shell
   $ mkdir web-service-gin
   $ cd web-service-gin/
   $
   ```

3. 创建一个模块，你能管理管理依赖

   运行`go mod init` 命令，给出你的代码将要在的模块路径。

   ```shell
   $ go mod init example/web-service-gin
   go: creating new go.mod: module example/web-service-gin
   $
   ```

   这个命令创建了一个`go.mod`文件，在这个文件中，你能添加跟踪列表。

接下来，我们将设计数据结构，用于处理数据。

### 3.2 创建数据

为了保持本教程的东西简单，你将存储数据到内存中。一个非常典型的API将使用数据库进行交互。

注意，存储数据在内存中，意味着在每次停止服务的时候专辑集合会丢失，当你启动它的时候会重新创建它。

#### （1）写代码

1. 使用编辑器，在`web-service-gin` 目录中创建一个`main.go`文件。你将在这个文件中写代码。

2. 进入`main.go` 文件，粘贴下面的一句声明：
   ```go
   package main
   ```

   一个独立的程序（与库相反）总是在`main` 包中。

3. 包下方的声明，粘贴下面的声明作为`album`结构。你将使用它在内存中存储专辑数据。

   结构的标签像：`json:"artist"`指明了当结构的内容被序列化成json的时候，字段的名称是什么。没有它们，JSON将使用结构中的大写字段名称——一种在 JSON 中不常见的样式。

   ```go
   // album 代表一个专辑的数据
   type album struct {
     ID string `json:"id"`
     Title string `json:"title"`
     Artist string `json:"artist"`
     Price  float64 `json:"price"`
   }
   ```

4. 仅在结构下面添加，粘贴下面的专辑切片结构，包含了你将开始使用的数据

   ```go
   var albums = []album {
       {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
       {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
       {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
   }
   ```

接下来，你将写代码实现你第一个断点。


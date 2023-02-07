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

接下来，你将写代码实现你第一个端点。

### 3.3  返回所有的项目

当客户端进行一个`/albums`的GET请求时，你以JSON的形式返回所有的专辑。

为了做这个，要进行下面两步：

- 准备响应逻辑
- 编写请求路径到逻辑的映射；

注意，这与它们在运行时时相反的。但是你首先要添加依赖，并基于依赖进行编码。

1. 在上面部分，你添加下面的代码结构，粘贴下面的代码用于获取专辑列表。

   这个`getAlbums`函数根据切片列表创建一个JSON，写这个JSON到响应中。

   ```go
   // getAlbums 使用专辑列表作为JSON响应
   func getAlbums(c *gin.Context) {
     c.IndentedJSON(http.StatusOK, albums)
   }
   ```

   在这个代码中：

   - 写了一个`getAlbums` 函数，并获取一个`gin.Context`参数。注意，你能给这个函数任意到名字，Gin或Go都不需要特定的函数名格式。

     `gin.Context`是Gin非常重要的部分。它携带了请求详情，验证和序列化JSON，等等。（尽管名字相似，但是它不同于Go的内置上下文包）。

   - 调用`c.IndextedJSON`去序列化结构成为JSON，并且添加到响应中。

     函数的第一个参数你想发送给客户端的HTTP状态码。这里，你从`net/http`包，传递`StatusOK`状态，表明是`200 OK`。

     注意，你能替换`Context.IndentedJSON`使用`Context.JSON`，去发送一个更严谨的JSON。在实际中，缩进形式在调试时更容易使用，并且大小差异通常很小。

2. 在靠近`main.go`的顶部，在`albums`切片声明下面，粘贴下面的代码，分配处理函数到端点路径。

   这建立起了`getAlbums`处理请求到 `/albums`端点路径到联系。

   ```go
   func main() {
     router := gin.Default()
     router.GET("/albums", getAlbums)
     router.Run("localhost:8080")
   }
   ```

   在这个代码中：

   - 使用`Deault`初始化Gin 路由；

   - 使用`GET`函数去关联`GET`到HTTP方法和`/albums`路径，带有一个处理函数；

     注意，你传递`getAlbums`函数名。这是不同于传递函数结果，你可以通过传递`getAlbums()`。

   - 使用`Run`函数去绑定路由到`http.Server`，并启动服务。

3. 在`main.go`函数的顶部，下面的包声明，导入包，你将需要支持你已经写的代码。

   第一行代码看起来像这样：

   ```go
   import (
     "net/http"
     "github.com/gin-gonic/gin"
   )
   ```

4. 保存`main.go`文件。

### 3.4 运行代码

1. 从跟踪Gin模块，作为依赖

   在命令行，为你的模块使用`go get`去添加`github.com/gin-gonic/gin`模块作为依赖。使用一个点参数，意味着要为当前的目录代码获取依赖。

   ```shell
   $ go get .
   go: downloading github.com/gin-gonic/gin v1.8.1
   $
   ```

   Go解决并下来依赖，去满足你在之前步骤中添加的导入声明。

2. 在包含`main.go`文件的命令行中，运行代码。使用一个点参数，意思是运行当前目录的代码。

   ```shell
   $ go run .
   [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
   
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:	export GIN_MODE=release
    - using code:	gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
   [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
   Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
   [GIN-debug] Listening and serving HTTP on localhost:8080
   
   ```

   一旦代码运行，你就开始了一个HTTP服务，你能用来发送请求。

3. 在一个新的命令行窗口，使用`curl`去制造一个请求，用于运行你的web服务。

   该命令应显示您为服务播种的数据。

   ```shell
   $ curl http://localhost:8080/albums
   [
       {
           "id": "1",
           "title": "Blue Train",
           "artist": "John Coltrane",
           "price": 56.99
       },
       {
           "id": "2",
           "title": "Jeru",
           "artist": "Gerry Mulligan",
           "price": 17.99
       },
       {
           "id": "3",
           "title": "Sarah Vaughan and Clifford Brown",
           "artist": "Sarah Vaughan",
           "price": 39.99
       }
   ]$
   ```

   你已经开始了一个API。在下面的部分，你将使用代码创建另一个端点，用于处理`POST`请求来添加一个项目。

## 四、写一个端点用于添加一个新项目

当客户创造一个`POST`的`/albums`请求，你想添加在请求体中描述的album到已经存在的albums的数据中。

为了做这件事，你要做下面的事情：

- 编写添加一个新album到存在列表的逻辑；
- 一点代码用于路由`POST`请求到你的逻辑；

### 4.1 写代码

1. 添加代码用于添加album数据到album列表中；

   在`import` 语句后面到某处，粘贴下面的代码。（文件的结尾是写这些代码的好地方，但是Go不强迫你声明函数代码的顺序）

   ```go
   // postAlbums 添加一个album 从接受的JSON请求体中
   func postAlbums (c *gin.Context) {
     var newAlbum album
     // 调用BindJSON去绑定接受的JSON，变成一个newAlbum
     if err := c.BindJSON(&newAlbum); err != nil {
       return
     }
     
     // 添加新的album到切片中
     albums = append(albums, newAlbum)
     c.IndentedJSON(http.StatusCreated, newAlbum)
   }
   ```

   在这个代码中：

   - 使用`Context.BindJSON`去绑定请求体成为newAlbum；
   - 添加从JSON中初始的`album`结构到albums切片；
   - 添加`201`状态码去响应，同时使用JSON代表你添加的album；

2. 改变`main`函数，以便于它包含`router.POST` 函数，像下面：

   ```go
   func main() {
   	router := gin.Default()
   	router.GET("/albums", getAlbums)
     router.POST("/albums", postAlbums)
   	router.Run("localhost:8080")
   }
   ```

   在这个代码中：

   - 关联`POST`方法到`/albums`路径，带有`postAlbums`函数；

     使用`Gin`，你能关联一个带有“方法路径”组合的处理器。在这种方法中，你能分别路由请求，发送到单个路径基于客户端使用的方法。

### 4.2 运行代码

1. 如果从上一步，服务仍然在运行，停止它。

2. 在`main.go`文件的目录下的命令行中，运行下面的代码：

   ```go
   $ go run .
   [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
   
   [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:	export GIN_MODE=release
    - using code:	gin.SetMode(gin.ReleaseMode)
   
   [GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
   [GIN-debug] POST   /albums                   --> main.postAlbums (3 handlers)
   [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
   Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
   [GIN-debug] Listening and serving HTTP on localhost:8080
   
   ```

3. 在另一个命令行窗口，使用`curl`去创建一个请求，到一个运行的web服务

   ```shell
   $ curl http://localhost:8080/albums \
   --include \
   --header "Content-Type: application/json" \
   --request "POST" \
   --data '{"id":"4","title": "晴天", "artist":"周杰伦","price":49.99}'
   ```

   命令行将展示头信息和用来添加albums的JSON。

   ```shell
   $ curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id":"4","title": "晴天", "artist":"周杰伦","price":49.99}'
   HTTP/1.1 201 Created
   Content-Type: application/json; charset=utf-8
   Date: Sun, 28 Aug 2022 03:33:58 GMT
   Content-Length: 87
   
   {
       "id": "4",
       "title": "晴天",
       "artist": "周杰伦",
       "price": 49.99
   } $
   ```

4. 就像上一部分，使用`curl`取出所有的albums，能用来验证新的album是否添加进去

   命令行将展示所有的album列表。

   ```shell
   $ curl http://localhost:8080/albums
   [
       {
           "id": "1",
           "title": "Blue Train",
           "artist": "John Coltrane",
           "price": 56.99
       },
       {
           "id": "2",
           "title": "Jeru",
           "artist": "Gerry Mulligan",
           "price": 17.99
       },
       {
           "id": "3",
           "title": "Sarah Vaughan and Clifford Brown",
           "artist": "Sarah Vaughan",
           "price": 39.99
       },
       {
           "id": "4",
           "title": "晴天",
           "artist": "周杰伦",
           "price": 49.99
       }
   ]$
   ```

下一部分，我们将添加代码，用于处理`GET`用于特定的项。

## 五、写一个处理器，返回特定的项目

当一个客户端制造一个GET请求`/albums/[id]`，你想要返回与路径参数ID匹配的album。

为了做这个事情：

- 添加请求album的请求逻辑；
- 将路径影射到逻辑；

### 5.1 写代码

1. 在上一部分你添加的`postAlbums`函数下面，粘贴下面的代码，用于取回特定的album。

   这个`getAlbumByID`函数将从请求路径中提炼出ID，然后定位到匹配到album。

   ```go
   // getAlbumByID 定位于客户端发送请求到id匹配的album，然后返回album作为响应
   func getAlbumByID(c *gin.Context) {
     id := c.Param("id")
     
     // 遍历albums列表，查找与参数ID匹配的album
     for _, a := range albums {
       if a.ID == id {
         c.IndentedJSON(http.StatusOK, a)
         return
       }
     }
     c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album not found"})
   }
   ```

   在这个代码中：

   - 使用`Context.Param`从URL路径中提取id路径参数。当你映射请求器到请求路径，你将包含一个占位符，用于路径上的参数。

   - 遍历在切片结构中的album，查看一个与ID字段匹配的项。如果找到，你能序列化album结构成为JSON，并且返回它作为响应，带有`200 OK`HTTP码。

     就像上面提到的，一个真实的服务使用数据库去做查询。

   - 返回404错误，带有`http.Status.NotFound`，如果album没有发现。

   2. 最后，改变你的`main`函数，以便于它包含一个新的调用`rounter.GET`，路径是`/albums/:id`，就像下面展示的：

      ```go
      func main() {
      	router := gin.Default()
      	router.GET("/albums", getAlbums)
      	router.POST("/albums", postAlbums)
        router.GET("/albums/:id", getAlbumByID)
      	router.Run("localhost:8080")
      }
      ```

      在这个代码中：

      - 关联`/albums/:id`路径，带有`getAlbumByID`函数。在Gin中，项目前的冒号，表明这个项是路径参数。

   ### 5.2 运行代码

   1. 如果从上一步服务仍然在运行，停止它。

   2. 在包含`main.go` 文件的目录的命令行下，运行代码，启动服务：

      ```shell
      $ go run .
      [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
      
      [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
       - using env:	export GIN_MODE=release
       - using code:	gin.SetMode(gin.ReleaseMode)
      
      [GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
      [GIN-debug] POST   /albums                   --> main.postAlbums (3 handlers)
      [GIN-debug] GET    /albums/:id               --> main.getAlbumByID (3 handlers)
      [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
      Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
      [GIN-debug] Listening and serving HTTP on localhost:8080
      
      ```

   3. 在一个不同的命令行窗口，使用`curl`制造一个请求到你运行的web服务。

      ```shell
      curl http://localhost:8080/albums/2
      ```

      这个命令将展示一个id在使用的album的JSON。如果album没有发现，你将获取一个错误消息。

      ```shell
      $ curl http://localhost:8080/albums/2
      {
          "id": "2",
          "title": "Jeru",
          "artist": "Gerry Mulligan",
          "price": 17.99
      }lifeideMacBook-Pro:~ lifei$ curl http://localhost:8080/albums/4
      {
          "Message": "album not found"
      }$
      ```

      

   

# lagou

Golang编写的小工具，可以从拉勾网爬取指定岗位的职位内容信息，并支持对IT相关职位进行关键词词频分析排序，更方便的了解到岗位所需要技能。使用net/http标准库进行网络请求，方便拓展，可以方便的替换cookie。使用清华大学开放中文词库和基于[huichen/sego](https://github.com/huichen/sego)的分词器进行分词，支持可配置的词过滤，同义替换，支持自定义词典。使用[gizak/termui](https://github.com/gizak/termui)以表格的形式在命令行显示数据并支持分页。不依赖数据库等开发环境，方便使用，包含较为完备的单元测试。

- 选择岗位
![选择岗位](http://www.shawpo.me/media/posts/lagou/lagoukdselect.png)

- 关键词词频排序
![关键词词频排序](http://www.shawpo.me/media/posts/lagou/lagoudisplay.png)

- 表格分页
![表格分页](http://www.shawpo.me/media/posts/lagou/lagoudispalypage.png)

- cookie替换：浏览器打开拉勾网，登录后打开任意一个职位页面，在浏览器开发者工具中找到请求cookie，并用得到的cookie替换spider/request.go中的静态变量COOKIE。
![cookie替换](http://www.shawpo.me/media/posts/lagou/lagoucookie.png)

### 运行

```
go get github.com/shawpo/lagou/...

cd $GOPATH github.com/shawpo/lagou/

go run main.go
```

### 更多分析结果截图

- java关键词词频排序1
![java关键词词频排序1](http://www.shawpo.me/media/posts/lagou/lagoujava1.png)

- java关键词词频排序2
![java关键词词频排序2](http://www.shawpo.me/media/posts/lagou/lagoujava2.png)

- python关键词词频排序
![python关键词词频排序](http://www.shawpo.me/media/posts/lagou/lagoupython.png)
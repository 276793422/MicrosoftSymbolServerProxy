# MicrosoftSymbolServerProxy
微软符号服务器的一个跳板

由于我在工作中，不定时地需要使用微软的调试工具——Windbg。
不定期地需要通过它来下载微软的符号。

但是近期出现了一个问题，就是微软的符号服务器无法连接了。
是因为微软符号服务器的302跳转连接被砍掉了。

我也很悲痛。

如果是我自己做的话，我还能忍受，
但是当我来到了一个新的公司，
这个奇葩公司有各种奇葩限制。
比如，连接公司VPN之后，无法连接公司内网的其他机器，包括内网符号服务器，
断开VPN之后，又无法连接微软服务器，无法下载符号。
这就导致，如果我要下载公司的符号，就无法下载微软的符号，
如果我要下载微软的符号，我就无法下载公司的符号。
好纠结。

今天，我终于不就结了，我决定了，用我的香港服务器，做个墙外的跳板，
给我本地做个定向带里来下载微软符号。

写代码，30分钟解决，很简单，我是用的go语言，echo http 服务端框架，
然后基于这个框架，做了一个代理的分支。

整套东西30来分钟就解决了，其实应该可以10分钟或者5分钟解决的，
但是我好久不写go了，找IDE也找了半天。。。

好了，就说到这，
版本库里面包括可执行文件和代码，
没有技术含量，随便玩吧。

呃，还是再说一下使用方法吧。
其实只要将bin目录里面的文件放置到某个目录中，
然后设置好json 中的内容，之后启动程序就好了。

符号服务器设置：
SRV*d:\symbol\mssymbols*http://XXXXXX/download/symbols
然后WinDBG就可以使用了。

具体的工作流程是，客户端向服务端请求指定的路径，
服务端如果不存在指定的路径，就去微软服务器询问同样的路径，
并且获取指定的文件，先下载到服务器，
然后服务器将下载到的数据发送给本地。

服务器的外网速度如果飞快的话，那么其实不会有太大问题。
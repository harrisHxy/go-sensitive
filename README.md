### GO-敏感词过滤服务

    基于go语言和开源的包封装的一个敏感词过滤服务，可微服务独立部署，提供api给业务使用，占用内存少，搜索匹配快
    
#### 使用步骤

    1，# git clone git@github.com:harrisHxy/go-sensitive.git
    
    2, # go get github.com/anknown/ahocorasick
       # go get github.com/gin-gonic/gin
       
    3，创建mysql敏感词表 ’sensitive_words‘：
        CREATE TABLE `sensitive_words` (
          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
          `words` varchar(20) NOT NULL DEFAULT '' COMMENT '敏感词',
          `add_time` int(10) NOT NULL DEFAULT '0',
          PRIMARY KEY (`id`)
        ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4
        并插入一些测试敏感词数据
        
    4，修改main.go中25行的mysql配置
    
    5，#go run main.go
    
    6，在浏览器输入：http://localhost:8282/match?words=你去死吧
    
    7，可以看到返回json数据
    
    8，done

#### 服务框架

    github.com/gin-gonic/gin

#### 算法

    Double Array Trie + AhoCorasick

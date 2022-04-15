# Counter_Golang  
Implement a counter module that requires in-process coroutines to be safe, asynchronous, and high-performance statistics based on key-value indicators.  
  
# 文件介绍与成果展示  
根目录下的lib.go是二面的笔试题（new version），another文件夹下的lib.go是上次的版本。我在一个新建项目下引用该包，go.mod配置如下：  
![Moodle Pic](https://github.com/EkkoSummer/goCounter/blob/main/image/pre.png)  
运行结果如下：  
![Moodle Pic](https://github.com/EkkoSummer/goCounter/blob/main/image/v4.png)   
其中第一行为本次新版本的运行时间，第二行为老版本的运行时间，可见更新后运行时间提升了三倍。个人分析认为由于本次手写了读写锁，而上一次是用封装好的sync.Map，sync.Map再去调用具体功能时造成了额外的时间开销。  

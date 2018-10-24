
# 尝试通过整式是学习的方式学习一下Go语言

## 小目标 写1000行代码


## 环境搭建
- goLang + GOPATH（项目目录） + GOROOT（环境目录） 等环境变量
- 编辑器 Sublime + goSublime（ https://packagecontrol.io/packages/GoSublime 英文不好影响了效率 ） + gofmt
- 


## 框架
- 基础知识笔记
    + 常量 const  iota 可以自增    fallthrough (which 向下执行)
    + goroute go func() 高并发的方式
    +  多个  goroute 通过 channel 通信   make 创建 
    ```
        pipe := make(chan int, 3)
        pipe <- 1
        pipe <- 2
        t1 := <-pipe
    ```
    + 包内大写 可以被引用  包内小写 只能本文件使用
    + 包的概念 （编码都是 utf8）
        * 顺序全局变量  init 函数  main 函数
        * import 导入但是不引用 使用 下划线 _ "packageName" 包_   X "packageName" X 是别名
    + 数组[] 、 切片[:] 、 MAP[KEY->VALUE] (使用 make 初始化)

- 线程同步  互斥锁、读写锁    import sync   原子操作 sync/atomic
- 模块 http
-

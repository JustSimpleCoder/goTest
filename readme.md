
# 尝试通过整式是学习的方式学习一下Go语言

## 小目标 先写10000行代码


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
    + 结构体  new(Student)  &Student  都是初始化结构体的方法  返回指针  stu.  == ×stu  tag 信息   匿名字段
    ```
    type ST struct {
        Name string `json:"stname"` `db:"name"`
        Age  int    `json:"stage"`
        Math
    }
    ST.Math.scores = 100

    func(receiver type) method(参数) 返回值 {}
    ```
    + 面向接口编程 interface
        * 只能有方法
        * 接口可以嵌套
        * if v,ok = f.(interfacefunName);ok{}  是否实现了此接口
    + 反射 动态的获取变量类型、如果是结构体可以调用里面的方法（个人预测项目会使用较多） 
    ```
    v := reflect.ValueOf(x)
    v.Elem().SetInt(10)
    ```

- 线程同步  互斥锁、读写锁    import sync   原子操作 sync/atomic
- file 文件操作
    + os.OpenFile  写   os.Open 读
    + ioutil  统一读出写入

- goroutine   M P G
    + V1.8 下要设置CPU 核数
    + 可以通过 全局变量数据传递
    + 管道定义 var x chan type   用make 初始化   close(strChan) 关闭  range 遍历   var x chan<- int  只写  var x <-chan int 只读
    + 管道的 select 操作 解决一直阻塞  select  case default
    + recover() 捕获错误
- 单元测试
    + 文件命名 xxxx_test.go    函数名  Testxxxxx(t *testing.T) Test开头
- Net Socket 服务
- 



## 项目其他
- 服务注册和服务发现
- 一致性hash

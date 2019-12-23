# GO

## 创始人

![20191223100829.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191223100829.png)

Rob Pike

Unix 的早期开发者  UTF-8 创始⼈人

![20191223100920.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191223100920.png)

Ken Thompson

Unix 的创始⼈人 C 语⾔言创始⼈人 1983 年年获图灵奖

![20191223101003.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191223101003.png)

Robert Griesemer

Google V8 JS Engine 开发者 Hot Spot 开发者

## GO语言特点

#### 1. 为并发而生

Go语言在多核并发上拥有原生的设计优势，Go语言从底层原生支持并发。

Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。Go 语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用CPU性能。开启一个goroutine的消耗非常小（大约2KB的内存），你可以轻松创建数百万个goroutine。

goroutine的特点：

1. goroutine具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存。
2. goroutine的启动时间比线程快。
3. goroutine原生支持利用channel安全地进行通信。
4. goroutine共享数据结构时无需使用互斥锁。

#### 2. 简单

Go 语言简单易学，学习曲线平缓，只有25个关键字。其语法在C语言的基础上进行了大幅的简化，去掉了不需要的表达式括号，循环也只有 for 一种表示方法，就可以实现数值、键值等各种遍历。

#### 3. 效率:快速的编译时间，开发效率和运⾏效率⾼

![20191223111832.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191223111832.png)

Go语言实现了开发效率与执行效率的完美结合，让你像写Python代码（效率）一样编写C代码（性能）。

#### 4. ⾃自由⾼高效:组合的思想、无侵⼊式的接⼝

Go语⾔言⽀支持当前所有的编程范式，包括过程式编程、⾯面向对象编程、⾯面向 接⼝口编程、函数式编程。程序员们可以各取所需、⾃自由组合、想怎么玩就怎么玩。

#### 5. 稳定性

Go拥有强⼤大的编译检查、严格的编码规范和完整的软件⽣生命周期⼯工具，具
有很强的稳定性，稳定压倒⼀一切。那么为什什么Go相⽐比于其他程序会更更稳定呢? 这是因为Go提供了了软件⽣生命周期(开发、测试、部署、维护等等)的各个环节 的⼯工具，如go tool、gofmt、go test。



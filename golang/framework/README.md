框架源码
<pre>
	cfg -- config.ini 读取, 写log
	db -- 数据库驱动
	helper -- 辅助函数，candy funcs....
    	stack.go -- 打印堆栈消息
        profilingtool.go -- 打印GC相关信息
		rand.go	-- 全局快速随机数发生器
	hub -- HUB服务器
    	sys.go -- 系统goroutine打印系统状态和gc
	misc -- 算法等
    	timer -- 通用计时器
        packet -- 解析包
	types -- 玩家数据结构
    tats -- 统计服务器（UDP)
    	protos -- 协议相关处理，协议号和对应的处理函数
    agent -- GAME服务器(UDP)
    	stats_client -- 统计服务器客户端程序
	gamedata -- 提供一个类似的内存二维表。
 ==========未实现==========

    scripts -- awk bash 脚本
    event -- EVENT服务器
    inspect -- Telnet Console for GS
</pre>

问题总结：

1.在显示查找结果/viewlog时，总是显示乱码，经过两个小时的排查，问题主要有两个地方
	1）首先是调用render_template函数时，对最后两个参数的赋值应该保持与原有类型不变，不该赋值为‘var’形式
	2）对于flask下html的for循环使用，和python应该遵循一样的缩进原则

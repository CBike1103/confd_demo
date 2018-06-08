import os,sys,re,traceback

#从命令行参数获取变化的key
para = sys.argv
k = para.pop()

#从对应的log文件获取value
try:
	f = open(k+'.log', 'r')
	v = f.read()
finally:
	if f:
		f.close()

		#写入新的文件
		filePath = r''+k+'.php'
		class_file = open(filePath,'w')
		class_file.writelines(k+':'+v)
		class_file.close()
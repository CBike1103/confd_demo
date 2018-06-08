import os,sys,re,traceback

filePath = r'demo.php'
class_file = open(filePath,'w')
class_file.writelines(str(sys.argv))
class_file.close()
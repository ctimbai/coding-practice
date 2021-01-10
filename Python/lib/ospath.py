#!/usr/bin/env python3
#coding=utf-8

import os.path

path = '/home/python/doc/file.txt'

print(os.path.basename(path))
print(os.path.dirname(path))

info = os.path.split(path) #分成文件名和目录
path2 = os.path.join('/', 'home', 'python', 'doc', 'file1.txt')
print(info, path2)

l = [path, path2]
print(os.path.commonprefix(l)) #查询多个路径的共同部分

path = '/home/python/../.'
print(os.path.normpath(path)) #去除路径中的冗余，比如'/home/python/../.'被转化为'/home'

print(os.path.exists(path))
print(os.path.getatime(path)) #上次读取的时间
print(os.path.getmtime(path)) #上次修改的时间

print(os.path.isfile(path)) #是否是文件
print(os.path.isdir(path)) #是否是目录
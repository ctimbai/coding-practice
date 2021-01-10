#!/usr/bin/env python3
#coding=utf-8

def change_integet(a):
    a = a + 1
    return a #不用定义返回值

a = 1
change_integet(1)
print(a)

def change_list(a):
    a[0] += 1
    return a

a = [1,2,3]
change_list(a)
print(a)

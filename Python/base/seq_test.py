#!/usr/bin/env python3
#coding=utf-8

# seq include tuple and list
t1 = (2, 1.2, 'love', False) # 不可变
l1 = [1, True, 'smile']

print(t1, type(t1))
print(l1, type(l1))

print(t1[:])
print(t1[:1])
print(t1[1:])
print(t1[-1])

print(l1[:])
print(l1[:1])
print(l1[1:])
print(l1[-1])

# string is spec tuple
str = 'bacckd'
print(str[2:4])
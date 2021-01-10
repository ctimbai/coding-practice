#!/usr/bin/env python3
#coding=utf-8

f = open("e:\Git\coding-practice\Python\\base\\test.txt", "r+")
c = f.read(5) #读取10个字节
print(c)
c = f.readline()
print(c)
c = f.readlines()
print(c)

f.write("I like apple")
c = f.readlines()
print(c)
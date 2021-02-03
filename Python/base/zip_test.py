#!/usr/bin/env python3
#coding=utf-8

ta = [1,2,3]
tb = [7,8,9]

for i, j in zip(ta, tb):
    print(i, j)

# cluster
zipped = zip(ta, tb)
print(zipped)

# decompose
na, nb = zip(*zipped)
print(na, nb)
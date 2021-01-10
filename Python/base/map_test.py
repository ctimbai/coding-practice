#!/usr/bin/env python3
#coding=utf-8

m = {'tom':1, 'sam': 5}

m['lili'] = 7

print(m, len(m), type(m))

for k in m:
    print(k, m[k])

print(m.keys())
print(m.values())
print(m.items())

del m['tom']
print(m)

m.clear()
print(m)
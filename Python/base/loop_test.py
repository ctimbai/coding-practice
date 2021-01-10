#!/usr/bin/env python3
#coding=utf-8

# 1
for a in [1, 2.2, 'hello']:
    print(a)

# 2
idx = range(5)
print(idx)

# 3
for idx in range(10):
    print(idx ** 2)

# 4
for i in range(10):
    if i == 2:
        continue
    elif i == 8:
        break
    print(i)

# 5 while
i = 0
while i < 10:
    if i == 2:
        print(i)
    i += 1
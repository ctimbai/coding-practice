#!/usr/bin/env python3
#coding=utf-8

class Bird(object):
    have_feather = True
    way_of_reproduction = 'egg'

    def __init__(self, info):
        print("hello", info)

    def move(self, dx, dy):
        position = [0, 0]
        position[0] = position[0] + dx
        position[1] = position[1] + dy
        return position
    
    def show_feather(self):
        print(self.have_feather)

class Chicken(Bird):
    have_feather = False #覆盖
    way_of_move = 'walk'
    possible_in_KFC = True

    def __init__(self, info):
        print("Hi", info)

class Oriole(Bird):
    way_of_move = 'fly'
    possible_in_KFC = False

summer = Chicken("world")
print(summer.have_feather)
print(summer.move(3, 4))

print(dir(list))

l1 = [1,2,3]
l2 = [4,5,6]
l = l1 + l2
print(l) # use __add__()

print(l.count(5))
print(l.index(2))
l.sort()
print(l)
print(l.pop())
l.remove(2)
print(l)
l.insert(0, 9)
print(l)

class super_list(list):
    def __sub__(self, b):
        a = self[:]
        b = b[:]
        while len(b) > 0:
            elb = b.pop()
            if elb in a:
                a.remove(elb)
        return a

print(super_list([1,2,3]) - super_list([3,4,5]))
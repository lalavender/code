#!/usr/bin/python3
# -*- coding: utf-8 -*-

from collections import deque

"""
可以当作限定大小的列表使用
"""

d = deque('ghji')
for i in d:
    print(i)
d.append('x')
d.appendleft('a')
print(d)
d.pop()
d.popleft()
print(d)
print(list(d))
d.extend('456789')
print(d)
d.rotate(1) # 向右侧反转
d.rotate(-1) # 向左侧反转
deque(reversed(d))
d.clear()




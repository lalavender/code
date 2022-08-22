#!/usr/bin/python3
# -*- coding: utf-8 -*-

import itertools

# 无限循环迭代器
n = itertools.count(1)
for i in n:
    print(i)
    break

# 无限循环序列
cs = itertools.cycle("ABC")
for i in cs:
    print(i)
    break

# 指定序列重复次数
ns = itertools.repeat("a", 6)
for i in ns:
    print(i)
    break
# 根据序列截取出有限的序列
ts = itertools.takewhile(lambda x: x < 10, n)
print(list(ts))

# 组合序列
for i in itertools.chain("avaffaf", "dadwaedweqwefsedf"):
    print(i)
    break

# 相邻重复的元素挑出来放在一起
for i,j in itertools.groupby("AAAAAABBBVFFRREF"):
    print(i,list(j))

for key, group in itertools.groupby('AaaBBbcCAAa', lambda c: c.upper()):
    print(key, list(group))

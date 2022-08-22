#!/usr/bin/python3
# -*- coding: utf-8 -*-


from collections import defaultdict

"""
defaultdict是内置数据类型dict的一个子类，
基本功能与dict一样，只是重写了一个方法__missing__(key)和增加了一个可写的对象变量default_factory。
"""

s = [('red', 1), ('blue', 2), ('red', 3), ('blue', 4), ('red', 1), ('blue',
4)]
d = defaultdict(set)
for k,v in s:d[k].add(v)
print(d.items())
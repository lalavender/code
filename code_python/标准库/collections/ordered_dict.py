#!/usr/bin/python3
# -*- coding: utf-8 -*-
from collections import OrderedDict
"""
有序字典：记住了字典插入的顺序
"""
d = {"banana":3,"apple":2,"pear":1,"orange":4}
# 按键排序
d = OrderedDict(sorted(d.items(),key = lambda t:t[0]))
print(d)
# 按值排序
d = OrderedDict(sorted(d.items(),key = lambda t:t[1]))
print(d)
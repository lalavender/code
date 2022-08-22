#!/usr/bin/python3
# -*- coding: utf-8 -*-
from collections import namedtuple
"""
命令元组，给元组中的每个位置赋予含义
可以通过名称来获取字段信息而不仅仅是通过位置索引
"""


point = namedtuple('Point', ['x', 'y'])
t = [33,44]
# ----------------------------
# 从已有的序列中创建对象
point._make(t)
p = point(x=33, y=44)
print(p)
# ----------------------------
# 返回一个OrderDict，由名称到对应值建立的映射
p_dict = p._asdict()
print(p_dict)
# ----------------------------
# 回一个新的namedtuple对象，用新值替换指定名称中的值
p_replace = p._replace(x = 55)
print(p_replace)
# ----------------------------
print(p._fields)

# 当名称存储在字符串中，可以使用getattr()函数进行检索
print(getattr(p,'x'))

# 使用**操作符，可以将一个字典转换成namedtuple
d  = {'x':11,'y':22}
print(point(**d))










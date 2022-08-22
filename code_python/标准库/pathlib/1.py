#!/usr/bin/python3
# -*- coding: utf-8 -*-
import os
import pathlib

# 获取当前路径，返回对象，可用str()转换
print(pathlib.Path.cwd())
cwd = pathlib.Path.cwd()

# 获取用户home目录
print(pathlib.Path.home())

# 当前文件路径绝对路径
print(pathlib.Path(__file__).resolve())

# 文件属性
file = pathlib.Path(__file__).resolve()
print(file.stat())
print(file.stat().st_size)
print(file.stat().st_atime)
print(file.stat().st_ctime)
print(file.stat().st_mtime)

# 路径组成部分
print(file.name)  # 文件名/目录名
print(file.stem)
print(file.suffix)
print(file.parent)
print(list(file.parents))
print(file.anchor)

x = [path for path in cwd.iterdir() if cwd.is_dir()]
print(x)

# 路径拼接
print(pathlib.Path.home().joinpath('dir', 'file.txt'))

# 路径判断
print(cwd.is_file())
print(cwd.is_dir())
print(cwd.exists())

# 创建文件
file = pathlib.Path('test.txt')
file.touch(exist_ok=True) # exist_ok=True 文件存在则忽略
# 创建目录
path = pathlib.Path("test1/test2/test3")
path.mkdir(parents=True)  # parents=True 创建多层

# 删除目录
path.rmdir()  # 只能删除一级且为空
# 删除文件
pathlib.Path(cwd/'test.txt').unlink()

# 打开文件
file = pathlib.Path('test.txt')
file.write_text("test info")
print(file.read_text())
file.write_bytes(b"test info")
print(file.read_bytes())

# 移动文件
file = pathlib.Path('test.txt')
file.touch()
tmp = file.replace('test1/demo1') #移动目录要存在

# 重命名
new_file = tmp.with_name('new.txt')
tmp.replace(new_file)

# 修改后缀名
file = pathlib.Path('test1/new.txt')
new_file = file.with_suffix('.json')
file.replace(new_file)
print('-'*20)
import shutil
shutil.rmtree('test1')












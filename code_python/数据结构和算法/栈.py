#!/usr/bin/python3
# -*- coding: utf-8 -*-

class Stack:
    def __init__(self):
        self.items = []

    def is_empty(self):
        return self.items == []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        return self.items.pop()

    def peek(self):
        return self.items[len(self.items) - 1]

    def size(self):
        return len(self.items)


# 匹配括号

def par_checker(str_symbol):
    s = Stack()
    balanced = True
    idx = 0
    while idx < len(str_symbol) and balanced:
        symbol = str_symbol[idx]
        if symbol == "([{":
            s.push(symbol)
        else:
            if s.is_empty():
                balanced = False
            else:
                top = s.pop()
                if not matches(top, symbol):
                    balanced = False
        idx += 1
    if balanced and s.is_empty():
        return True
    else:
        return False


def matches(open, close):
    opens = "([{"
    closers = ")]}"
    return opens.index(open) == closers.index(close)


# 十进制转二进制
def divide_by2(num):
    s = Stack()
    while num > 0:
        rem = num % 2
        s.push(rem)
        num = num // 2
    str_bin = ""
    while not s.is_empty():
        str_bin += str(s.pop())
    return str_bin


# 十进制转化为十六以下任意进制
def base_converter(num, base):
    digits = "0123456789ABCDEF"
    s = Stack()
    while num > 0:
        rem = num % base
        s.push(rem)
        num //= base
    str_new = ""
    while not s.is_empty():
        str_new += digits[s.pop()]
    return str_new


#通用中缀转后缀

if __name__ == '__main__':
    print(par_checker("((()))"))
    print(par_checker("((()))))"))
    print(divide_by2(42))
    print(base_converter(25,2))
    print(base_converter(25,16))

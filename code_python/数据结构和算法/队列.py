#!/usr/bin/python3
# -*- coding: utf-8 -*-

class Queue:
    def __init__(self):
        self.items = []

    def is_empty(self):
        return self.items == []

    def enqueue(self, item):
        self.items.insert(0, item)

    def dequeue(self):
        return self.items.pop()

    def size(self):
        return len(self.items)


def hot_potato(name_list, n):
    q = Queue()
    for name in name_list:
        q.enqueue(name)
    while q.size() > 1:
        for i in range(n):
            q.enqueue(q.dequeue())
        q.dequeue()
    return q.dequeue()


class Printer:
    def __init__(self, ppm):
        self.page_rate = ppm  # 打印速度
        self.cur_task = None  # 打印任务
        self.time_remaining = 0  # 任务倒计时

    def tick(self):
        """ 打印1s"""
        if self.cur_task is not None:
            self.time_remaining = self.time_remaining - 1
            if self.time_remaining <= 0:
                self.cur_task = None

    def is_busy(self):
        if self.cur_task is not None:
            return True
        else:
            return False

    def start_next(self, new_task):
        self.cur_task = new_task
        self.time_remaining = new_task.get_pages() * 60 / self.page_rate


print(hot_potato(["1", "2", "3", "4", "5", "6"], 7))



#双端队列


#!/usr/bin/python3
# -*- coding: utf-8 -*-

from collections import Counter



c = Counter("gallahad")
print(c)
c = Counter({'red': 4, 'blue': 2})
print(c)
c = Counter(cats=4, dogs=8)
print(c)
c = Counter(['eggs', 'ham'])
print(c['123'])
print(c['eggs'])

cnt = Counter()
wordList = ["a", "b", "c", "c", "a", "a"]
for word in wordList:
    cnt[word] += 1
print(cnt)
print(list(cnt.elements()))
print(list(cnt.most_common()))
c = Counter(a=4,b=2,c=0,d=-2)
d = Counter(a=1,b=2,c=-3,d=4)
c.subtract(d)
print(c)
c.update(d)










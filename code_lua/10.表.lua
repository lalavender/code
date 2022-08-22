-- 简单的 table
mytable = {}
print("---------------------------------------")
print("mytable 的类型是 ",type(mytable))
mytable[1]= "Lua"
mytable["wow"] = "修改前"
print("---------------------------------------")
print("mytable 索引为 1 的元素是 ", mytable[1])
print("mytable 索引为 wow 的元素是 ", mytable["wow"])

-- alternatetable和mytable的是指同一个 table
alternatetable = mytable
print("---------------------------------------")
print("alternatetable 索引为 1 的元素是 ", alternatetable[1])
print("mytable 索引为 wow 的元素是 ", alternatetable["wow"])

alternatetable["wow"] = "修改后"
print("---------------------------------------")
print("mytable 索引为 wow 的元素是 ", mytable["wow"])

-- 释放变量
alternatetable = nil
print("---------------------------------------")
print("alternatetable 是 ", alternatetable)

-- mytable 仍然可以访问
print("---------------------------------------")
print("mytable 索引为 wow 的元素是 ", mytable["wow"])

mytable = nil
print("---------------------------------------")
print("mytable 是 ", mytable)


-- table操作函数
fruits = {"banana","orange","apple"}
print("---------------------------------------")
print("连接后的字符串 ",table.concat(fruits))
print("连接后的字符串 ",table.concat(fruits,", "))
print("连接后的字符串 ",table.concat(fruits,", ", 2,3))

print("---------------------------------------")
fruits = {"banana","orange","apple"}

-- 在末尾插入
table.insert(fruits,"mango")
print("索引为 4 的元素为 ",fruits[4])

-- 在索引为 2 的键处插入
table.insert(fruits,2,"grapes")
print("索引为 2 的元素为 ",fruits[2])

print("最后一个元素为 ",fruits[5])
table.remove(fruits)
print("移除后最后一个元素为 ",fruits[5])

print("---------------------------------------")
fruits = {"banana","orange","apple","grapes"}
print("排序前")
for k,v in ipairs(fruits) do
        print(k,v)
end

table.sort(fruits)
print("排序后")
for k,v in ipairs(fruits) do
        print(k,v)
end


print(type("Hello world")) -- > string
print(type(10.4 * 3)) -- > number
print(type(print)) -- > function
print(type(type)) -- > function
print(type(true)) -- > boolean
print(type(nil)) -- > nil
print(type(type(X))) -- > string

-- 全局变量或者tabke赋值nil，相当于将他们删除

tab1 = {
    key1 = "val1",
    key2 = "val2",
    "val3"
}
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end

tab1.key1 = nil
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end

-- nil作比较时应该加引号
print(type(x))
print(type(x) == nil)
print(type(x) == "nil")
-- 如果没有使用type() 则不需要,type() 返回类型为string
print(x == nil)

-- lua 自动转换string和number
print("10" + 1)
print("10 + 1")
print(10 .. 20)

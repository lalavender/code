-- while 循环
a = 10
while (a < 20) do
    print("a 的值为:", a)
    a = a + 1
end

-- for 循环
for i = 10, 1, -1 do
    print(i)
end

--  f(x) 只在循环开始前执行一次
function f(x)
    print("function")
    return x * 2
end
for i = 1, f(5) do
    print(i)
end

-- 通过ipairs迭代器遍历数组,遇到nil停止
a = {"one", "two", "three"}
for i, v in ipairs(a) do
    print(i, v)
end

-- pairs 遍历键值对
x = {a="123",b="456"}
for i, v in pairs(x) do
    print(i, v)
end

-- if 语句
print("----------if 语句---------------")
a = 10;
if( a < 20 )
then
   print("a 小于 20" );
end
print("a 的值为:", a);
print("----------if else语句---------------")
a = 100;
if( a < 20 )
then
   print("a 小于 20" )
else
   print("a 大于 20" )
end
print("a 的值为 :", a)
print("----------if elseif else语句---------------")
a = 100
if( a == 10 )
then
   print("a 的值为 10" )
elseif( a == 20 )
then
   print("a 的值为 20" )
elseif( a == 30 )
then
   print("a 的值为 30" )
else
   print("没有匹配 a 的值" )
end
print("a 的真实值为: ", a )










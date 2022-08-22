string1 = 'Lua'
print('"字符串 1 是"', string1)
string2 = 'runoob.com'
print('字符串 2 是', string2)

string3 = [["Lua 教程"]]
print('字符串 3 是', string3)
print('---------字符串方法---------------------')
local s = 'lua'
print(string.upper(s))
print(string.lower(s))
print(string.gsub('aaaa', 'a', 'z', 3))
print(string.find('Hello Lua user', 'Lua', 1))
print(string.reverse('Lua'))
print(string.format('the value is:%d', 4))
print(string.char(97, 98, 99, 100))
print(string.byte('ABCD', 4))
print(string.byte('ABCD')) -- 默认第一个字符
print(string.len('abc'))
print(string.rep('abcd', 2))
print('www.runoob.' .. 'com')
for word in string.gmatch('Hello Lua user', '%a+') do
    print(word)
end
print(string.match('I have 2 questions for you.', '%d+ %a+'))
print(string.format('%d, %q', string.match('I have 2 questions for you.', '(%d+) (%a+)')))

print("-------------------字符串截取-----------------------")
local sourcestr = "prefix--runoobgoogletaobao--suffix"
print("\n原始字符串", string.format("%q", sourcestr))

-- 截取部分，第1个到第15个
local first_sub = string.sub(sourcestr, 4, 15)
print("\n第一次截取", string.format("%q", first_sub))

-- 取字符串前缀，第1个到第8个
local second_sub = string.sub(sourcestr, 1, 8)
print("\n第二次截取", string.format("%q", second_sub))

-- 截取最后10个
local third_sub = string.sub(sourcestr, -10)
print("\n第三次截取", string.format("%q", third_sub))

-- 索引越界，输出原始字符串
local fourth_sub = string.sub(sourcestr, -100)
print("\n第四次截取", string.format("%q", fourth_sub))

--[[
.(点): 与任何字符配对
%a: 与任何字母配对
%c: 与任何控制符配对(例如\n)
%d: 与任何数字配对
%l: 与任何小写字母配对
%p: 与任何标点(punctuation)配对
%s: 与空白字符配对
%u: 与任何大写字母配对
%w: 与任何字母/数字配对
%x: 与任何十六进制数配对
%z: 与任何代表0的字符配对
%x(此处x是非字母非数字字符): 与字符x配对. 主要用来处理表达式中有功能的字符(^$()%.[]*+-?)的配对问题, 例如%%与%配对
[数个字符类]: 与任何[]中包含的字符类配对. 例如[%w_]与任何字母/数字, 或下划线符号(_)配对
[^数个字符类]: 与任何不包含在[]中的字符类配对. 例如[^%s]与任何非空白字符配对
--]]

print("--------------字符串格式化----------------------")
--[[
    %c - 接受一个数字, 并将其转化为ASCII码表中对应的字符
    %d, %i - 接受一个数字并将其转化为有符号的整数格式
    %o - 接受一个数字并将其转化为八进制数格式
    %u - 接受一个数字并将其转化为无符号整数格式
    %x - 接受一个数字并将其转化为十六进制数格式, 使用小写字母
    %X - 接受一个数字并将其转化为十六进制数格式, 使用大写字母
    %e - 接受一个数字并将其转化为科学记数法格式, 使用小写字母e
    %E - 接受一个数字并将其转化为科学记数法格式, 使用大写字母E
    %f - 接受一个数字并将其转化为浮点数格式
    %g(%G) - 接受一个数字并将其转化为%e(%E, 对应%G)及%f中较短的一种格式
    %q - 接受一个字符串并将其转化为可安全被Lua编译器读入的格式
    %s - 接受一个字符串并按照给定的参数格式化该字符串

--]]
string1 = "Lua"
string2 = "Tutorial"
number1 = 10
number2 = 20
print(string.format("基本格式化 %s %s",string1,string2))
-- 日期格式化
date = 2; month = 1; year = 2014
print(string.format("日期格式化 %02d/%02d/%03d", date, month, year))
-- 十进制格式化
print(string.format("%.4f",1/3))
print(string.format("%c", 83))                 -- 输出S
print(string.format("%+d", 17.0))              -- 输出+17
print(string.format("%05d", 17))               -- 输出00017
print(string.format("%o", 17))                 -- 输出21
print(string.format("%u", 3.14))               -- 输出3
print(string.format("%x", 13))                 -- 输出d
print(string.format("%X", 13))                 -- 输出D
print(string.format("%e", 1000))               -- 输出1.000000e+03
print(string.format("%E", 1000))               -- 输出1.000000E+03
print(string.format("%6.3f", 13))              -- 输出13.000
print(string.format("%q", "One\nTwo"))         -- 输出"One\Two                            -- 　　Two"
print(string.format("%s", "monkey"))           -- 输出monkey
print(string.format("%10s", "monkey"))         -- 输出    monkey
print(string.format("%5.3s", "monkey"))        -- 输出  mon


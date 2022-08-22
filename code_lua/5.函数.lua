-- 基本函数定义
function max(num1, num2)
    local result = 0
    if (num1 > num2) then
        result = num1;
    else
        result = num2;
    end
    return result;
end
-- 调用函数
print("-------------------------")
print("两值比较最大值为 ", max(10, 4))
print("两值比较最大值为 ", max(5, 6))

-- 函数多返回值
function maximum(a)
    local mi = 1 -- 最大值索引
    local m = a[mi] -- 最大值
    for i, val in ipairs(a) do
        if val > m then
            mi = i
            m = val
        end
    end
    return m, mi
end
print("-------------------------")
print(maximum({8, 10, 23, 12, 5}))

-- 可变参数
function add(...)
    local s = 0
    for i, v in ipairs {...} do -- > {...} 表示一个由所有变长参数构成的数组
        s = s + v
    end
    return s
end
print("-------------------------")
print(add(3, 4, 5, 6, 7))

-- 获取可变参数的数量
function average(...)
    local result = 0
    local arg = {...}
    for i, v in ipairs(arg) do
        result = result + v
    end
    print("总共传入 " .. select("#", ...) .. " 个数")
    return result / select("#", ...)
end
print("-------------------------")
print("平均值为", average(10, 5, 3, 4, 5, 6))




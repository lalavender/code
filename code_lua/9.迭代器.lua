-- 泛型for迭代器
array = {'Google', 'Runoob'}

for key, value in ipairs(array) do
    print(key, value)
end

-- 无状态的迭代器
function square(iteratorMaxCount, currentNumber)
    if currentNumber < iteratorMaxCount then
        currentNumber = currentNumber + 1
        return currentNumber, currentNumber * currentNumber
    end
end

for i, n in square, 3, 0 do
    print(i, n)
end

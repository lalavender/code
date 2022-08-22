-- 一维数组
array = {'Lua', 'Tutorial'}

for i = 0, 2 do
    print(array[i])
end

-- 多维数组
array = {}
for i = 1, 3 do
    array[i] = {}
    for j = 1, 3 do
        array[i][j] = i * j
    end
end

-- 访问数组
for i = 1, 3 do
    for j = 1, 3 do
        print(array[i][j])
    end
end

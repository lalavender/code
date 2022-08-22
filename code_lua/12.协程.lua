co =
    coroutine.create(
    function(i)
        print(i)
    end
)

coroutine.resume(co, 1) -- 1
print(coroutine.status(co)) -- dead

print('----------')

co =
    coroutine.wrap(
    function(i)
        print(i)
    end
)

co(1)

print('----------')

co2 =
    coroutine.create(
    function()
        for i = 1, 10 do
            print(i)
            if i == 3 then
                print(coroutine.status(co2)) --running
                print(coroutine.running()) --thread:XXXXXX
            end
            coroutine.yield()
        end
    end
)

coroutine.resume(co2) --1
coroutine.resume(co2) --2
coroutine.resume(co2) --3

print(coroutine.status(co2)) -- suspended
print(coroutine.running())

print('----------')

-- 生产者消费者
local newProductor

function productor()
    local i = 0
    while true do
        i = i + 1
        send(i) -- 将生产的物品发送给消费者
    end
end

function consumer()
    while true do
        local i = receive() -- 从生产者那里得到物品
        print(i)
    end
end

function receive()
    local status, value = coroutine.resume(newProductor)
    return value
end

function send(x)
    coroutine.yield(x) -- x表示需要发送的值，值返回以后，就挂起该协同程序
end

-- 启动程序
newProductor = coroutine.create(productor)
consumer()

print('------------------------------------')
co =
    coroutine.create(
    function(a)
        local r = coroutine.yield(a + 1) -- yield()返回a+1给调用它的resume()函数，即2
        print('r=' .. r) -- r的值是第2次resume()传进来的，100
    end
)
status, r = coroutine.resume(co, 1) -- resume()返回两个值，一个是自身的状态true，一个是yield的返回值2
coroutine.resume(co, 100) --resume()返回true

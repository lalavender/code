#include<string.h>
#include<stdio.h>

int main(int argc, char const *argv[])
{

    // string.h 标准库

    // 变量： size_t --> unsigned int 
    // 宏： NULL
    // 函数：

    int x;
    x = strcmp("abcv","wqewq");
    printf("%d\n",x);
    x = strncmp("abcv","wqewq",3);
    printf("%d\n",x);
    return 0;
}

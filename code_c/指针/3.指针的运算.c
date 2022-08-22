#include <stdio.h>

int main(void) {

  //指针的运算
  // char a[20];
  // printf("数组a的地址：%p\n", &a);
  // int *ptr = (int *)a; //强制类型转换并不会改变a 的类型
  // printf("指针ptr：%p\n", ptr);
  // ptr++;    // ptr = ptr +sizeof(int)
  // printf("指针ptr：%p\n",ptr);

  /*int i;
  int array[20] = {0};
  int *ptr = array;
  for (i = 0; i < 20; i++) {
    (*ptr)++;
    printf("值：%d\n", *ptr);
    ptr++;
    printf("指针ptr：%p\n", ptr);
  }*/
  char a[20] = "You_are_a_girl";
  int *ptr = (int *)a;
  ptr += 1;
  printf("值：%d\n", *ptr);

  return 0;
}
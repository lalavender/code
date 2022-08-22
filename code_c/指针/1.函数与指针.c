#include <stdio.h>
#include <stdlib.h>

/*
** 1. 将指针传递给函数
** 2. 声明函数指针：
** 本质上，函数表示法就是指针表示法。函数名字经过求值会变成函数的地址，
** 然后函数参数会被传递给函数。我们将会看到，函数指针为控制程序的执行流提供了新的选择。
** 程序栈——> 栈帧 ——> 存放函数参数和局部变量
** 栈帧：1. 返回地址 2.局部数据存储 3. 参数存储 4. 栈指针和基指针
**
**
**
**
**
**
**
**
*/

// 用指针传递数据
void swap_with_pointers(int *p1, int *p2) {
  int tmp;
  tmp = *p1;
  *p1 = *p2;
  *p2 = tmp;
}
// 传递普通变量
void swap(int p1, int p2) {
  int tmp;
  tmp = p1;
  p1 = p2;
  p2 = tmp;
}

// 返回指针
/*
** 返回未初始化的指针
** 返回指向无效地址的指针
** 返回局部变量的指针
** 返回指针但是没有释放内存
**
**
**
**
*/

int *allocate_array(int size, int value) {
  int *arr = (int *)malloc(size * sizeof(int));
  for (int i = 0; i < size; i++) {
    arr[i] = value;
  }
  return arr;
}
/*
** 为什么不能使用局部数据指针?
** 一旦函数返回，返回的数组地址也就无效了
*/

// int *local_array(int size, int value) {
//   int arr[size];
//   for (int i = 0; i < size; i++) {
//     arr[i] = value;
//   }
//   return arr;
// }

/*
** 传递空指针
** 将指针传递给函数时，使用之前先判断它是否为空是个好习惯。
*/

int *null_pointer(int *arr, int size, int value) {
  if (arr != NULL) {
    for (int i = 0; i < size; i++) {
      arr[i] = value;
    }
  }
  return arr;
}

/*
** 传递指针的指针
** 传递指针目的是修改指针指向的值
** 传递指针的指针目的是修改原指针
*/
void *ppointer(int **arr, int size, int value) {
  *arr = (int *)malloc(size * sizeof(int));
  if (*arr != NULL) {
    for (int i = 0; i < size; i++) {
      *(*arr + 1) = value;
    }
  }
}

/*
** 函数指针：持有函数地址的指针
**
** 声明： void(*f)(void);
**       int (*f)(double);
**       int* (*f)(int,int);
** 函数指针建议使用fptr 做前缀
**
**
*/
// 函数指针的使用
int (*fptr)(int);
int sq(int n) { return n * n; }

// 传递函数指针,将函数作为参数
int add(int n1, int n2) { return n1 + n2; }
int sub(int n1, int n2) { return n1 - n2; }

typedef int (*fptrOpt)(int, int);

int compute(fptrOpt opt, int n1, int n2) { return opt(n1, n2); }

// 返回函数指针，需要声明返回类型为函数指针
fptrOpt select(char opcode) {
  switch (opcode) {
    case '+':
      return add;
    case '-':
      return sub;
    default:
      break;
  }
}

int evaluate(char opcode, int n1, int n2) {
  fptrOpt opt = select(opcode);
  return opt(n1, n2);
}

// 使用函数指针数组，根据条件选择要执行的函数

typedef int (*operations)(int, int);
operations opr[128] = {NULL};
// 或者使用下面的声明
// int (*oper[128])(int , int ) = {NULL};

void test_oper(){
  opr['+'] = add;
  opr['-'] = sub;
}

int test_evaluate(char opcode, int n1, int n2){
  fptrOpt operation;
  operation =  opr[opcode];
  return operation(n1,n2);


}





int main(int argc, char const *argv[]) {
  int n1 = 5;
  int n2 = 10;
  swap_with_pointers(&n1, &n2);
  printf("n1 = %d\n", n1);
  printf("n2 = %d\n", n2);
  printf("-----------------------------\n");
  swap(n1, n2);
  printf("n1 = %d\n", n1);
  printf("n2 = %d\n", n2);
  printf("-----------------------------\n");
  int *vector1 = allocate_array(5, 45);
  for (int i = 0; i < 5; i++) {
    printf("%d\n", vector1[i]);
  }
  free(vector1);
  printf("-----------------------------\n");
  // int *vector2 = local_array(5, 45);
  // for (int i = 0; i < 5; i++) {
  //   printf("%d\n", vector2[i]);
  // }
  printf("-----------------------------\n");
  int *v = NULL;
  ppointer(&v, 5, 45);

  printf("-------函数指针的使用----------------------\n");
  int n = 5;
  fptr = sq;
  printf("%d sq is %d\n ", n, fptr(n));

  printf("-------传递函数指针----------------------\n");
  printf("%d\n", compute(add, 5, 6));
  printf("%d\n", compute(sub, 5, 6));
  printf("-------返回函数指针----------------------\n");
  printf("%d\n", evaluate('+', 5, 6));
  printf("%d\n", evaluate('-', 5, 6));


  // 比较函数指针
  fptrOpt fptr2 = add;
  if (fptr2 == add){
    printf("fptr1 -> add \n");

  }else{
    printf("fptr1 !-> add \n");
  }


  // 转换函数指针


  return 0;
}

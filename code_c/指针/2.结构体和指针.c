#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct person {
  char* first_name;
  char* last_name;
  char* title;
  unsigned int age;
};

typedef struct person Person;

int main(int argc, char const* argv[]) {
  // 结构体指针的使用
  Person* ptrPerson;
  ptrPerson = (Person*)malloc(sizeof(Person));
  ptrPerson->first_name = (char*)malloc(strlen("Emily") + 1);
  strcpy(ptrPerson->first_name, "Emily");
  ptrPerson->age = 23;

  printf("%d\n", sizeof(Person));

  return 0;
}

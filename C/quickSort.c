/******************************************************************************
 *	File Name: quickSort.c
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

#include <stdio.h>
#include <stdlib.h>

void swap(int *p, int *q) 
{
    int tmp = *p;
    *p = *q;
    *q = tmp;
}
int Partition(int *arr, int left, int right)
{
    int i = left - 1;
    int j = left;

    int tmp = arr[right];
    for (; j <= right - 1; j ++) {
        if (arr[j] < tmp) {
            i ++;
            swap(&arr[i], &arr[j]);
        }
    }

    swap(&arr[right], &arr[i + 1]);
    return i + 1;
}

int Partition1(int *arr, int left, int right)
{
    int i = left, j = right;

    int mid = (left + right) / 2;
    int tmp = arr[mid];

    while (i < j) {
        while (arr[j] >= tmp && i < j) {
            j --;
        }
        while (arr[i] <= tmp && i < j) {
            i ++;
        }

        swap(&arr[i], &arr[j]);
    }
    swap(&arr[mid], &arr[i]);
    return i;
}
void QuickSort(int *arr, int left, int right)
{
    int mid;

    if (NULL == arr || left < 0 || right < 0) {
        return;
    }

    if (left < right) {
        mid = Partition1(arr, left, right);
        QuickSort(arr, left, mid - 1);
        QuickSort(arr, mid + 1, right);    
    }
}
int main() 
{
    int arr[8] = {2,2,1,4,2,6,9,5};
    int arrSize = 8;

    QuickSort(arr, 0, arrSize - 1);
    int i;  
    for (i = 0; i < arrSize; i ++) {
        printf("%d\n", arr[i]);
    }
    int a = 1;
    int b = 2;
    swap(&a, &b);
    printf("%d\n", a);
    return 0;
}

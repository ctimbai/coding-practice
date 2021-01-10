/******************************************************************************
 *	File Name: binarySearch.c
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

#include <stdio.h>
#include <stdlib.h>

int BinarySearch(int *arr, int size, int value)
{
    if (NULL == arr || size <= 0) {
        return -1;
    }

    int l = 0, r = size - 1;
    while (l <= r) {
        int mid = l + ((r - l) >> 1); 
        if (value < arr[mid]) {
            r = mid - 1;
        }
        else if (value > arr[mid]) {
            l = mid + 1;
        }
        else {
            return mid;
        }
    }
    return 0;
}
int main()
{
    int arr[5] = {1,3,5,6,8};
    int value = 4;

     printf("%d\n", BinarySearch(arr, 5, value));
    return 0;
}

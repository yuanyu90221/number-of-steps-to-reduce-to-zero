# number of steps to reduce a number to zero

## 題目解讀：

### 題目來源:
[number-of-steps-to-reduce-a-number-to-zero](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/)

### 原文:

Given a non-negative integer num, return the number of steps to reduce it to zero. If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

 
### 解讀：

給定一個正整數 num

依據以下的法則num遞減到 0

1 如果 num 是偶數 則讓 num = num/2

2 如果 num 不是偶數 則讓 num = num -1

寫一個function 計算出所需要的步驟數目

## 初步解法:
### 初步觀察:

首先是需要把每個步驟利用一個counter記錄下來

每做一步 把 counter加一

### 初步設計:

Given an integer num, num > 0

step 1: set a integer count = 0

step 2: check if num == 0 go to step 6

step 3: check if num is even , num = num /2

step 4: check if num is not even, num = num - 1

step 5: count = count + 1 go to step 2 

step 6: return count
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
func numberOfSteps (num int) int {
    count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num -= 1
		}
		count++
	}
	return count
}
```
## 測資的撰寫

```golang
package number_of_steps

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				num: 14,
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				num: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
```
## my self record
[leetcode golang 6th day](https://hackmd.io/I6Np7LmxQFuH8dpFoOZ5Wg?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)
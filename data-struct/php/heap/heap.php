<?php
/**
 * Created by PhpStorm.
 * User: baimayou
 * Date: 19-3-24
 * Time: 下午6:29
 * 堆：
 * 堆是一个完全二叉树；
 * 堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。
 *  堆是由数组组成,且从1开始
 */
//建立堆，向堆中插入元素(从下向上堆化)
function buildHeap(&$arr,$value){
    if(is_null($arr)){
        $arr[0] = null;
        $arr[1] = $value;
    }
    $num = count($arr);
    $arr[$num] = $value;
    $j = $num;
    $i = $j>>1;
    while($value > $arr[$i] && $i>=1){
        temp($arr,$i,$j);
        $j = $i;
        $i = $j>>1;
    }
}
//删除堆顶元素(将堆中的最后一个元素与堆顶元素交换位置,然后从上向下堆化)
function deleteHeap(&$arr){
    if(is_null($arr)){
        return false;
    }
    $num = count($arr);
    $arr[1] = $arr[$num - 1];
    unset($arr[$num-1]);
    $value = $arr[1];
    $i = 1;
    $j = $i * 2;
    while($value < $arr[$j] && $j<$num-1){
        temp($arr,$i,$j);
        $i = $j;
        $j = $i * 2;
    }
}
/*
 基于堆进行排序
  1，建堆：根据数组建堆
      ①从上向下堆化：和建堆的方法类似，将数组中的元素依次插入堆中
      ②从下向上堆化：从数组的 n/2 开始，调整此子树根节点和子节点的位置，是最大值成为根节点；聪明和n/2到1，且每次要调整以该节点为根的整棵树
  2，排序
    堆顶一定是最大或者最小的元素，所以从依次取出堆顶的元素，并调整堆即可
 * */
function sortHeap(&$arr){
    //建堆
    $num = count($arr);
    $p = floor($num/2);
    while($p >= 1){
        $i = $p;
        $j = $p * 2;
        while($j <= $num){
            if(isset($arr[$j+1]) && $arr[$j+1] > $arr[$j] && $arr[$i] < $arr[$j+1]){
                temp($arr,$i,$j+1);
            }elseif($arr[$i] < $arr[$j]){
                temp($arr,$i,$j);
            }
            $i = $j;
            $j = $j*2;
        }
        $p--;
    }
    //排序
    $resArr = [];
    for($m = $num-1;$m>0;$m--){
        $resArr[$m] = $arr[1];

        $arr[1] = $arr[$m];
        unset($arr[$m]);
        $i = 1;
        $j = $i * 2;
        //比较左右两侧的元素，选出最大值交换位置
        while($arr[$i] < ($arr[$j+1]>$arr[$j] ? $arr[$j+1] : $arr[$j]) && $j<=$num-1){
            $n = $arr[$j+1]>$arr[$j] ? $j+1 : $j;
            temp($arr,$i,$n);
            $i = $n;
            $j = $i * 2;
        }
    }
    $arr = $resArr;
}
function temp(&$arr,$i,$j){
    $temp = $arr[$i];
    $arr[$i] = $arr[$j];
    $arr[$j] = $temp;
}
$arr = [null,9,8,100,80,107,36,5,3,4,7,6];
//buildHeap($arr,100);
//var_dump($arr);
//deleteHeap($arr);
sortHeap($arr);
var_dump($arr);

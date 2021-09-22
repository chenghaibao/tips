<?php

//16. 最接近的三数之和
function threeSumClosest($nums, $target) {
    sort($nums);//排序
    $map=[];//哈希表，加快速度
    $res=null;//保存返回结果
    $len=sizeof($nums);
    foreach($nums as $k=>$v){
        if(isset($map[$v])){  //判断是否已经循环过了
            continue;
        }else{
            $map[$v]=1;
        }
        $l=$k+1;//对撞左指针                //左移
        $r=$len-1;       //对撞右指针       //右移
        while($l<$r){                      //左移小于右移
            $sum=$nums[$l]+$nums[$r]+$v;    //左移值加右移值加当前值
            if($res===null){                //结果为空
                $res=$sum;                  //保存结果
            }else{
                $res=abs($res-$target)>abs($sum-$target)?$sum:$res;  //当前移值是否较大
            }
            if($sum==$target){
                return $sum;//相等肯定是最小值，直接返回
            }else if($sum<$target){   //和小于   前移
                $l++;//移动左指针
            }else{                      //和大于  后移
                $r--;//移动右指针
            }
        }
    }
    return $res;
}


//27. 移除元素
function removeElement(&$nums, $val) {
    $j = 0;
    for( $i = 0;$i<count($nums);$i++ ){
        if( $nums[$i] != $val ){
            $nums[$j] = $nums[$i];
            $j++;
        }
    }

    var_dump($j);
    return $j;

}
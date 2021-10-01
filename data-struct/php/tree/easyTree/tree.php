<?php

/**
 *    PHP二叉树算法
 *    Created on 2017-8-7
 *    Author     entner
 *    Email     1185087164@qq.com
 */

/*
    假设我构造一颗如下的二叉树
            A
        B       C
      #   D   #   #
        #   #
*/

error_reporting(E_ALL ^ E_NOTICE);

$data = array(
    0=>'A',
    1=>'B',
    2=>'#',
    3=>'D',
    4=>'#',
    5=>'#',
    6=>'C',
    7=>'#',
    8=>'#',
);

Class BTNode{
    public $data;
    public $lChild;
    public $rChild;

    public function __construct($data = null){
        $this->data = $data;
    }
}

Class BinaryTree{

    public $root;
    public $btData;

    public function __construct($data = null){
        $this->root = null;
        $this->btData = $data;
        //$this->preOrderCreateBT($this->root);
    }

    public function preOrderCreateBT(&$root = null){
        $elem = array_shift($this->btData);    //移除数组头部，并作为结果返回
        if($elem == null){    #判断：当数组为空时
            return ;
        }else if($elem == '#'){    #判断：当数组为无效单元时，该节点是虚节点，退出当前递归，执行下一个递归
            $root = '#';
            return $root;
        }else{
            $root = new BTNode();
            $root->data = $elem;
            $this->preOrderCreateBT($root->lChild);
            $this->preOrderCreateBT($root->rChild);
        }
        return $root;
    }


    /**
     * TODO:先序遍历二叉树
     * @param $tree object 二叉树
     * @param $temp array  二叉树输出节点存储数组
     */
    public function printBTPreOrder($tree,&$temp){
        if($tree != null){
            $temp[] = $tree->data;
            $this->printBTPreOrder($tree->lChild,$temp);
            $this->printBTPreOrder($tree->rChild,$temp);
        }else{
            return ;
        }
        return $temp;
    }

    /**
     * TODO:中序遍历二叉树
     * @param $tree object 二叉树
     * @param $temp array  二叉树输出节点存储数组
     */
    public function printBTInOrder($tree,&$temp){
        if($tree != null){
            $this->printBTInOrder($tree->lChild,$temp);
            $temp[] = $tree->data;
            $this->printBTInOrder($tree->rChild,$temp);
        }else{
            return;
        }
        return $temp;
    }

    /**
     * TODO:中序遍历二叉树
     * @param $tree object 二叉树
     * @param $temp array  二叉树输出节点存储数组
     */
    public function printBTPosOrder($tree,&$temp){
        if($tree != null){
            $this->printBTPosOrder($tree->lChild,$temp);
            $this->printBTPosOrder($tree->rChild,$temp);
            $temp[] = $tree->data;

        }else{
            return;
        }
        return $temp;
    }

    /**
     * TODO:层序遍历二叉树（需要将书中节点放入队中）
     *
     */
    function PrintFromTopToBottom(&$root)
    {
        // write code here
        if($root == null){
            return ;
        }
        $queue = array();
        array_push($queue,$root);

        while(!is_null($node = array_shift($queue))){
            echo $node->data . ' ';
            if($node->left != null){
                array_push($queue,$node->lChild);

            }
            if($node->left != null){
                array_push($queue,$node->rChild);
            }
        }
    }

}


$object = new BinaryTree($data);
$tree = $object->preOrderCreateBT($root);

echo "<pre>";
print_r($tree);


$items = array(
    1 => array('id' => 1, 'pid' => 0, 'name' => '江西省'),
    2 => array('id' => 2, 'pid' => 0, 'name' => '黑龙江省'),
    3 => array('id' => 3, 'pid' => 1, 'name' => '南昌市'),
    4 => array('id' => 4, 'pid' => 2, 'name' => '哈尔滨市'),
    5 => array('id' => 5, 'pid' => 2, 'name' => '鸡西市'),
    6 => array('id' => 6, 'pid' => 4, 'name' => '香坊区'),
    7 => array('id' => 7, 'pid' => 4, 'name' => '南岗区'),
    8 => array('id' => 8, 'pid' => 6, 'name' => '和兴路'),
    9 => array('id' => 9, 'pid' => 7, 'name' => '西大直街'),
    10 => array('id' => 10, 'pid' => 8, 'name' => '东北林业大学'),
    11 => array('id' => 11, 'pid' => 9, 'name' => '哈尔滨工业大学'),
    12 => array('id' => 12, 'pid' => 8, 'name' => '哈尔滨师范大学'),
    13 => array('id' => 13, 'pid' => 1, 'name' => '赣州市'),
    14 => array('id' => 14, 'pid' => 13, 'name' => '赣县'),
    15 => array('id' => 15, 'pid' => 13, 'name' => '于都县'),
    16 => array('id' => 16, 'pid' => 14, 'name' => '茅店镇'),
    17 => array('id' => 17, 'pid' => 14, 'name' => '大田乡'),
    18 => array('id' => 18, 'pid' => 16, 'name' => '义源村'),
    19 => array('id' => 19, 'pid' => 16, 'name' => '上坝村'),
);

/**
 *TODO:通过引用方式实现无限极分类
 *
 */
function tree_Classify1($items){
    $tree=array();
    $packData=array();
    foreach ($items as  $data) {
        $packData[$data['id']] = $data;
    }
    foreach ($packData as $key =>$val){
        if($val['pid']==0){//代表跟节点
            $tree[]=& $packData[$key];
        }else{
            //找到其父类
            $packData[$val['pid']]['son'][]=& $packData[$key];
        }
    }
    return $tree;
}

/**
 *TODO:通过递归方式实现无限极分类
 *
 */
function tree_Classify2($items,$child='_child',$root = 0){
    $tree=array();
    foreach($items as $key=> $val){

        if($val['pid']==0){
            //获取当前$pid所有子类
            unset($items[$key]);
            if(! empty($items)){
                $child=tree_Classify2($items,$child,$val['id']);
                if(!empty($child)){
                    $val['_child']=$child;
                }
            }
            $tree[]=$val;
        }
    }
    return $tree;
}

echo "<pre>";
print_r(tree_Classify1($items,$child='_child',$root=0));


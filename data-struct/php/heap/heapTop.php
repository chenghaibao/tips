<?php
/**
 * Created by PhpStorm.
 * User: KeenSting
 * Date: 2017/12/7
 * Time: 下午5:47
 * Name: 梁小苍
 * Phone: 13126734215
 * QQ: 707719848
 * File Description: 堆排序
 */

//节点类
class Node{
    public $value;  //节点值
    public $left; //左子节点
    public $right; //右子节点
    public $father; //父节点
    public $pos_in_father; //在父节点的位置,左还是右

    public function __construct($num)
    {
        $this->value = $num;
        $this->left = $this->right = $this->father = null;
        $this->pos_in_father = null;//1 左,2右
    }
}

//测试堆排序类
class ATest{
    private $list; //存放节点列表,节点间的关系用引用表示
    private $num;  //节点数
    private $result; //结果集
    public function __construct(array $data)
    {
        $this->list = $this->result = [];
        $this->num = count($data);
        $this->createHeap($data,$this->num);
        $this->initHeap();
    }

    //执行排序
    public function run()
    {
        //交换堆顶和堆尾的节点值,移除尾部节点,加入到输出中
        while(true)
        {
            $size = count($this->list);
            if($size>2)
            {
                array_push($this->result,$this->list[0]->value);//堆顶的数加入结果集
                $tail = array_pop($this->list);
                $this->list[0]->value = $tail->value;//堆尾的数放到堆顶,在此调整
                if($father = &$tail->father)//释放尾部节点,修改父节点中子节点信息,置null
                {
                    if($tail->pos_in_father==1)
                        $father->left = null;
                    else
                        $father->right = null;
                }
                unset($tail);
                $this->nodeAdjust($this->list[0]);

            }else//只有两个元素的时候就释放掉
            {
                array_push($this->result,$this->list[0]->value);
                array_push($this->result,$this->list[1]->value);
                unset($this->list);
                break;
            }
        }
        //打印降序结果
        print_r($this->result);
        //打印升序结果
        print_r(array_reverse($this->result));
    }

    //创建堆结构,结果存放在list中,按从上到下,从左到右排列
    private function createHeap($data,$length)
    {
        //使用array_shift作为队列输出
        $root = new Node($data[0]);
        $queue = [];
        $index = 0;
        array_push($queue,$root);

        while($queue)
        {
            $father = array_shift($queue);
            $index++;
            if($index<$length)//创建左节点
            {
                $node1 = new Node($data[$index]);
                $father->left = $node1;
                $node1->father = $father;
                $node1->pos_in_father = 1;
                array_push($queue,$node1);
                $index++;
            }
            if($index<$length)//创建右节点
            {
                $node2 = new Node($data[$index]);
                $father->right = $node2;
                $node2->father = $father;
                $node2->pos_in_father = 2;
                array_push($queue,$node2);
            }

            array_push($this->list,$father);

        }

    }

    //初始化为大顶堆
    private function initHeap()
    {
        for($i=$this->num-1;$i>=0;$i--)
        {
            $tmp = &$this->list[$i];
            $this->nodeAdjust($tmp);
        }
    }

    //递归调整节点以及其子节点,保证大顶堆的特性
    private function nodeAdjust(&$tmp)
    {
        if(($left = &$tmp->left)!=null)//左子树
        {
            if(($right = &$tmp->right)!=null)//右子树在
            {

                if($right->value > $left->value)//先比较子节点大小
                {
                    if($tmp->value < $right->value)//父节点和右节点交换
                    {
                        $tmp_val = $right->value;
                        $right->value = $tmp->value;
                        $tmp->value = $tmp_val;
                        $this->nodeAdjust($right);
                    }
                }else
                {
                    if($tmp->value < $left->value)//父节点和左节点交换
                    {
                        $tmp_val = $left->value;
                        $left->value = $tmp->value;
                        $tmp->value = $tmp_val;
                        $this->nodeAdjust($left);
                    }
                }
            }else{//没有右子树

                if($tmp->value < $left->value)
                {

                    $tmp_val = $left->value;
                    $left->value = $tmp->value;
                    $tmp->value = $tmp_val;
                    $this->nodeAdjust($left);
                }
            }
        }else
            echo '不存在子树----';

    }
}

$a = new ATest([3,1,9,2,7,6,11,5,4,13]);
$a->run();
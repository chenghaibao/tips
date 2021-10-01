<?php

class Node {
    public $left;
    public $right;
    public $data;

    public function __construct ($data) {
        $this->data = $data;
        $this->left = null;
        $this->right = null;
    }
}

class BinaryTree {
    public $root = null;

    //插入节点
    public function insert ($data) {
        $newNode = new Node($data);
        if ($this->root == null) {
            $this->root = $newNode;
        } else {
            $this->insertNode($this->root, $newNode);
        }
    }

    private function insertNode ($node, $newNode) {
        if ($newNode->data < $node->data) {
            if ($node->left == null) {
                $node->left = $newNode;
            } else {
                $this->insertNode($node->left, $newNode);
            }
        } else {
            if ($node->right == null) {
                $node->right = $newNode;
            } else {
                $this->insertNode($node->right, $newNode);
            }
        }
    }

    //查询
    public function search ($data) {
        return $this->searchNode($this->root, $data);
    }

    private function searchNode ($node, $data) {
        if ($node == null) {
            return false;
        }
        if ($data < $node->data) {
            return $this->searchNode($node->left, $data);
        } else if ($data > $node->data) {
            return $this->searchNode($node->right, $data);
        } else {
            return $node;
        }
    }

    //删除
    public function delete ($data) {
        return $this->deleteNode($this->root, $data);
    }

    private function deleteNode ($node, $data) {
        if ($node == null) {
            return null;
        }
        if ($data < $node->data) {
            $node->left = $this->deleteNode($node->left, $data);
            return $node;
        } else if ($data > $node->data) {
            $node->right = $this->deleteNode($node->right, $data);
            return $node;
        } else {
            if ($node->left == null && $node->right == null) {
                $node = null;
                return $node;//根节点
            } else if ($node->left == null) {
                $node = $node->right;
                return $node;
            } else if ($node->right == null) {
                $node = $node->left;
                return $node;
            } else {
                //找出要删除的节点，用他左边子节点去替换要删除的节点
                $aux = $this->findMinNode($node->right);
                $node->data = $aux->data;
                $node->right = $this->deleteNode($node->right, $aux->data);
                return $node;
            }
        }
    }

    //找出左侧最小节点
    private function findMinNode ($node) {
        if ($node == null) {
            return null;
        }
        while ($node && $node->left != null) {
            $node = $node->left;
        }
        return $node;
    }

    public function min () {
        return $this->minNode($this->root);
    }

    //找出左侧最小节点
    private function minNode ($node) {
        if ($node == null) {
            return null;
        }
        while ($node && $node->left != null) {
            $node = $node->left;
        }
        return $node->data;
    }

    public function max () {
        return $this->maxNode($this->root);
    }

    //找出右侧最大节点值
    private function maxNode ($node) {
        if ($node == null) {
            return null;
        }
        while ($node && $node->right != null) {
            $node = $node->right;
        }
        return $node->data;
    }

    //层数
    public function getHeight ($node) {
        if ($node == null) {
            return null;
        }
        $leftH = $this->getHeight($node->left);
        $rightH = $this->getHeight($node->right);

        return max($leftH, $rightH) + 1;
    }

    //前序遍历
    public function preOrder ($node) {
        if ($node == null) {
            return;
        }
        echo $node->data.'->';
        $this->preOrder($node->left);
        $this->preOrder($node->right);
    }

    //中序遍历
    public function inOrder ($node) {
        if ($node == null) {
            return;
        }
        $this->inOrder($node->left);
        echo $node->data.'->';
        $this->inOrder($node->right);
    }

    //后序遍历
    public function postOrder ($node) {
        if ($node == null) {
            return;
        }
        $this->postOrder($node->left);
        $this->postOrder($node->right);
        echo $node->data.'->';
    }
}

$nodes = [6, 3, 8, 2, 5, 1, 7];
$bTreeObj = new BinaryTree();
foreach ($nodes as $node) {
    $bTreeObj->insert($node);
}

printf("先序遍历:\n");
$bTreeObj->preOrder($bTreeObj->root);
echo PHP_EOL;

printf("中序遍历\n");
$bTreeObj->inOrder($bTreeObj->root);
echo PHP_EOL;

printf("后序遍历\n");
$bTreeObj->postOrder($bTreeObj->root);

print_r($bTreeObj->search(3));

$bTreeObj->delete(3);

printf("删除后中序遍历！\n");
$bTreeObj->inOrder($bTreeObj->root);

printf("最小值: %d\n", $bTreeObj->min());

printf("最大值: %d\n", $bTreeObj->max());

printf("层数: %d\n", $bTreeObj->getHeight($bTreeObj->root));
<?php

class Node {
    public
        $left;
    public
        $right;
    public
        $data;

    public
    function __construct ($data) {
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
            //当添加完一个结点后，如果 (左子树的高度 - 右子树的高度) > 1, 右旋转
            if ($this->height($node->left) - $this->height($node->right) > 1) {
                //如果它的左子树的右子树高度大于它的左子树的高度
                if ($node->left != null && $this->height($node->left->right) > $this->height($node->left->left)) {
                    //LR型：左旋-右旋
                    $this->leftRightRotate($node);
                } else {
                    //LL型：右旋转
                    $this->rightRotate($node);
                }
            }
        } else {
            if ($node->right == null) {
                $node->right = $newNode;
            } else {
                $this->insertNode($node->right, $newNode);
            }
            //当添加完一个结点后，如果: (右子树的高度-左子树的高度) > 1, 左旋转
            if ($this->height($node->right) - $this->height($node->left) > 1) {
                //如果它的右子树的左子树的高度大于它的右子树的右子树的高度
                if ($node->right != null && $this->height($node->right->left) > $this->height($node->right->right)) {
                    //RL型：右旋-左旋
                    $this->rightLeftRotate($node);
                } else {
                    //RR型：左旋转
                    $this->leftRotate($node);
                }
            }
        }
    }

    public function height ($node) {
        if ($node == null) {
            return null;
        }
        $leftHeight = $this->height($node->left);
        $rightHeight = $this->height($node->right);
        return 1 + max($leftHeight, $rightHeight);
    }

    //[RR型：左旋转]
    public function leftRotate (&$node) {
        //创建新的结点，以当前根结点的值
        $newNode = new Node($node->data);

        //把新的结点的左子树设置成当前结点的左子树
        $newNode->left = $node->left;

        //把新的结点的右子树设置成带你过去结点的右子树的左子树
        $newNode->right = $node->right->left;

        //把当前结点的值替换成右子结点的值
        $node->data = $node->right->data;

        //把当前结点的右子树设置成当前结点右子树的右子树
        $node->right = $node->right->right;

        //把当前结点的左子树(左子结点)设置成新的结点
        $node->left = $newNode;
    }

    //[LL型：右旋转]
    public function rightRotate (&$node) {
        //创建新的结点，以当前根结点的值
        $newNode = new Node($node->data);

        //把新的结点的右子树设置成当前结点的右子树
        $newNode->right = $node->right;

        //把新的结点的左子树设置成带你过去结点的左子树的右子树
        $newNode->left = $node->left->right;

        //把当前结点的值替换成左子结点的值
        $node->data = $node->left->data;

        //把当前结点的左子树设置成当前结点左子树的左子树
        $node->left = $node->left->left;

        //把当前结点的右子树(右子结点)设置成新的结点
        $node->right = $newNode;
    }

    //[RL型：右旋-左旋]
    public function rightLeftRotate (&$node) {
        //先对右子结点进行右旋转
        $this->rightRotate($node->right);
        //然后在对当前结点进行左旋转
        $this->leftRotate($node);
    }

    //[LR型：左旋-右旋]
    public function leftRightRotate (&$node) {
        //先对当前结点的左结点(左子树)->左旋转
        $this->leftRotate($node->left);
        //再对当前结点进行右旋转
        $this->rightRotate($node);
    }

    function isBalanced ($node) {
        return abs($this->leftHeight($node) - $this->rightHeight($node)) < 2;
    }

    function rightHeight ($node) {
        if ($node->right == null) {
            return 0;
        }
        return $this->rightHeight($node->right) + 1;
    }

    function leftHeight ($node) {
        if ($node->left == null) {
            return 0;
        }
        return $this->leftHeight($node->left) + 1;
    }

    //查询
    public function search ($data) {
        if ($node == null) {
            return null;
        }
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
        $node = $this->deleteNode($this->root, $data);
    }

    public function deleteNode ($node, $data) {
        if ($node == null) {
            return null;
        }

        if ($data < $node->data) {
            $node->left = $this->deleteNode($node->left, $data);
            if ($this->height($node->right) - $this->height($node->left) > 1) {
                $currentNode = $node->right;
                if ($node->right != null && $this->height($currentNode->left) > $this->height($currentNode->right)) {
                    //RL型：右旋-左旋
                    $currentNode = $this->rightLeftRotate($node);
                } else {
                    //RR型：左旋转
                    $currentNode = $this->leftRotate($node);
                }
            }
        } else if ($data > $node->data) {
            $node->right = $this->deleteNode($node->right, $data);
            if ($this->height($node->left) - $this->height($node->right) > 1) {
                $currentNode = $node->left;
                if ($node->left != null && $this->height($currentNode->right) > $this->height($currentNode->left)) {
                    //LR型：左旋-右旋
                    $currentNode = $this->leftRightRotate($node);
                } else {
                    //LL型：右旋转
                    $currentNode = $this->rightRotate($node);
                }
            }
        } else {
            if ($node->left == null && $node->right == null) {
                $node = null;
            } else if ($node->left == null) {
                $node = $node->right;
            } else if ($node->right == null) {
                $node = $node->left;
            } else {
                //找出要删除的节点，用他左边子节点去替换要删除的节点
                $aux = $this->findMinNode($node->right);
                $node->data = $aux->data;
                $node->right = $this->deleteNode($node->right, $aux->data);
            }
        }
        return $node;
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

$nodes = [10, 20, 7, 17, 18, 22, 23, 24, 25];
$avlObj = new BinaryTree();

foreach ($nodes as $node) {
    $avlObj->insert($node);
}

printf("中序遍历\n");
$avlObj->inOrder($avlObj->root);
echo PHP_EOL;

$avlObj->delete(20);

printf("删除后中序遍历！\n");
$avlObj->inOrder($avlObj->root);

printf("层数: %d\n", $avlObj->rightHeight($avlObj->root));

printf("平衡因子: %d\n", $avlObj->isBalanced($avlObj->root));

printf("右子树高度: %d\n", $avlObj->rightHeight($avlObj->root));

printf("左子树高度: %d\n", $avlObj->leftHeight($avlObj->root));
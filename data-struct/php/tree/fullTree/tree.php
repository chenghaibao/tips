<?php

class Node {
    public $value;
    public $leftNode;
    public $rightNode;
}

/* 找到空节点 */
function findEmpytNode ($node, $parent = null) {
    if (empty($node->value)) {
        return $node;
    } else {
        if (empty($node->leftNode->value)) {
            return $node->leftNode;
        } else if (empty($node->rightNode->value)) {
            return $node->rightNode;
        } else {
            if (empty($parent) || $node->value == $parent->rightNode->value) {
                return findEmpytNode($node->leftNode, $node);
            } else {
                return findEmpytNode($parent->rightNode, $node);
            }
        }
    }
}

/* 添加节点 */
function addNode ($node, $value) {
    $emptyNode = findEmpytNode($node);
    setNode($emptyNode, $value);
}

/* 设置节点 */
function setNode ($node, $value) {
    $node->value = $value;
    $node->leftNode = new Node();
    $node->rightNode = new Node();
}

/* 打印 */
function printTree ($node, $parent = null) {
    if (empty($node->value)) return;
    echo $node->leftNode->value;
    echo $node->rightNode->value;
    if (empty($parent) || $node->value == $parent->rightNode->value) {
        printTree($node->leftNode, $node);
    } else {
        printTree($parent->rightNode, $node);
    }
}

$head = new Node();
setNode($head, 1);
addNode($head, 2);
addNode($head, 3);
addNode($head, 4);
addNode($head, 5);
addNode($head, 6);
printTree($head);


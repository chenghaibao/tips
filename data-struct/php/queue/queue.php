<?php

$_stack = [];

// php内部函数实现方式
/**
 *
 * $q = new SplQueue();
 * $q->push(1);
 * $q->push(2);
 * $q->push(3);
 * $q->pop();
 * print_r($q);
 */


// 方法实现方式

function in (&$_stack, $content) {
    $_stack[] = $content;
}

function out (&$_stack) {
    if ($_stack) {
        foreach ($_stack as $k => $value) {
            unset($_stack[$k]);
            break;
        }
    }
}


in($_stack, 3);
in($_stack, 500);
var_dump($_stack);
out($_stack);
var_dump($_stack);
out($_stack);
var_dump($_stack);
out($_stack);
var_dump($_stack);
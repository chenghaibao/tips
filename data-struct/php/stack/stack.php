<?php

$_stack = [];

// php内部函数实现方式
/**
 *
 * function push (&$_stack, $content) {
 * array_push($_stack, $content);
 * }
 *
 * function pop (&$_stack) {
 * array_pop($_stack);
 * }
 *
 * push($_stack);
 * push($_stack, 500);
 * var_dump($_stack);
 * pop($_stack);
 * var_dump($_stack);
 * pop($_stack);
 * var_dump($_stack);
 */


// 方法实现方式

function push (&$_stack, $content) {
    $_stack[] = $content;
}

function pop (&$_stack) {
    $len = count($_stack)-1;
    if($len >= 0 ){
        unset($_stack[$len]);
    }
}


push($_stack, 3);
push($_stack, 500);
var_dump($_stack);
pop($_stack);
var_dump($_stack);
pop($_stack);
var_dump($_stack);
pop($_stack);
var_dump($_stack);
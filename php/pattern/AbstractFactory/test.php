<?php
/**
 * 创建型模式
 *
 * php单例模式
 *
 */


// 注册自加载
spl_autoload_register('autoload');

function autoload($class)
{
    require dirname($_SERVER['SCRIPT_FILENAME']) . '//..//' . str_replace('\\', '/', $class) . '.php';
}

/************************************* test *************************************/

$animal = new \AbstractFactory\Factory();
$kind = $animal->createPig();
$kind->eat();


<?php
/**
 * 创建型模式
 *
 * php单例模式
 *
 * @author  TIGERB <https://github.com/TIGERB>
 * @example 运行 php test.php
 */


// 注册自加载
spl_autoload_register('autoload');

function autoload($class)
{
    require dirname($_SERVER['SCRIPT_FILENAME']) . '//..//' . str_replace('\\', '/', $class) . '.php';
}

/************************************* test *************************************/


$instance = new \template\SmsOne([
    "mobile" => 13614658086
]);
$instance->send(13614658086);

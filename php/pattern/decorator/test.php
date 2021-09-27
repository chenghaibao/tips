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

$template = new \decorator\CouponMessageTemplate();
$message = new \decorator\Message();

// 老系统，用不着过滤，只有内部用户才看得到
$message->send($template);

// 新系统，面向全网发布的，需要过滤一下内容哦
$message->msgType = 'new';
$template = new \decorator\AdFilterDecoratorMessage($template);
$template = new \decorator\SensitiveFilterDecoratorMessage($template);

// 过滤完了，发送吧
$message->send($template);

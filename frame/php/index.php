<?php

/**
 * 入口文件
 * 1.定义常量
 * 2.加载函数库
 * 3.启动框架
 */

define("HB", __DIR__);
define("CORE", HB.'/core');
define("APP", HB.'/app');
define("MODULE", 'app');
define("DEBUG", true);

if (DEBUG) {
    ini_set("display_errors", 'On');
} else {
    ini_set("display_errors", 'Off');
}

include CORE."/common/function.php";
include CORE."/hb.php";
include "vendor/autoload.php";


spl_autoload_register('\core\hb::load');

\core\hb::run();
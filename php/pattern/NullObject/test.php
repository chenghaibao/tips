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

function autoload ($class) {
    require dirname($_SERVER['SCRIPT_FILENAME']).'//..//'.str_replace('\\', '/', $class).'.php';
}

/************************************* test *************************************/
//创建一个老师：路飞
$teacher = new \nullObject\Teacher('路飞');

// 创建学生
$mc = new \nullObject\Student('麦迪');

// 老师提问
$teacher->doThing($mc);
$teacher->doThing('小李');
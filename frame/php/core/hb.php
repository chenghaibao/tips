<?php

namespace core;

use core\lib\log;
use core\lib\route;

class hb {

    /**
     * @var array
     */
    public static $classMAp = [];

    /**
     * @var array
     */
    public  $assign = [];

    static public function run () {
        log::init();
        $route = new route();
        $ctrl = $route->ctrl;
        $action = $route->action;
        $ctrlFile = APP."/ctrl"."/".$ctrl."Ctrl.php";
        $ctrlClass = '\\'.MODULE."\ctrl\\".$ctrl."Ctrl";
        if (is_file($ctrlFile)){
            include $ctrlFile;
            $class = new $ctrlClass();
//            log::log("ctrl:$ctrlClass    action:$action");
            $class->$action();
        }else{
            throw new \Exception("不存在控制器");
        }
    }

    static public function load ($class) {
        // 自动加载
        if (isset($classMap[$class])) {
            return true;
        } else {
            $class = str_replace("\\", '/', $class);
            $file = HB.'/'.$class.'.php';
            if (is_file($file)) {
                include $file;
                self::$classMAp[$class] = $class;
            } else {
                return false;
            }
        }
    }

    public function assign($name,$value){
            $this->assign[$name] = $value;
    }

    public function display($file){
        $viewFile = APP."/".$file;
        if(is_file($viewFile)){
            extract($this->assign);
            include $viewFile;
        }

    }
}
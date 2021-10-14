<?php

namespace core\lib;

class conf {

    static public $conf = [];

    static public function get ($name, $file) {

        if (isset(self::$conf[$file])) {
            return self::$conf[$file][$name];
        } else {
            $path = HB.'/core/config/'.$file.'.php';
            if (is_file($path)) {
                $conf = include $path;
                if (isset($conf[$name])) {
                    self::$conf[$file] = $conf;
                    return $conf[$name];
                } else {
                    throw new \Exception("找不到配置项");
                }
            } else {
                throw new \Exception("不存在配置文件");

            }
        }

    }

    static public function all ($file) {

        if (isset(self::$conf[$file])) {
            return self::$conf[$file];
        } else {
            $path = HB.'/core/config/'.$file.'.php';
            if (is_file($path)) {
                $conf = include $path;
                self::$conf[$file] = $conf;
                return $conf;
            }
        }
    }
}
<?php

namespace factory;

class Sample {

    public static $instance;


    public function __construct () {

    }

    public static function produce ($type = '') {
        switch ($type) {
            case "Pig":
                self::$instance = new Pig();
                break;
            case "Dog":
                self::$instance = new Dog();
                break;
            default:
                echo "不存在";
                break;
        }

        return self::$instance;
    }

}
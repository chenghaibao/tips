<?php

namespace AbstractFactory;

class Dog implements AnimalInterface{

    public function __construct(){
        echo "生产了一个狗"."\n";
    }

    public function eat () {
        // TODO: Implement eat() method.
        echo "狗吃肉"."\n";
    }
}
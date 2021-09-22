<?php

namespace AbstractFactory;

class Pig implements AnimalInterface{

    public function __construct(){
        echo "生产了一个猪"."\n";
    }

    public function eat () {
        // TODO: Implement eat() method.
        echo "猪吃饲料"."\n";
    }
}
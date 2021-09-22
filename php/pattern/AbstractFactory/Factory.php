<?php

namespace AbstractFactory;

class Factory implements KindInterface {

    public function createDog () {
        // TODO: Implement createDog() method.
        return new Dog();
    }

    public function createPig () {
        // TODO: Implement createPig() method.
        return new Pig();
    }
}
<?php

namespace strategy;


class  UserSelect {

    /** var StrategyInterface */
    public $instance;

    public function __construct (StrategyInterface $strategy) {
        $this->instance = $strategy;
    }

    public function something () {
        $this->instance->doThing();
    }
}
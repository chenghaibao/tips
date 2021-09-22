<?php
namespace state;
class ZooState {

    public $instance;

    private $_currentSeason;
    private $_season = ["Dog", "Pig"];

    function __construct ($season = 'Dog') {
        $this->_currentSeason = $season;
        $this->setState($this->_currentSeason);
    }

    public function setState ($state) {
        switch ($state) {
            case "Dog" :
                $this->instance = new \state\Dog();
                break;
            case "Pig" :
                $this->instance = new \state\Pig();
                break;
            default :
                echo "不存在。\n";
                break;
        }
    }

    public function eat () {
        $this->instance->eat();
    }

    public function drink () {
        $this->instance->drink();
        $this->nextAnimal();
    }

    private function nextAnimal () {
        $nowKey = (int)array_search($this->_currentSeason, $this->_season);
        if ($nowKey < 1) {
            $nextSeason = $this->_season[$nowKey + 1];
        } else {
            $nextSeason = 'Dog';
        }
        $this->_currentSeason = $nextSeason;
        $this->setState($this->_currentSeason);
    }
}
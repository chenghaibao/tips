<?php
namespace observer;

class Observerable implements ObserverableInterface{

    private $_observers = [];

    private $_name = '【被观察者:大萝卜】';

    public function __get($name='')
    {
        return $this->$name;
    }

    public function attach (ObserverInterface $observer) {
        // TODO: Implement attach() method.
        if (!in_array($observer, $this->_observers, true)) {
            $this->_observers[] = $observer;
        }
    }

    public function delete (ObserverInterface $observer) {
        // TODO: Implement delete() method.
        foreach ($this->_observers as $k => $v) {
            if ($v === $observer) {
                unset($this->_observers[$k]);
            }
        }
    }

    public function notify () {
        // TODO: Implement notify() method.
        foreach ($this->_observers as $v) {
            $v->doThing($this);
        }
    }
}
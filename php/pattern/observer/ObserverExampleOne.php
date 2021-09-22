<?php
namespace observer;

class ObserverExampleOne implements ObserverInterface{

    public function doThing (ObserverableInterface $observable) {
        // TODO: Implement doThing() method.
        echo "{$observable->_name}干个吊\n";
    }
}
<?php
namespace observer;

class ObserverExampleTwo implements ObserverInterface{

    public function doThing (ObserverableInterface $observable) {
        // TODO: Implement doThing() method.
        echo "{$observable->_name}干你\n";
    }
}
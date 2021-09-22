<?php
namespace observer;

interface ObserverInterface{
    public function doThing(ObserverableInterface $observable);
}
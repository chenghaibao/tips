<?php
namespace observer;

interface ObserverableInterface{

    public function attach(ObserverInterface $observer);
    public function delete(ObserverInterface $observer);
    public function notify();
}
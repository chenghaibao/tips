<?php
namespace nullObject;

class Student extends Person
{

    function doThing ($person) {
        // TODO: Implement doThing() method.
        echo "老师‘{$person->name}’让学生‘{$this->name}’回答了一道题~ \n";
    }
}
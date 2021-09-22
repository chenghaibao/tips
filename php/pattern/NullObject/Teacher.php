<?php
namespace nullObject;


class Teacher extends Person {

    function doThing($person)
    {
        if (!$person instanceof Student) {
            $person = new NullPerson('');
        }
        $person->doThing($this);
    }
}
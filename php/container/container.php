<?php
// 接口类型
interface D{}
class childD implements D{
public function aa(){
var_dump("test");
}
}
function testD(D $d){
$d->aa();
}
testD(new childD());

// 回调匿名函数类型
function testE(Callable $e, string $data){
$e($data);
}
testE(function($data){
var_dump($data);
}, '回调函数');


interface DemoInterface {
public function tt ();
}

class DemoClass implements DemoInterface {
public function tt () {
var_dump(__METHOD__);
}
public function yy () {}
}

class Container {
public function get (string $class) {
$map = [
DemoInterface::class => DemoClass::class,
];
return new $map[$class];
}
}


$container = new Container();
$class = $container->get(DemoInterface::class);
$class->tt();


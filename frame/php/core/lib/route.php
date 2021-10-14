<?php

namespace core\lib;
class route{

    /**
     * @var mixed|string
     */
    public $ctrl = "index";
    /**
     * @var mixed|string
     */
    public $action = "index";
    public function __construct () {
        /**
         * 1.隐藏index.php 静态缀
         * 2.获取url的参数部分
         * 3.返会对应的控制器和方法
         */
       if(isset($_SERVER['REQUEST_URL']) && $_SERVER['REQUEST_URL']  != "/"){
           $path = $_SERVER['REQUEST_URL'];
           $pathArr = explode("/",trim($path,"/"));
           if(isset($pathArr[0])){
               $this->ctrl = $pathArr[0];
           }
           unset($pathArr[0]);
           if(isset($pathArr[1])){
               $this->action = $pathArr[1];
               unset($pathArr[1]);
           }else{
               $this->action = conf::get("ACTION","route");
           }

           $count = count($pathArr)+2;
           $i = 2;
           while($i<$count){
               if(isset($pathArr[$i+1])){
                   $GET[$pathArr[$i]] = $pathArr[$i+1];
               }
               $i = $i+2;
           }
       }
    }
}
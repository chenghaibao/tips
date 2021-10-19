<?php

namespace app\ctrl;

use core\lib\conf;
use core\lib\model;
use Tools\Tools;

class indexCtrl extends \core\hb {

//    public function index(){
//        p('it is pool');
//        $model = new model();
//        $model->model();
////        $model->query();
//    }

    public function index(){
        // ctrl
        p('it is pool');
        Tools::hi();
        // model
//        $model = new model();
//        $model->model();
 //        $model->query();

        // view
//        $this->assign("data","hello");
//        $this->display("view/index.html");


//        conf
//        $temp = conf::get("CTRL","route");
//        $temp = conf::get("ACTION","route");
    }
}
<?php

namespace template;

abstract class Sms {

    public $_config = [];
    protected $_text = '[xx公司]你好，你的验证码是';

    final function __construct($config=[])
    {
        $this->init($config);
    }

    abstract public function init();
    abstract public function SendSms();

    function makeText()
    {
        $this->_text .= rand(000000, 999999);
    }

    final function send($mobile=0)
    {
        // 生成文本
        $this->makeText();
        // 发送短信
        $this->sendSms($mobile);
    }
}
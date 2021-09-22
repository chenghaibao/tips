<?php
namespace template;

class SmsTwo extends Sms{

    public function init ($config = []) {
        // TODO: Implement init() method.
        $this->_config =$config;
    }

    public function SendSms ($mobile=0) {
        // TODO: Implement SendSms() method.
        echo "厂商‘two’给手机号{$mobile}发送了短信：{$this->_text} \n";
    }
}
<?php
namespace decorator;
// 客户端，发送接口，需要使用模板来进行短信发送
class Message
{
    public $msgType = 'old';
    public function send(MessageTemplate $mt)
    {
        // 发送出去咯
        if ($this->msgType == 'old') {
            echo '面向内网用户发送' . $mt->message() . PHP_EOL;
        } else if ($this->msgType == 'new') {
            echo '面向全网用户发送' . $mt->message() . PHP_EOL;
        }

    }
}

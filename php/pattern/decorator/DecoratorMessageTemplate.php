<?php
namespace decorator;
// 我们来准备好装饰上面那个过时的短信模板
abstract class DecoratorMessageTemplate implements MessageTemplate
{
    public $template;
    public function __construct($template)
    {
        $this->template = $template;
    }
}

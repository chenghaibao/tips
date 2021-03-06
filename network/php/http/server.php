<?php
/**
 * @socket 通信的整个过程
 * @socket_create   //创建套接字
 * @socket_bind     //绑定IP和端口
 * @socket_listen   //监听相应端口
 * @socket_accept   //接收请求
 * @socket_read     //获取请求内容
 * @socket_write    //返回数据
 * @socket_close    //关闭连接
 */
class MyServer{
    private $ip;
    private $port;
    private $webroot;
    //将常用的MIME类型保存在一个数组中
    private $contentType=array(
        ".html"=>"text/html",
        ".htm"=>"text/html",
        ".xhtml"=>"text/html",
        ".xml"=>"text/html",
        ".php"=>"text/html",
        ".java"=>"text/html",
        ".jsp"=>"text/html",
        ".css"=>"text/css",
        ".ico"=>"image/x-icon",
        ".jpg"=>"application/x-jpg",
        ".jpeg"=>"image/jpeg",
        ".png"=>"application/x-png",
        ".gif"=>"image/gif",
        ".pdf"=>"application/pdf",
    );
    public function __construct($ip="127.0.0.1",$port=65500){
        set_time_limit(0);
        $this->ip=$ip;
        $this->port=$port;
        $this->webroot=__DIR__.'/www';
        echo "\nServer init sucess\n";
    }
    public function listen(){
        $socket=socket_create(AF_INET,SOCK_STREAM,SOL_TCP);
        if(!$socket)
            echo "CREATE ERROR:".socket_strerror(socket_last_error()).'\n';
        $bool=socket_bind($socket,$this->ip,$this->port);
        if(!$bool)
            echo "BIND ERROR:".socket_strerror(socket_last_error()).'\n';
        while(true){
            $bool=socket_listen($socket);
            if(!$bool)
                echo "LISTEN ERROR:".socket_strerror(socket_last_error()).'\n';
            $new_socket=socket_accept($socket);
            if(!$new_socket)
                echo "ACCPET ERROR:".socket_strerror(socket_last_error()).'\n';
            $string=socket_read($new_socket,20480);
            $data=$this->request($string);
            $num=socket_write($new_socket,$data);
            if($num==0)
                echo "WRITE ERROR:".socket_strerror(socket_last_error())."\n";
            else
                echo "request already succeed\n";
            socket_close($new_socket);
        }
    }
    /**
     * [读取get或post请求中的url，返回相应的文件]
     * @param  [string]
     * @return [string]
     * http头
     * method url protocols
     */
    public function request($string){
        echo $string;
        $pattern="/\s+/";

        $request=preg_split($pattern,$string);
        if(count($request)<3)
            return "request error\n";
        $filename=$this->webroot.$request[1];
        echo "filename:".$filename."\n";
        $type=$this->setContentType($filename);
        if(file_exists($filename)){
            $data=file_get_contents($filename);
            return $this->addHeader($request[2],200,"OK",$data,$type);
        }
        else{
            $data="this resource is not exists";
            return $this->addHeader($request[2],1000,"not exists",$data,$type);
        }
    }
    private function addHeader($protocol,$state,$desc,$str,$type){
        return "{$protocol} {$state} {$desc}\r\nContent-type:{$type}\r\n"."Content-Length:".
            strlen($str)."\r\nServer:MyServer\r\n\r\n".$str;
    }
    private function setContentType($filename){
        $type=substr($filename,strpos($filename,'.'));
        if(isset($this->contentType[$type]))
            return $this->contentType[$type];
        else
            return "text/json";
    }
}
$server=new MyServer();
$server->listen();
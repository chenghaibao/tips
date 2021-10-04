//客户端代码：
<?php
//error_reporting(E_ALL | E_STRICT);
set_time_limit(0);
ob_implicit_flush();

$sock = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);

$msg = "ping";
$len = strlen($msg);
socket_sendto($sock, $msg, $len, 0, "127.0.0.1", 8012);

$from_remote = "127.0.0.1";
$port_remote = 8012;
socket_recvfrom($sock, $buf, 65535, 0, $from_remote, $port_remote);
echo "$buf".PHP_EOL;


socket_close($sock);

?>
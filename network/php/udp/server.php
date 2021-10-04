//服务端
<?php

error_reporting(E_ALL | E_STRICT);
set_time_limit(0);
ob_implicit_flush();

$socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);

if ($socket === false) {
    echo "socket_create() failed:reason:".socket_strerror(socket_last_error())."\n";
}

$ok = socket_bind($socket, '127.0.0.1', 8012);

if ($ok === false) {
    echo "socket_bind() failed:reason:".socket_strerror(socket_last_error($socket));
}

while (true) {
    $from_remote = '';
    $port_remote = 0;
    $num = socket_recvfrom($socket, $buf, 65535, 0, $from_remote, $port_remote);
    if ($num === false) {
        break;
    }
    echo "Received $buf from remote address $from_remote and remote port $port_remote".PHP_EOL;
    //sleep(1000);

    $msg = random_int(4, 100);
    //$msg = "Pong!";
    $msg = "er".$msg;
    $len = strlen($msg);
    socket_sendto($socket, $msg, $len, 0, $from_remote, $port_remote);

}

socket_close($socket);
?>
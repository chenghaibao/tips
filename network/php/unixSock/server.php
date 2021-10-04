<?php
$listen_socket = socket_create( AF_UNIX, SOCK_STREAM, 0 );
// 这会儿就不需要把socket bind到IP+PORT上了
// 而是bind到一个sock文件上
$file = "./server.sock";
socket_bind( $listen_socket, $file );
socket_listen( $listen_socket );
while ( true ) {
    $connection_socket = socket_accept( $listen_socket );
    $ret = socket_recv( $connection_socket, $recv_content, 2048, 0 );
    echo $recv_content;
    $encode_ret = "higood";
    socket_write( $connection_socket, $encode_ret, strlen( $encode_ret ) );
}
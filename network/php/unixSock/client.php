<?php
$conn_socket = socket_create( AF_UNIX, SOCK_STREAM, 0 );
$file = "./server.sock";
socket_connect( $conn_socket, $file );
socket_write( $conn_socket, "HI,I am coming." );
socket_recv( $conn_socket, $recv_content, 2048, 0 );
echo $recv_content;
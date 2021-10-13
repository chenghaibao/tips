<?php
$conn_socket = socket_create( AF_UNIX, SOCK_STREAM, 0 );
$file = "/Users/chenghaibao/Downloads/go/tips/app_nodes/master-cYmpkBCuBXM3wnMA.sock";
socket_connect( $conn_socket, $file );
socket_write( $conn_socket, "HI,I am coming." );
while(true){
    socket_recv($conn_socket, $recv_content, 2048, 0 );
    echo $recv_content;
    if($recv_content == "海"){
       socket_write( $conn_socket, "HI,I am coming." );
    }

    //fork子进程 逻辑   (没测试  )
     if($recv_content['type'] ?? "" == "create" ){
           $pid = pcntl_fork();
           $conn_socket = socket_create( AF_UNIX, SOCK_STREAM, 0 );
           socket_connect( $conn_socket, $recv_content['unix']);
           socket_write($conn_socket, "fork-success{$pid}");
     }

}

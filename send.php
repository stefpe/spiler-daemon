<?php


$addr = 'tcp://127.0.0.1:9001';
$fp = @stream_socket_client($addr);
fwrite($fp, "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy");
fclose($fp);

?>


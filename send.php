<?php


$addr = 'tcp://127.0.0.1:9001';
$fp = @stream_socket_client($addr);
fwrite($fp, '{"key1":"value1", "key2":"value2"}');
fclose($fp);

?>

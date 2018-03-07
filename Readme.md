# SP⚡LER PROFILER DAEMON

## installation
make install

## run the daemon
make run endpoint=127.0.0.1:9001 address=127.0.0.1:8000

## build an executable for your platform
make build

### configuration options of the daemon
-endpoint -> bind address of the daemon
-address -> server address which accepts the data (spiler portal address)

### test with netcat ubuntu package
nc 127.0.0.1 9001

### test with php file (client and server)
php send.php -> sends data to 127.0.0.1:9000
php -S localhost:8000 -> runs local php server which accepts the data at: localhost:8000/server.php

Simple TCP based Web Server in go for listening to multiple connections along with graceful shutdown

cmd:-`go run main.go -p :<port-name> -c<connection-method-name>`

The connection methods can be tcp/tcp6/tcp4/unix/unixpacket

Test commands:-
```
curl http://localhost:8000 & curl http://localhost:8000
```

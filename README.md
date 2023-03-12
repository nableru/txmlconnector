# txmlconnector

### Description

[gRPC](https://grpc.io/) interface around 
[TransaqConnector](https://www.finam.ru/howtotrade/tconnector/)
to be able to connect from different languages via TCP (remote procedure call) and linux as well 

#### Service proto description

```protobuf
syntax = "proto3";

option go_package = "github.com/lebedev_k/txmlconnector";

package transaqConnector;

message DataRequest {
}

message DataResponse {
  string message = 1;
}

message SendCommandRequest {
  string message = 1;
}

message SendCommandResponse {
  string message = 1;
}

service ConnectService {
  rpc FetchResponseData(DataRequest) returns (stream DataResponse) {}
  rpc SendCommand(SendCommandRequest) returns (SendCommandResponse) {}
}
```

### Starting server in docker
```
docker run kmlebedev/txmlconnector server
```

### Start server and client
```
docker-compose -f docker/txmlconnector-compose.yaml -p txmlconnector up
```

### Starting server in Linux/MacOSX

#### Requirements
##### MacOSX

```shell
brew install mingw-w64 
brew cask install wine-stable
```
##### Debian/Ubintu

```shell
sudo apt install wine64
```

#### Set environment variables

```shell
export TC_LOG_LEVEL=debug
export TC_DLL_PATH=txmlconnector64-6.19.2.21.21.dll
```

#### Example output:

```bash
make server
CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/server.exe server/main.go
wine64 bin/server.exe
00ea:fixme:process:SetProcessPriorityBoost (FFFFFFFFFFFFFFFF,1): stub
time="2020-12-17T20:57:01+05:00" level=info msg="Initialize txmlconnector"
InitCrashHandler: Z:\Users\kmlebedev\go\src\txmlconnector\bin\server-201217-205701.mdmp
00f0:fixme:ver:GetCurrentPackageId (0x29d5fd90 0x0): stub
time="2020-12-17T20:57:01+05:00" level=info msg="Server running ..."
00ea:fixme:winsock:set_dont_fragment IP_DONTFRAGMENT for IPv4 not supported in this platform
00ea:fixme:winsock:set_dont_fragment IP_DONTFRAGMENT for IPv6 not supported in this platform
time="2020-12-17T20:57:01+05:00" level=info msg="Press CRTL+C to stop the server..."
```

### Starting client in Linux/MacOSX

#### Set environment variables
[TRANSAQ Connector request access to demo account](https://www.finam.ru/howtotrade/tconnector00002/?program=Transaq%20Connector)

```
export TC_LOGIN="TCNN9979"
export TC_PASSWORD="n3Z4W4"
export TC_HOST="tr1-demo5.finam.ru"
export TC_PORT="3939"
export TC_LOG_LEVEL="DEBUG"
```

#### Example output:
```bash
make client
go build -ldflags "-s -w" -o bin/client client/main.go
bin/client
INFO[0000] Client running ...                           
INFO[0001] res <result success="true"/> 
```

### Example apps
#### Clickhouse-exporter
Historical Moscow Exchange stock quote prices exporting to [clickhouse server](https://clickhouse.tech/) through transaq 
#### Starting
```
docker-compose -f docker/transaq-clickhouse-exporter-compose.yaml -p transaq up
```

### Links
 - [TXmlConnector last(2.21.21) version](https://files.comon.ru/usercontent/TXmlConnector.zip)
 - [TRANSAQ Connector Forum](http://www.transaq.ru/forum/index.php?board=6.0)
 - [TransaqConnector Linux Finam Forum](https://forum.finam.ru/posts/t109457-TransaqConnector-Linux)
 - [TransaqGrpcWrapper](https://github.com/ivanantipin/transaqgrpc)
 - [Creating a simplegRPC client and server application with Golang](http://www.inanzzz.com/index.php/post/fczr/creating-a-simple-grpc-client-and-server-application-with-golang)
 - [Using Go to call Windows API](https://medium.com/@justen.walker/breaking-all-the-rules-using-go-to-call-windows-api-2cbfd8c79724)
 - [How to Set Up gRPC Server-Side Streaming with Go](https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/)
 - [BB](https://www.investopedia.com/terms/b/bollingerbands.asp)
 - https://iqsignal.net/coint
 - https://www.gonum.org/post/intro_to_gonum/
 
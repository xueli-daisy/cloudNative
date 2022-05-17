1. modified simple_server.go
2. add metircs/metrics.go 
3. access http://IP:Port/images
4. access metrics from http://IP:Port/metrics
![image](https://user-images.githubusercontent.com/60275859/164990520-6869c979-577b-4c09-a8a0-8c995a142600.png)
6. build new image by using 'sudo docker build HTTPserver -t xueli17/httpserver:v1.0-metrics'
7. encounter error "go get: no install location for directory /build outside GOPATH
	For more details see: 'go help gopath'"
  when using Dockerfile
  Solve this problem by executing " GO111MODULE=on go mod vendor" in project folder "/home/palden/cloudnative/HTTPserver" before go build
  ![image](https://user-images.githubusercontent.com/60275859/165399320-2b339f0a-5b19-4577-8e3a-1dba88a8e606.png)
  




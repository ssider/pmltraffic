# Docker Env 

## HTTP/1.1

### server

- 默认启动

```
$ docker run -p 8080:80 core.harbor.pml.com.cn/library/traffic:v1.0 http-server
 listening on 80
```

- 指定端口

```
$ docker run -p 8080:8000 core.harbor.pml.com.cn/library/traffic:v1.0 http-server -port 8000
 listening on 8000
```

### client

- 访问服务端

```
$ docker run  -it core.harbor.pml.com.cn/library/traffic:v1.0 curl 192.168.1.142:8080
protocol: HTTP/1.1 
time: 2019-04-08 02:14:00.146 
add return pod name
envs:
  POD_NAME:
  POD_NAMESPACE:
```

- 服务端延迟

```
$ docker run  -it core.harbor.pml.com.cn/library/traffic:v1.0 bash -c "time curl 192.168.1.142:8080?delay=1000"
protocol: HTTP/1.1
time: 2019-04-08 02:12:18.710
envs:
  POD_NAME:
  POD_NAMESPACE:

real    0m1.010s
user    0m0.000s
sys     0m0.000s
```

## HTTP/2.0 (h2c)

### server

```
$ docker run -p 8080:80 core.harbor.pml.com.cn/library/traffic:v1.0 h2c-server
 listening on 80
```

### client

```
$ docker run -it core.harbor.pml.com.cn/library/traffic:v1.0 nghttp http://192.168.1.142:8080
protocol: HTTP/2.0
time: 2019-04-08 02:17:30.456
envs:
  POD_NAME:
  POD_NAMESPACE:
```

##  GRPC

### server

- 默认启动

```
$ docker run -p 9090:9090 core.harbor.pml.com.cn/library/traffic:v1.0 grpc-server
 listening on 9090
```

- 指定端口

```
$ docker run -p 9090:8080 core.harbor.pml.com.cn/library/traffic:v1.0 grpc-server -port 8080
 listening on 8080
```


### client

- 访问服务端

```
$ docker run -it core.harbor.pml.com.cn/library/traffic:v1.0 grpc-client -addr 192.168.1.142:9090
2019/04/08 02:24:49 connect to 192.168.1.142:9090
client:
time: 2019-04-08 02:24:49.899
envs:
  POD_NAME:
  POD_NAMESPACE:
```

- 服务端延迟

```
$ docker run -it core.harbor.pml.com.cn/library/traffic:v1.0 bash -c "time grpc-client -addr 192.168.1.142:9090 -delay 1000"
2019/04/08 02:26:02 connect to 192.168.1.142:9090
client:
time: 2019-04-08 02:26:03.674
envs:
  POD_NAME:
  POD_NAMESPACE:


real    0m1.010s
user    0m0.000s
sys     0m0.000s
```

# Kubernetes Env

## 部署

```
kubectl apply -f config/
```

## HTTP/1.1 访问

```
$ kubectl exec -it $(kubectl get pod -l app=test-client -o jsonpath='{.items[0].metadata.name}') -c client -- curl test-http
protocol: HTTP/1.1
time: 2019-04-08 03:50:40.608 
envs:
    POD_NAME: test-http-5bf668cd95-h5wws
    POD_NAMESPACE: demo
```

## HTTP/2.0 访问

```
$ kubectl exec -it $(kubectl get pod -l app=test-client -o jsonpath='{.items[0].metadata.name}') -c client -- nghttp  http://test-h2c 
protocol: HTTP/2.0
time: 2019-04-08 07:01:56.303 
envs:
    POD_NAME: test-h2c-7c7f8669b7-jdsk8
    POD_NAMESPACE: default
```

## GRPC 访问

```
$ kubectl exec -it $(kubectl get pod -l app=test-client -o jsonpath='{.items[0].metadata.name}') -c client -- grpc-client -addr test-grpc:9090
2019/04/08 07:02:41 connect to test-grpc:9090
client: test-client-7f87987f77-ctvcd 
time: 2019-04-08 07:02:41.584 
envs:
    POD_NAME: test-grpc-7fd68db664-r2n29
    POD_NAMESPACE: default
```
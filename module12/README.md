1. Install Istio
```
   curl -L https://istio.io/downloadIstio | sh -

   Istio has been successfully downloaded into the istio-1.13.3 folder on your system.

   istioctl install --set profile=demo -y

   ✔ Istio core installed
   ✔ Istiod installed
   ✔ Egress gateways installed
   ✔ Ingress gateways installed
   ✔ Installation complete Making this installation the default for injection and validation.

   Thank you for installing Istio 1.13.

   kubectl get pods -n istio-system
   NAME                                    READY   STATUS    RESTARTS   AGE
   istio-egressgateway-5f686bbdfd-vczvt    1/1     Running   0          3m22s
   istio-ingressgateway-699d557cd4-fsljq   1/1     Running   0          3m22s
   istiod-5458bd7bc4-zl6gg                 1/1     Running   0          3m34s
```
2.  Create new namespace
```
kubectl create namespace httpserver
namespace/httpserver created
kubectl label ns httpserver istio-injection=enabled
namespace/httpserver labeled
kubectl get ns -L istio-injection
NAME              STATUS   AGE     ISTIO-INJECTION
cert-manager      Active   29d
cloudnative       Active   37d
default           Active   78d
httpserver        Active   3m49s   enabled
```
3. Install gateway
```
kubectl apply -f gateway.yaml -n httpserver                                                        
gateway.networking.istio.io/httpserver created

kubectl get gw -n httpserver
NAME         AGE
httpserver   3m47s
```
4. Install virtualService
```
kubectl apply -f virtualservice.yaml -n httpserver
virtualservice.networking.istio.io/httpserver created

kubectl get vs -n httpserver
NAME         GATEWAYS         HOSTS   AGE
httpserver   ["httpserver"]   ["*"]   43s
```
5. Deploy httpserver pod
```
kubectl apply -f httpserver-deployment.yaml -n httpserver
deployment.apps/httpserver created
kubectl get pods -n httpserver
NAME                         READY   STATUS    RESTARTS   AGE
httpserver-bddd8b95f-b5srr   2/2     Running   0          19s
```
6. Deploy httpserver service
```
kubectl apply -f svc.yaml -n httpserver
service/httpserver created

kubectl get svc -n httpserver
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpserver   ClusterIP   10.107.165.24   <none>        80/TCP    22m
curl 10.107.165.24/healthz
working
```
7. 
```
kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      10.98.193.125   <none>        80/TCP,443/TCP                                                               45h
istio-ingressgateway   LoadBalancer   10.109.60.29    <pending>     15021:32532/TCP,80:30427/TCP,443:30777/TCP,31400:30609/TCP,15443:32442/TCP   45h
istiod                 ClusterIP      10.105.235.48   <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        45h

curl 10.109.60.29/healthz
working
```





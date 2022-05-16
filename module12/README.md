1. Install Istio
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

2. 

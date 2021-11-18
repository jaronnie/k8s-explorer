# k8s-image-pull-policy

k8s 拉取镜像策略之 always 是如何实现的

## 探究过程

```shell
docker build -t "gocloudcoder/kube-image-pull:develop" .
docker push gocloudcoder/kube-image-pull:develop
kubectl apply -f deploy-service.yaml
```

**docker images**

![image-20211119013519266](https://resource.gocloudcoder.com/image-20211119013519266.png)

**docker ps | grep kube-image-pull**

![image-20211119013620285](https://resource.gocloudcoder.com/image-20211119013620285.png)

重新修改镜像，上传到 dockhub

修改 main.go 中打印的内容为 Hello World1

```shell
docker rmi gocloudcoder/kube-image-pull
docker build -t "gocloudcoder/kube-image-pull:develop" .
docker push gocloudcoder/kube-image-pull:develop
```

**docker images**

![image-20211119014513818](https://resource.gocloudcoder.com/image-20211119014513818.png)

发现端倪了吧

原来的镜像 TAG 变成了 none, 新更新的镜像是最新的 develop tag

这就是 k8s image pull 策略之 always 原理！

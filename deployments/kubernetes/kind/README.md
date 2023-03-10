```bash
# 列出所有的 k8s 集群
kind get clusters

# 创建
kind create cluster --config kind-example-config.yaml

# 删除
kind delete cluster -n my-k8s
```

## 存储
[CSI Hostpath Driver](https://github.com/kubernetes-csi/csi-driver-host-path)：一个示例CSI Driver(非生产环境)，它在单个节点上创建一个本地目录作为卷。

[Local Path Provisioner](https://github.com/rancher/local-path-provisioner):使用 Kubernetes 动态配置持久性本地存储


[Open-Local](https://github.com/alibaba/open-local)：Open-Local是由多个组件构成的本地磁盘管理系统，目标是解决当前 Kubernetes 本地存储能力缺失问题。通过Open-Local，使用本地存储会像集中式存储一样简单。

[NFS CSI driver for Kubernetes](https://github.com/kubernetes-csi/csi-driver-nfs)：该驱动程序用于Kubernetes访问Linux节点上的NFS服务器。



docker inspect kind-control-plane | grep IPAddress



---
[Quick start](https://kind.sigs.k8s.io/docs/user/quick-start/)
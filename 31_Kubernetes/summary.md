# Resume Kurbentes

1. Kubernetes (K8s) adalah sistem sumber terbuka untuk mengotomatiskan penerapan, penskalaan, dan pengelolaan aplikasi dalam container

2. What in can do
<br>\- service discovery and load balancing
<br>\- horizontal scaling
<br>\- automated rollouts and rollbacks.
<br>\- secret and configuration management
<br>What Kurbenetes is not
<br>\-Does not limit the types of applications supported
<br>\-Does not deploy source code and does not build your application 
<br>\-Does not provide application-level services, such as middleware
<br>\-Does not dictate logging, monitoring, or alerting solutions.

3. Kurbenetes Workloads
* Namespace : provide for a scope of Kubernetes resource, carving up your cluster in smaller units.
* Pod : is a collectin of containers shaaring a network and mount namespace and is the basic unit of deployment in Kubernetes
* Labels : the mechanism you use to organize Kubernetes objects
* Deploment : a supervisor for pods, giving you fine-grained control over how and when a new pod version is rolled out as well as rolled back to a previous state
* Rolling updates : allow Deployments update to take place with zer downtime by incrementally updating Pods instances with new ones. 
* Secret : porvide you with a mechanism to use such information in a safe and reliable way.
* Service : is an abstractiono for pods, providing a stable, so called virtual IP (VIP) address.
* Ingress : API object that manages external access to the services in a cluster, typically HTTP

# Langkah-langkah

* Buat akun okteto
* Download kube-config
* Update `~/.kube/config` atau `C:\\Users\<user-name>\.kube\config`
    - Download kubectl
    - Bikin akun okteto
    - Hubungkan kubectl ke cluster okteto dengan cara:
        - Menjalankan ini di terminal:
            - Windows: `$Env:KUBECONFIG=("c:\\User\<user-name>\Downloads\okteto-kube.config;$Env:KUBECONFIG;$HOME\.kube\config")`
            - Linux: `export KUBECONFIG=/mnt/c/Users/<user-name>/Downloads/okteto-kube.config:$KUBECONFIG:$HOME/.kube/config`
        - Supaya tidak perlu menjalankan perintah di atas setiap kali hendak manage okteto:
        Tambahkan script di atas ke `~/.bashrc` atau `~/.zshrc` (tergantung terminal yg dipakai)
        - Jika yang di atas gagal, salin isi okteto-kube.config ke `~/.kube/config`
* Pastikan `kubectl config get-contexts` mengarah ke okteto, jika tidak lakukan `kubectl config use-context <context-name>`
* Deploy mysql (`./mysql/deploy.sh`)
* Deploy go app (`./go-app/deploy.sh`)


# Catatan

* Cloud provider lain, seperti AWS dan GCP memiliki mekanisme sendiri untuk update kubeconfig. Biasanya harus menginstall CLI tool yang mereka sediakan
* Untuk MySQL sebaiknya menggunakan `statefulSet`, bukan `deployment`, namun dalam contoh ini kita menggunakan `deployment` saja, karena okteto free plan tidak mengijinkan akses `volume` API
* Supaya service bisa diakses dari luar cluster, kita perlu menggunakan `ingress`

* Okteto free plan tidak mengijinkan membuat volume dan volume claim, jadi semua bagian yang terkait dengan storage tidak bisa dipakai

sumber: https://github.com/goFrendiAsgard/alta-batch-3/blob/master/14-kubernetes/README.md
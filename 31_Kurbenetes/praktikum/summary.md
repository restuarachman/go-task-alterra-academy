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
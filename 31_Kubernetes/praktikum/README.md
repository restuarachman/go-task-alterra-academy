# Task Deployment Kubernetes
1. Deploy you application by using kubectl from local machine
    * Containerize go app (docker build)
    ![](./1.jpg)
    
    * push to image registry (docker push)
    ![](./2.jpg)

    * Create kubernetes manifest
    ![](./3.jpg)

    * Try to perform kubectl apply
    ![](./4.1.jpg)
    ![](./4.2.jpg)


# Result 
1. Deploy MYSQL
![](./deploy-mysql.jpg)
![](./okteto-mysql.jpg)

2. Deploy go-app
![](./deploy-app.jpg)
![](./okteto-app.jpg)

3. Test app on Postman
    * Checking deployment name `kubectl get deployment`
    ![](./check-deployment.jpg)

    * Forward a local port to a port on the Deployment (Port Foward) `kubectl port-forward deployment/go-app 8000:8000`
    ![](./port-forward.jpg)

    * Test on postman using port forward (`localhost:8000`)
    ![](./postman.jpg)


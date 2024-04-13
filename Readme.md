# 1. Mysqlを用意する

- namespace: database を作成する
- mysql-deployment.yaml で mysql をデプロイする(5.7だとエラーが出るので8.0以上を使う)
- mysql-service.yaml で mysql-service を作成する

- コマンド集
    ```
    kubectl apply -f database-ns.yaml -n database
    kubectl apply -f mysql-deployment.yaml -n database
    kubectl apply -f mysql-service.yaml -n database

    kubectl get ns -n database
    kubectl get pod -n database
    kubectl get svc -n database

    kubectl exec -n database -it mysql-deployment-<ID> bash
    ```

# 2. sample-appを用意する
- namespace=sample (terminalで実行)
- namespace: sample を作成する

    ``` 
    kubectl create namespace $namespace
    ```

- sample-app-deployment.yaml で sample-app をデプロイする

    ```
    kubectl create deploy sample-app --image=ghcr.io/nakamasato/fastapi-sample:v1.0 --dry-run=client -o yaml > sample-app-deployment.yaml
    ```

- env.testからConfigmapを作成する

    ```
    kubectl create cm sample-app --from-env-file=env.txt --dry-run=client -o yaml > sample-app-configmap.yaml
    ```

- Secret を作成する

    ```
    kubectl create secret generic sample-app --from-literal=MYSQL_PASSWORD=password --dry-run=client -o yaml > sample-app-secret.yaml
    ``` 

- serviceを作成する

    ```
    kubectl create service clusterip sample-app --tcp=80 --dry-run=client -o yaml > sample-app-service.yaml
    ```

- portforwardする

    ```
    kubectl port-forward svc/sample-app 8080:80 -n $namespace
    ```
- I AM user を用意（例:eks-setup-user）
- 上記IAMのprofileを用意
    ```
    aws configure --profile eks-setup-user
    ```

- eksctlをインストール

    ```
    curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp

    sudo mv /tmp/eksctl /usr/local/bin

    eksctl version
    ```

- EKS clusterを作成　(20分ほどかかる)
    ```
    eksctl create cluster --name test-cluster --region ap-northeast-1 --profile eks-setup-user
    ```

- eksの動作確認
- kubeconfigを作成
    ```
    aws eks update-kubeconfig --name test-cluster --profile eks-setup-user
    ```

- nginxのpodを用意
    ```
    kubectl create deploy nginx --image=nginx --replicas=3

    kubectl get pods
    ```

- serviceを用意
    ```
    kubectl create service loadbalancer nginx --tcp=80 (--dry-run=client -o yaml)

    kubectl get svc
    ```

- オブジェクトを削除
    ```
    kubectl delete svc nginx
    kubectl delete deploy nginx
    ```

- EKS clusterを削除
    ```
    eksctl delete cluster --name=test-cluster --region ap-northeast-1 --profile eks-setup-user
    ```
- RUNはimage build時に実行される
- CMDはコンテナ起動時に実行される。コンテナ起動時のプロセスを指定できる。
- ENTRYPOINTもコンテナ起動時のプロセスを指定できる。
- 例えば下記のようにした場合goコマンドを引数で渡さずにgo を実行できる 
    ```
    ENTRYPOINT ["go"]

    docner container run hoge/go:latest version // ENTRYPOINTでgoをしているのでgoコマンドは不要
    ```
- ENTRYPOINTはイメージの作成者側でコンテナの用途をある程度制限したい場合に活用できる
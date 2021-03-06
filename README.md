# BbMatching
野球チームのマッチングアプリケーション

### 使用技術・フレームワーク
* MySQL 5.7.26
* Go 1.10.4
* Goa 2.0.0
* Firebase

### GoによるAPIサーバーの構築
##### Goaによるデザインファーストな開発
DSLでAPIのデザインを記述して、コードを自動生成する。

* swaggerドキュメントを自動で生成してくれる
* swagger-ui を利用することで、ドキュメントの閲覧、APIの実行が可能

公式ドキュメント
https://goa.design

##### Swagger-ui によるAPI管理画面
APIの管理画面の例
 <img src=https://github.com/natsu-summer72/BbMatching/blob/master/example/swagger_console.png>
 
 APIの実行も可能
 <img src=https://github.com/natsu-summer72/BbMatching/blob/master/example/API_console.png>
 
 APIの実行例
 <img src=https://github.com/natsu-summer72/BbMatching/blob/master/example/API_result.png>

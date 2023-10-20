## Directory Structure
アプリケーションの中心となるロジックはinternalに格納され以下の様な構造になっている。
<pre>
── internal
│   ├── adapters //データストレージや、サードパーティのサービスにアクセスするような内から外に出る処理を管理する。
│   │   └── psql
│   │       ├── factory.go
│   │       ├── member.go
│   │       └── organization.go
│   ├── domains // 依存が極端にすくないビジネスロジックなどを管理する。
│   │   ├── error.go
│   │   ├── member.go
│   │   ├── member_test.go
│   │   └── organization.go
│   ├── ports // ユーザなどからリクエストなど、外から内にくる処理を管理する。
│   │   └── http
│   │       └── factory.go
│   └── usecases // ビジネスロジックを組み合わせて達成されるユースケースを管理する。
│       ├── factory.go
│       └── i_repository_keeper.go
</pre>

adapters層には依存することはなく（直接ファイルの中でimportすることはなく)インターフェースを通じてのみ必要な操作ができるようになっている。
そのため、初期化のタイミングでadapter層で作成された構造体がdiされるという仕組みになっている。

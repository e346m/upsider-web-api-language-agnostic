## 開発環境の構築について

### Demoについて

この環境はdemo用になるべく再現性が高くなるようにnix前提に設定されています。
動作確認を行うためにnixのdevcontainerを用意しているので、codespacesの方から起動して利用するようにして下さい。

コンテナ作成後にnixの環境構築フローが自動で開始されるので、見守って下さい。

構築後のフロー

1. secretのdecrypt

`git-crypt unlock symmetric_key_file`

通常では鍵となるファイルはコミットせずに、別の経路で共有されるものですが今回はdemo用に特別コミットしています。

unlockすることでconfig/secret.yamlが読み取れるようになります。

2. seedデータの作成

`make seed`

3. 開発環境の立ち上げ

`make dev`

goのREST APIサーバが立ち上がります。

### Applicationの開発について

task runnnerとしてmakeを利用しています。 `make`と打つとhelpが参照できるので必要な操作についてはmakeを通して実行することができます。
<pre>
Usage: make [target]
Targets:
api-schema-type-gen            generate response and request struct from openapi.yaml
db-down                        stop postgres
db                             start postgres
dev                            start api server
download                       start postgres
migrate-create                 make migrate-create fileName=
migrate-down                   rollback migration schema
migrate-up                     run migration schema
seed                           seed data
sqlboiler-gen                  generate code from schema
test                           run all test
</pre>

#### データベースの変更について

基本的な流れは以下のようになります。
1. `make migrate-create fileName=xxx`
2. migration fileにDDLの実装
3. `make migrate-up`
4. `make sqlboiler-gen`

特に最後の工程ではデータベースのスキーマからqueryのボイラープレートを作成しており、それを利用してアプリケーションからアクセスする仕組みになります。

#### エンドポイントの変更について

基本的な流れは以下のようになります
1. openapi.yamlの編集
2. `make api-schema-type-gen`

openapiからクライアントのコードやサーバサイドのコードを生成することを前提にopenapi.yamlをベースに編集します。
サーバサイドでは、httpサーバに柔軟性をもたせたいのでリクエストとレスポンスのstructの生成のみを最後の工程で行っています。

### 環境の更新について

基本的には[Nix](https://nixos.org/)の日本語の文献は少なく、学習コストが高いのでチーム開発ではDockerを用いましょう。

nixで管理されていないバイナリはバージョンの不整合や、build時の環境変数の違いなどによって環境に差分が出てしまうため、そういったバイナリの使用は避けたいと考えられています。(curlとかで直接ダウンロードしてきたりと)
<pre>
@e346m ➜ /workspaces/upsider-web-api-language-agnostic (fix-missing-direnv-command) $ echo $PATH
/nix/store/ay0p9mbw1w3zkvwzx3c94xq7x8jrn9wq-patchelf-0.15.0/bin:/nix/store/18bs92p6yf6w2wwxhbplgx02y6anq092-gcc-wrapper-12.3.0/bin:/nix/store/h5kvfrjmpw792v8jg7nrzfkffmn0iyy8-gcc-12.3.0/bin:/nix/store/f6in5kb2y5v06zinz1a6xy6cyg67q026-glibc-2.37-8-bin/bin:/nix/store/y9gr7abw
...
xfjqspcc9442hi0lm0szv3sw75zswvml-file-5.45/bin:/workspaces/upsider-web-api-language-agnostic/.direnv/bin:/vscode/bin/linux-alpine/f1b07bd25dfad64b0167beb15359ae573aecd2cc/bin/remote-cli:/home/vscode/.nix-profile/bin:/nix/var/nix/profiles/default/bin:/nix/var/nix/profiles/default/bin:/nix/var/nix/profiles/default/sbin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/home/vscode/.local/bin
</pre>
上記のような制約により再現性を高める仕組みが提供されていますが、反面、nixのパッケージレポジトリで管理されていないツールについては自前でビルドする必要があり、ビルド自体の難易度も高いのでチームの合意なく取り入れるのはやめましょう。　（nix自体の仕組み、各言語のビルドの仕組み、そのパッケージ自体のビルドの仕組みを把握することが難易度をあげている）


一応、簡単なケースで開発環境にライブラリ等を入れるような場合、[nix channel](https://search.nixos.org/packages)でほしいライブラリがあるかを調べて、以下のbuildInputsに追加するとターミナル上で利用できるようになります。
```nix
{
  description = "Web API with go";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    sqlboiler.url = "github:DGollings/nix-sqlboiler";
  };

  outputs = { self, nixpkgs, sqlboiler }:
    ...
    {
      devShell.${system} = mkShell {
        buildInputs = [
          go_1_21
          postgresql_16
          oapi-codegen
          sqlboiler.packages.${system}.sqlboiler
          go-migrate
          air
          git-crypt
          jq
        ];


      ...
    }
}


```

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

//エントリーポイントについて

// configと環境変数について後で

// desu masu統一

# backend

## Usage

### evans

```bash
evans --proto rpc/user/v1/user.proto repl -p 8080
```


### 参考

Clean Architecture で実装するときに知っておきたかったこと(<https://christina04.hatenablog.com/entry/go-clean-architecture>)


ざっくりDDD・クリーンアーキテクチャにおける各層の責務を理解したい①（ドメイン層・ユースケース層編）(<https://qiita.com/kotobuki5991/items/22712c7d761c659a784f>)

> Entity（およびValue Object）
ソフトウェアによって解決したい対象領域(課題)をモデリングし、コードに落とし込んだもの。
Entityは一意な識別子を持ち、変更される場合があります。
一方でValue Objectは不変であり、識別子を持ちません。
不変であるので、外部からフィールドを変更できないよう定義し、コンストラクタで初期化します。

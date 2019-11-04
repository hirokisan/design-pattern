# design-pattern

コードレベル、アーキテクチャレベルで使われるベストプラクティスであるデザインパターンを学習するリポジトリです

デザインパターンの学習を通して、良い設計・悪い設計を理解できれば良いなという気持ちです

言語仕様によってはデザインパターンが役立たない場合もあるかもしれませんが、その背景を理解すれば設計の際に役立つだろうと思っています

## コードレベル
- GoF23
  - 生成
    - [Abstract Factory](/code/creational/abstract_factory/)
    - [Builder](/code/creational/builder/)
    - [Factory Method](/code/creational/factory-method/)
    - [Prototype](/code/creational/prototype/)
    - [Singleton](/code/creational/singleton/)
  - 構造
    - [Adapter](/code/structural/adapter/)
    - [Bridge](/code/structural/bridge/)
    - [Composite](/code/structural/composite/)
    - [Decorator](/code/structural/decorator/)
    - [Facade](/code/structural/facade/)
    - [Flyweight](/code/structural/flyweight/)
    - [Proxy](/code/structural/proxy/)
  - 振る舞い
    - [Chain of Responsibility](/code/behavioral/chain_of_responsibility/)
    - [Command](/code/behavioral/command/)
    - [Interpreter](/code/behavioral/interpreter/)
    - [Iterator](/code/behavioral/iterator/)
    - [Mediator](/code/behavioral/mediator/)
    - [Memento](/code/behavioral/memento/)
    - [Observer(Pub/Sub)](/code/behavioral/observer/)
    - [State](/code/behavioral/state/)
    - [Strategy](/code/behavioral/strategy/)
    - [Template Method](/code/behavioral/template_method/)
    - [Visitor](/code/behavioral/visitor/)
- マルチスレッド
  - Active Object
  - Balking
  - Double-checked locking
  - Future
  - Guarded suspension
  - Lock
  - Monitor
  - Producer-consumer
  - Reactor
  - Readers-writer lock
  - Scheduler
  - Thread pool
  - Thread-specific storage
  - Two-phase termination
- その他
  - Abstract Class
  - Before/After
  - Generation Gap
  - Hook Operation
  - Immutable
  - Marker Interface
  - Monostate
  - Null Object
  - Sharable
  - Single-Active-Instance Singleton

## アーキテクチャレベル
- レイヤードアーキテクチャ
- ヘキサゴナルアーキテクチャ
- オニオンアーキテクチャ
- クリーンアーキテクチャ

## 関連
- アナリシスパターン
- [アンチパターン](https://github.com/hirokisan/anti-pattern)
- ソフトウェア法則・原則
  - SOLID
    - 単一責任の原則(SRP:single responsibility principle)
    - 解放閉鎖の原則(OCP:open closed principle)
    - リスコフの置換原則(LSP:liskov substitution principle)
    - インターフェイス分離の原則(ISP:interface segregation principle)
    - 依存性逆転の原則(DIP:dependency inversion principle)
    - 最小知識の原則(Principle of Least Knowledge)
  - KISS(keep it simple stupid)
  - DRY(dont repeat yourself)
  - YAGNI(You aren't gonna need it)
  - Worse is better
  - ヴィルトの法則
  - ブルックスの法則
  - コンウェイの法則
  - ホフスタッターの法則
  - 驚き最小の原則
  - ボーイスカウトの規則
  - パーキンソンの法則
  - ハインリッヒの法則
  - リーナスの法則
  - ヒックの法則
  - マーフィーの法則
  - リーマンの法則
  - ポステルの法則
  - パレートの法則
  - スタージョンの法則
  - ピーターの法則
  - ザヴィンスキーの法則
  - フィッツの法則

## 参考
- [デザインパターン(wikipedia)](https://ja.wikipedia.org/wiki/デザインパターン_%28ソフトウェア%29)
- [デザインパターン一覧(wikipedia)](https://ja.wikipedia.org/wiki/デザインパターンの一覧)
- [Java言語で学ぶデザインパターン入門](https://www.amazon.co.jp/dp/4797327030)
- [GoF以外のパターン](https://www.hyuki.com/dp/dpinfo.html)
- いろんなリポジトリがあるので、参考になりそう :sparkles:
  - [tmrts/go-patterns](https://github.com/tmrts/go-patterns)
  - [bvwells/go-patterns](https://github.com/bvwells/go-patterns)
  - [monochromegane/go_design_pattern](https://github.com/monochromegane/go_design_pattern)
- [ソフトウェア原則](http://objectclub.jp/technicaldoc/object-orientation/principle/)
- [開発者が知っておくべきSOLIDの原則](https://postd.cc/solid-principles-every-developer-should-know/)
- [何かのときにすっと出したい、プログラミングに関する法則・原則一覧](https://qiita.com/hirokidaichi/items/d6c473d8011bd9330e63)
- [人名を冠したソフトウェア開発の19の法則](https://www.yamdas.org/column/technique/19laws.html)

# design-pattern

コードレベル、アーキテクチャレベル等で使われるベストプラクティスであるデザインパターンを学習するためのリポジトリです

デザインパターンを学習することを通して、良い設計・悪い設計を理解できれば良いなという気持ちです

言語仕様によってはデザインパターンが役立つ場合も役立たない場合もあるかもしれませんが、設計の姿勢を理解すれば設計の際に役立つだろうと思っています

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
    - Observer(Pub/Sub)
    - State
    - Strategy
    - [Template Method](/code/behavioral/template_method/)
    - Visitor
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
- アンチパターン

## 参考
- [デザインパターン(wikipedia)](https://ja.wikipedia.org/wiki/デザインパターン_%28ソフトウェア%29)
- [デザインパターン一覧(wikipedia)](https://ja.wikipedia.org/wiki/デザインパターンの一覧)
- [Java言語で学ぶデザインパターン入門](https://www.amazon.co.jp/dp/4797327030)
- [GoF以外のパターン](https://www.hyuki.com/dp/dpinfo.html)
- [tmrts/go-patterns](https://github.com/tmrts/go-patterns)

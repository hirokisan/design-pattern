# chain of responsibility

責任の鎖

## [前提] どのような課題があるのか
特定の処理をするには権限が必要であり、クライアントは呼び出したインスタンスが権限を持っているのか知らないが、とりあえず処理してもらいたい場面がある

## [アプローチ] どのように課題に向き合うのか
呼び出されたインスタンスは自分よりも権限のあるインスタンスを呼び出して処理をお願いする

## 参考
- [Chain of Responsibility パターン](https://www.techscore.com/tech/DesignPattern/ChainOfResponsibility.html/)

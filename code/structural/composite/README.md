# composite

再帰的な構造を表現する

## [前提] どのような課題があるのか
容器の中に中身があり、その中身が容器となりその中に中身がある
クライアント側はそれが容器なのか中身なのか気にせずに使用したい

## [アプローチ] どのように課題に向き合うのか
再帰的な構造をinterfaceとして切り出して、各階層でinterfaceを実装する
容器に中身を設定して容器と中身を同じように扱えるようにする

## 参考
- [Composite パターン](https://www.techscore.com/tech/DesignPattern/Composite.html/)


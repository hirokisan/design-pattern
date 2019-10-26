# adapter

既存のクラス(構造体)を別のinterfaceを満たす実装に変更する

## [前提] どのような課題があるのか
新たに実装したいinterfaceがある
既存のクラスがそのinterfaceの実装に近い場合がある
新たにクラスを用意するよりは既存のクラスを活用した方がコスパが良い

## [アプローチ] どのように課題に向き合うのか
既存のクラスにinterfaceを満たすmethodを実装する
あるいは、既存のクラスを呼び出すクラスを用意する

## 参考
- [Adapterパターン 1](https://www.techscore.com/tech/DesignPattern/Adapter/Adapter1.html/)
- [Adapterパターン 2](https://www.techscore.com/tech/DesignPattern/Adapter/Adapter2.html/)

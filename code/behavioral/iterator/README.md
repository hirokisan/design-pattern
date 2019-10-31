# iterator

集約オブジェクトによる走査処理

## [前提] どのような課題があるのか
集約オブジェクトを呼び出す側で走査するとロジックが分散しがち
走査にも柔軟性をもたせたい

## [アプローチ] どのように課題に向き合うのか
集約オブジェクトからiteratorを生成する
iterator側で走査処理をする

## 参考
- [Iteratorパターン 1](https://www.techscore.com/tech/DesignPattern/Iterator/Iterator1.html/)

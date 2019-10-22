# builder

同じ作成過程で異なる表現形式の結果を得るための機構

## [前提] どのような課題があるのか
特定のプロセスで特定の材料を用いて結果を得る場合に、プロセスと材料の組み合わせによって結果は変わりうる
プロセスと材料の組み合わせに柔軟に対応できるようにしたい

## [アプローチ] どのように課題に向き合うのか
プロセス(作成過程)と材料（表現形式）をそれぞれ定義し、組み合わせることで結果が得られるようにする
directorが作成過程を管理し、builderが表現形式を管理する

## 参考
- [Builder パターン](https://www.techscore.com/tech/DesignPattern/Builder.html/)
- [Desing Patterns in Golang: Builder](http://blog.ralch.com/tutorial/design-patterns/golang-builder/)

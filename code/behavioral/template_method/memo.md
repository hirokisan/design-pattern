# memo

golang では継承という仕組みが用意されていないのでサブクラスで実装するということにはならない

しかし、interfaceを実装するという行為そのものが template method が示す概念と一致していると感じている

impl_1とimpl_2を比べると以下の違いがある

- impl_1 : interfaceを用意してinterfaceを満たす実装をする
  - あくまでinterfaceを満たす実装なので、methodというわけではなさそうだけど似たようなことはできそう
  - interfaceなので、他のmethodを持っていても許されてしまうので、templateにはなっていない
- impl_2 : structを用意して関数を埋め込む
  - 構造体の関数を外部から注入するなりして実装は切り替えられるので、こちらのほうがtemplate methodパターンに近そう
  - あらかじめ定義された構造体であり、可変なのは定義された関数の実装部分だけなのでこちらがtemplate methodと言えそう

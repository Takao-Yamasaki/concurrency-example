# concurrency-example

## 並行処理と並列処理の違い
### 1. 時間軸の観点
- 並行処理: ある時間の範囲において、複数のタスクを扱うこと
- 並列処理: ある時間の点において、複数のタスクを扱うこと

### 2. 「プログラム構成」と「プログラム実行」の観点
- 並行処理: 複数処理を独立して実行できる`構成`のこと
  - 工場で言うと、`設備`の話で、1度に多くのことを`扱う`こと
- 並列処理: 複数処理を同時に`実行`すること
  - 工場で言うと、`生産`の話で、一度に多くのことを`行う`こと

### 3. ソフトウェア・ハードウェアという観点
- 並行性(Concurrency): ソフトウェアの話(プログラムコード)
- 並列性(Parallelism): ハードウェアの話(フログラムプロセス)

## Goの「並行」処理
- Goでは、「並行処理」のための機構を、ゴルーチンやチャネルを使って提供している
- Goで作れるのはソフトウェアなので、「並行性」の話

## 並行処理をするメリット
1. 実行時間が早くなるかもしれないから
  - 並行な構成で書かれたコードは、複数のCPUに渡されて、並列に実行される可能性がある
2. 現実世界での事象が独立性・並列性を持つから

## 並行処理の難しさ
1. コードの実行順が予測できない
2. Race Condition(競合状態)を避ける必要がある
3. 共有メモリに正しくアクセスしないといけない
  - Race Conditionを避けるために、メモリに参照禁止のロックをかける

## gorutineとは？
- 他のコードに対して`並行`にしている関数のこと
- goroutineで`並行`に実装しても、`並列`に実行されるとは限らない
## gorutineの特徴
- 順不同で実行される（コードの実行順が予測できない）
```
0: alpha
5: zeta
4: pi
2: delta
6: eta
7: theta
8: epsilon
3: gamma
1: beta
This is the second thing to be printed!!
```

## 参考
- Goでの並行処理を徹底解剖！(Zenn)
https://zenn.dev/hsaki/books/golang-concurrency/viewer/intro
# umamusume-loop-factor
因子ループおぼえられましぇん

```shell
$ umamusume-loop-factor % go run main.go トウカイテイオー ニシノフラワー マックイーン ダイワスカーレット
F1: トウカイテイオー 親: ニシノフラワー + マックイーン
----- 以降ここから繰り返し -----
F1: ダイワスカーレット 親: F1 トウカイテイオー + ニシノフラワー
F1: マックイーン 親: F1 トウカイテイオー + F1 ダイワスカーレット
F1: ニシノフラワー 親: F1 マックイーン + F1 ダイワスカーレット
F2: トウカイテイオー 親: F1 マックイーン + F1 ダイワスカーレット
----- ここまで繰り返し -----
F2: ダイワスカーレット 親: F2 トウカイテイオー + F1 ニシノフラワー
```

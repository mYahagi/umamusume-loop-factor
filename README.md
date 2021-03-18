# umamusume-loop-factor
因子ループおぼえられましぇん

```shell
$ umamusume-loop-factor % go run main.go トウカイテイオー ニシノフラワー マックイーン ダイワスカーレット
トウカイテイオー
   ├─  ニシノフラワー
   └─  マックイーン
----- 以降ここから繰り返し -----
ダイワスカーレット
   ├─  トウカイテイオー
   │         ├─ニシノフラワー
   │         └─マックイーン
   └─  マックイーン
マックイーン
   ├─  トウカイテイオー
   │         ├─ニシノフラワー
   │         └─マックイーン
   └─  ダイワスカーレット
             ├─トウカイテイオー
             └─マックイーン
ニシノフラワー
   ├─  マックイーン
   │         ├─トウカイテイオー
   │         └─ダイワスカーレット
   └─  ダイワスカーレット
             ├─トウカイテイオー
             └─マックイーン
トウカイテイオー
   ├─  ニシノフラワー
   │         ├─マックイーン
   │         └─ダイワスカーレット
   └─  マックイーン
             ├─トウカイテイオー
             └─ダイワスカーレット
----- ここまで繰り返し -----
```

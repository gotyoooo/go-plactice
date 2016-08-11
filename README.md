# go-plactice

```
$ go test -run NONE -bench . -benchmem
```

* ベンチマーク関数は各自、約1秒間ずつ実行される
* 以下のようにやればテスト時間をずらせる

```
$ go test -run NONE -bench .  -benchmem -benchtime 10s
```

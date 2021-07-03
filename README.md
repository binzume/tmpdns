# TmpDNS

設定ファイルなどは存在せず，起動時のコマンドライン引数のみで完結する簡易DNSです．

ACMEv2のDNS認証用に作りました．
APIがないネームサーバを使っている場合，_acme-challengeのNSレコードに自身が管理するホストを登録し，そこでDNSサーバを起動することで認証できます．

# Usage

```shell
go install github.com/binzume/tmpdns@latest
tmpdns -p 53 "hoge.example.com.:txt:hello! world" "fuga.example.com.:a:192.168.0.1"
tmpdns -p 53 -z example.com. "hoge:txt:hello" "fuga:a:192.168.0.1"
```

(FQDNを示す `.` の書き忘れに注意)

## Docker image

https://hub.docker.com/r/binzume/tmpdns

```shell
dokcer pull binzume/tmpdns
docker run --rm -p 53:53/udp --name tmpdns binzume/tmpdns -z example.com. "hoge:txt:hello" "fuga:a:192.168.0.1"
```

query sample:

```shell
dig fuga.example.com @localhost  # 192.168.0.1
dig txt hoge.example.com @localhost  # "hello"
```

## Flags

- -p: port (default:53)
- -z: zone (default:.)

## acme.sh

ACMEv2のDNS認証用のDNSサーバとして利用できます (tmpdnsを実装した当初の目的です)

参考： https://qiita.com/binzume/items/698d12779b8ad5cda423

事前にお使いのDNSサービスに例えば `_acme-challenge.example.com` のNSレコードを登録してください．

NSレコードが正しく設定されているか確認するために，

```bash
docker run --rm -p 53:53/udp binzume/tmpdns "_acme-challenge.exmaple.com.:txt:hello"
```

としてtmpdnsを起動した状態で，

```bash
dig _acme-challenge.exmaple.com txt
```

でhelloが返ってくれば正しく設定できています．(確認が終わったらtmpdnsは終了してください)

あとは[acme.sh](https://github.com/acmesh-official/acme.sh)をインストールし，スクリプトをdnsapi下にコピーして使ってください．

```bash
cp dns_tmpdns.sh ~/.acme.sh/dnsapi
acme.sh --issue --dnssleep 10 --dns dns_tmpdns -d example.jp -d *.example.jp
```

# License

MIT License

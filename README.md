# miniDSP構築インターンシップ事前準備関連

## Mockサーバの準備
### バイナリの利用
[Release Page](https://github.com/fan-ADN/ready-for-miniDSP-Internship/releases/tag/v1.0.1)から必要なバイナリをダウンロードして、解凍してご利用ください。

- [Windows x86_64](https://github.com/fan-ADN/ready-for-miniDSP-Internship/releases/download/v1.0.1/checker_windows_x86-64.zip)
- [Mac OSX](https://github.com/fan-ADN/ready-for-miniDSP-Internship/releases/download/v1.0.1/checker_mac.zip)
- [Linux](https://github.com/fan-ADN/ready-for-miniDSP-Internship/releases/download/v1.0.1/checker_linux_x86-64.zip)

### ソースからコンパイル

Go言語(v 1.14)がコンパイルに必要となります。
``` shell
git clone https://github.com/fan-ADN/ready-for-miniDSP-Internship.git
cd tools/checker
go build
```

## Mockサーバの確認

ダウンロードした、ファイルをターミナル（コマンドプロンプト）　から実行します。バイナリのファイル名はご自身のダウンロードされたファイル名などに適宜置き換えてください。

``` shell
./checker version
```

`intern-dsp version: 1, revision 1 `このように表示されればMockサーバの準備は完了しています。

## 事前ダウンロードファイル

[adds.json](https://raw.githubusercontent.com/fan-ADN/ready-for-miniDSP-Internship/v1.0.1/adds.json)をダウンロードし、開発予定のディレクトリ配下に配置してください。

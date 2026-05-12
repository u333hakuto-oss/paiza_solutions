## [paiza レベルアップ問題集](https://paiza.jp/works/mondai)によるGo言語学習の記録
### [query_primer__range_of_score__go.go](query_primer__range_of_score__go.go)  
[クエリメニュー 点の幅](https://paiza.jp/works/mondai/query_primer/query_primer__range_of_score/edit?language_uid=go)  
Logic: 単純なループではO(KN)、最大で10^9となるため、前処理O(N)・クエリo(k√n)の平方分割を採用。  
Pitfalls: バケットサイズを100に固定しており処理効率が悪かったため、√nに変更。
逐一の出力でクエリ数の増加に伴い処理に時間がかかるため、バッファ出力に変更。

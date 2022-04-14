# ethereum-block-indexer-service

<pre>
api-service 和 ethereum-block-indexer-service
同時寫在這個專案裡

解說：
api-service
簡單的使用 gin 和 go-ethereum 完成簡單的查詢功能, 使用 controller 透過 service 做查詢功能, 
assembler 完成組裝回傳格式的功能, 這樣設計的理由是希望 service 裡面的功能可以相依性低一些, 如果之後想擴充功能, 
此 api 因為有 assembler 的原因, api 的格式就不會因為擴充 service 裡的功能而影響到

ethereum-block-indexer-service
我根據題目的意思先把 目前 Header 到創始 block 的資料先爬下來, 接著繼續爬新的資料

我的想法是把這個 chain 當成一段長度, 然後根據 env/config.yml 裡 server.IndexerNumber 這個變數, 可以把這段長度切成 n 段

|------------|------------|-------- ~~~ -------|----------------->

0          <---1        <---2            <---Header0(n)          Header1

從每個切開的斷點往前爬, 當前面資料都爬完以後
接著以 Header0 當作新的起點, 並重新抓新的 Header1 
            
~~ ---------|------|------|----  ~~~  ------|------------>
         Header0                          Header1

把 Header0 和 新 Header1 之間當作一段新的長度, 接著再切 n 段, 之後繼續上述的方式持續下去
決定把 server.IndexerNumber 切出來是希望使用者可以根據自身電腦決定數字應該怎麼調整, 數字越大爬下來的數度會越快, 
但也會越吃系統資源, 使用者重新調整後, 可以重啟服務,
此服務在 init 的時候就會先把爬過的資料 cached 起來, 所以不用擔心重覆爬的問題

Build:
make docker 先把 docker 利用 docker-compose 架起來(主要是需要db)
make init   把需要的套件和 db migrate 上去

go run main.go 
就會同時 listen on 8080 (可以在env/config.yml 修改) 用來打 api query
同時在 background run crawler


Others:
使用 https://github.com/uber-go/zap 這個套件 在 log/api.log 裡面寫紀錄 log 
使用原因是在測試期間希望可以知道到底出過什麼問題, 在運行放置測試的時間內, 可以透過這樣的方式紀錄, 方便事後查看

</pre>
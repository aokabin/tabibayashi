curl -F "user_id=hogehoge" -F "beacon_id=zzzzzzz" -F "send_date=1505704785" -F "steps=200" http://localhost:1323/v1/visit
curl -F "user_id=hogehoge" -F visits='[{"beacon_id":"yyyyyyy","send_date":"1505704785","steps":"100"}]' http://localhost:1323/v1/music
curl -F "user_id=hogehoge" http://localhost:1323/v1/music_url
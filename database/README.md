docker_databaseの起動の仕方と繋ぎ方
後々.mdの書き方にする


docker build -t todo-db . 
docker run --name mysql -p 3307:3306 todo-db

mysql -u test_u -ptest_pw -h 127.0.0.1 -P 3307
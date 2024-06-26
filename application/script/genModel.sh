#!/usr/bin/env bash

# 使用方法：
# ./model.sh usercenter user
# ./model.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./model

# 数据库配置
host=172.22.121.50
port=31820
dbname=beyond_$1
username=root
passwd=123456


echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero
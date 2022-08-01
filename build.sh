#! /bin/bash
# 获取当前路径
CUR=$(pwd)

# 判断配置文件是否存在
# 注意判断语句前后要有空格
if [[ ! -f "${CUR}/config/app.ini" ]];then
	echo 'please create app.ini fie '
	sleep 2
	exit -1
else
	echo "app.ini is existed"
fi


cd "${CUR}/core"



# 设置环境变量
export GOARCH=amd64
# linux
export GOOS=linux
# windows
# export GOOS=windows
echo 'build start'
go build -o 'core.exe' './core.go'
chmod 777 'core.exe'
echo 'build success'

cd '../'

echo 'docker start---'
docker build . -t 'dengchongsen/cloud_disk'

echo 'docker push---'
docker push dengchongsen/cloud_disk
echo 'push success'
sleep 5
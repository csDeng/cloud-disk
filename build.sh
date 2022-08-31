#! /bin/bash
# 获取当前路径
CUR=$(pwd)

# 判断配置文件是否存在
# 注意判断语句前后要有空格
if [ ! -f "${CUR}/config/app.ini" ];then
	echo 'please create app.ini fie '
	sleep 2
	exit -1
else
	echo "app.ini is existed"
fi


cd "${CUR}/core"

if [ -f "./core.exe" ];then
	echo 'start rm core.exe'
	rm "./core.exe"
	if [ $? -ne 0 ]; then
		echo 'rm core.exe failed'
		exit -1
	else 
		echo 'rm core.exe success!!!'
	fi
else 
	echo 'core.exe is not existed. pro continue....'
fi


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

read -p "please input a tag:\r\n" tag

image="dengchongsen/cloud_disk:${tag}"

echo 'docker start---'
docker build . -t $image

echo 'docker push---'
docker push $image
echo 'push success'
sleep 30
projectName="agollo-demo"
echo "构建$projectName....."
echo "=============================="
echo "设置GOPATH.."
srcBaseDir="$(pwd)"
cd ..
baseDir="$(pwd)"
export GOPATH=$baseDir

echo "编译中.."
cd $srcBaseDir
rm -rf build
go build -o "$projectName" main/check.go

echo "构建运行环境.."
mkdir build
cp -rf $projectName build/
cp -rf seelog.xml build/
cp -rf app.properties build/
# AListMobileService
```shell
#build and push mobile lib
#install gomobile(at system cli)
go install golang.org/x/mobile/cmd/gobind@latest
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
gomobile version
go get -u golang.org/x/mobile/...
#
export GO111MODULE="off"
gomobile bind -target=android -o=./build/mobile.aar ./libs/mobile
export MACOSX_DEPLOYMENT_TARGET=10.6
export IPHONEOS_DEPLOYMENT_TARGET=9.0
go clean -cache
gomobile bind -ldflags '-w -s -extldflags "-lresolve"' --target=ios,macos,iossimulator -o ./build/AListMobile.xcframework ./libs/mobile

#push ios,macos libs
#https://gitee.com/AListMobile/mobile-lib-podspec
#git tag -a 0.0.1 -m '0.0.1'
#git pus --tags
#pod trunk push ./AListMobile.podspec --skip-import-validation --allow-warnings

#
#modify version in pom
#push android lib
mvn gpg:sign-and-deploy-file -DrepositoryId=ossrh -Dfile=mobile.aar -DpomFile=mobile.pom -Durl=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
# https://s01.oss.sonatype.org/
mvn deploy:deploy-file -Dfile=mobile.aar -DgroupId=cloud.iothub -DartifactId=mobile -Dversion=0.0.1 -Dpackaging=aar -DrepositoryId=github -Durl=https://maven.pkg.github.com/AListMobile/gateway-go
```
```shell
#for build windows dll
echo "building windows dll"
#brew install mingw-w64
#sudo apt-get install binutils-mingw-w64
# shellcheck disable=SC2034
export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export GOOS=windows GOARCH=amd64
go build -tags windows -ldflags=-w -trimpath -o ./build/windows/mobile_amd64.dll -buildmode=c-shared ./libs/dynamic
```
```shell
#for build linux/android so file
echo "building linux/android so file"
#linux和Android共用动态链接库
export CGO_ENABLED=1
export GOARCH=amd64
export GOOS=linux
go build -tags linux -ldflags=-w -trimpath -o build/linux/libmobile_amd64.so -buildmode c-shared ./libs/dynamic

# shellcheck disable=SC2034
export CGO_ENABLED=1
export GOARCH=arm64
export GOOS=linux
#sudo apt install gcc-aarch64-linux-gnu
export CC=aarch64-linux-gnu-gcc
##sudo apt install g++-aarch64-linux-gnu
#export CXX=aarch64-linux-gnu-g++
##sudo apt-get install binutils-aarch64-linux-gnu
#export AR=aarch64-linux-gnu-ar
go build -tags linux -ldflags=-w -trimpath -o build/linux/libmobile_arm64.so -buildmode c-shared ./libs/dynamic
```
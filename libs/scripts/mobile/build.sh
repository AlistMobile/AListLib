./clean.sh
export GO111MODULE="off"
gomobile bind -target=android -o=./mobile.aar ./libs/mobile
export MACOSX_DEPLOYMENT_TARGET=10.14
gomobile bind -ldflags '-w -s -extldflags "-lresolve"' --target=ios,macos,iossimulator -o AListMobile.xcframework ./libs/mobile
#https://gitee.com/OpenIoThub/mobile-lib-podspec
#git tag -a 0.0.1 -m '0.0.1'
#git pus --tags
#pod trunk push ./OpenIoTHubMobile.podspec --skip-import-validation --allow-warnings

mvn gpg:sign-and-deploy-file -DrepositoryId=ossrh -Dfile=mobile.aar -DpomFile=mobile.pom -Durl=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
mvn deploy:deploy-file -Dfile=mobile.aar -DgroupId=cloud.iothub -DartifactId=mobile -Dversion=0.0.1 -Dpackaging=aar -DrepositoryId=github -Durl=https://maven.pkg.github.com/AListMobile/gateway-go
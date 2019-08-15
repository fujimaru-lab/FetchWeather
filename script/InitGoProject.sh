# !/bin/sh

# 2019/08/12 ShellName: InitGoProject.sh 
# Goのプロジェクトの標準的なプロジェクト構成を作成するシェルスクリプト
# シェルスクリプトの引数に渡したプロジェクト名でGoプロジェクトを作成する
# 引数が渡されなかった場合は、エラーメッセージを出力し終了
# 同じプロジェクト名が存在した場合は、エラーメッセージを出力し終了
# common.env内にて、git-hubのアカウント名を設定する必要がある。-> GITHUB_USERNAME
# arg1:Goプロジェクト名

. ./common.env

GO_PROJ_HOME=${GOPATH}/src/github.com/${GITHUB_USERNAME}/

if [ ! -e ${GO_PROJ_HOME} ]; then
    mkdir ${GO_PROJ_HOME}
fi

cd ${GO_PROJ_HOME}

GO_PROJ_NAME=$1

if [ -z "${GO_PROJ_NAME}" ]; then
    echo "You shoud pass argument."
    echo "process end..."
    exit 111
fi

if [ -e ./${GO_PROJ_NAME} ]; then
    echo "Same Project is already exists."
    echo "process end..."
    exit 222
else
    # プロジェクトのルートフォルダを生成
    mkdir ${GO_PROJ_NAME}
    cd ./${GO_PROJ_NAME}

    # git init
    git init
    
    # プロジェクトの各子フォルダを生成
    # /cmd
    mkdir ./cmd
    touch ./cmd/main.go
    echo 'package main' > ./cmd/main.go
    echo '' >> ./cmd/main.go
    echo 'func main() {' >> ./cmd/main.go
    echo '}' >> ./cmd/main.go

    # /internal
    mkdir ./internal
    
    # /pkg
    mkdir ./pkg

    # /vendor
    mkdir ./vendor

    # /api
    mkdir ./api

    # /web
    mkdir ./web
    mkdir ./web/template
    mkdir ./web/image
    mkdir ./web/css

    # /configs
    mkdir ./configs

    # /init
    mkdir ./init

    # /script
    mkdir ./script

    # /build
    mkdir ./build

    # /deployments
    mkdir ./deployments

    # /doc
    mkdir ./doc
    touch ./doc/readme.md
    echo "# ${GO_PROJ_NAME}" >> ./doc/readme.md
    echo '' >> ./doc/readme.md
    echo 'Project Start Date:'`date` >> ./doc/readme.md

    # /test
    mkdir ./test

    # /resource
    mkdir ./resource

    git add .
    git commit -m "first commit. create project ${GO_PROJ_NAME}"

    # end message
    echo "Your ${GO_PROJ_NAME} is ready to start."
    echo "Fun Go-ing :)"
fi

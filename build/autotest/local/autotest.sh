#!/usr/bin/env bash

set -e
set -o pipefail
#set -o verbose
#set -o xtrace

# os: ubuntu16.04 x64
# first, you must install jq tool of json
# sudo apt-get install jq
# sudo apt-get install shellcheck, in order to static check shell script
# sudo apt-get install parallel

PWD=$(cd "$(dirname "$0")" && pwd)
export PATH="$PWD:$PATH"

CLI="./dplatform-cli"

sedfix=""
if [ "$(uname)" == "Darwin" ]; then
    sedfix=".bak"
fi

dplatformConfig="dplatform.test.toml"
dplatformBlockTime=1
autoTestCheckTimeout=10

function config_dplatform() {

    # shellcheck disable=SC2015
    echo "# config dplatform solo test"
    # update test environment
    sed -i $sedfix 's/^Title.*/Title="local"/g' ${dplatformConfig}
    # grep -q '^TestNet' ${dplatformConfig} && sed -i $sedfix 's/^TestNet.*/TestNet=true/' ${dplatformConfig} || sed -i '/^Title/a TestNet=true' ${dplatformConfig}

    if grep -q '^TestNet' ${dplatformConfig}; then
        sed -i $sedfix 's/^TestNet.*/TestNet=true/' ${dplatformConfig}
    else
        sed -i $sedfix '/^Title/a TestNet=true' ${dplatformConfig}
    fi

    #update fee
    #    sed -i $sedfix 's/Fee=.*/Fee=100000/' ${dplatformConfig}

    #update wallet store driver
    #    sed -i $sedfix '/^\[wallet\]/,/^\[wallet./ s/^driver.*/driver="leveldb"/' ${dplatformConfig}
}

autotestConfig="autotest.toml"
autotestTempConfig="autotest.temp.toml"
function config_autotest() {

    echo "# config autotest"
    #delete all blank lines
    #    sed -i $sedfix '/^\s*$/d' ${autotestConfig}

    if [[ $1 == "" ]] || [[ $1 == "all" ]]; then
        cp ${autotestConfig} ${autotestTempConfig}
        sed -i $sedfix 's/^checkTimeout.*/checkTimeout='${autoTestCheckTimeout}'/' ${autotestTempConfig}
    else
        #copy config before [
        # sed -n '/^\[\[/!p;//q' ${autotestConfig} >${autotestTempConfig}
        #pre config auto test
        {

            echo 'cliCmd="./dplatform-cli"'
            echo "checkTimeout=${autoTestCheckTimeout}"
        } >${autotestTempConfig}

        #specific dapp config
        for dapp in "$@"; do
            {
                echo "[[TestCaseFile]]"
                echo "dapp=\"$dapp\""
                echo "filename=\"$dapp.toml\""
            } >>${autotestTempConfig}

        done
    fi
}

function start_dplatform() {

    echo "# start solo dplatform"
    rm -rf ../local/datadir ../local/logs ../local/grpc33.log
    ./dplatform -f dplatform.test.toml >/dev/null 2>&1 &
    local SLEEP=1
    sleep ${SLEEP}

    # query node run status
    ${CLI} block last_header

    echo "=========== # save seed to wallet ============="
    result=$(${CLI} seed save -p 1314fuzamei -s "tortoise main civil member grace happy century convince father cage beach hip maid merry rib" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    echo "=========== # unlock wallet ============="
    result=$(${CLI} wallet unlock -p 1314fuzamei -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    echo "=========== # import private key returnAddr ============="
    result=$(${CLI} account import_key -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944 -l returnAddr | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== # import private key mining ============="
    result=$(${CLI} account import_key -k 4257D8692EF7FE13C68B65D6A52F03933DB2FA5CE8FAF210B5B8B80C721CED01 -l minerAddr | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== #transfer to miner addr ============="
    hash=$(${CLI} send coins transfer -a 10000 -n test -t 12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)
    echo "${hash}"
    sleep ${dplatformBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

    echo "=========== #transfer to token amdin ============="
    hash=$(${CLI} send coins transfer -a 10 -n test -t 1Q8hGLfoGe63efeWa8fJ4Pnukhkngt6poK -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)

    echo "${hash}"
    sleep ${dplatformBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

    echo "=========== #config token blacklist ============="
    rawData=$(${CLI} config config_tx -c token-blacklist -o add -v BTC)
    signData=$(${CLI} wallet sign -d "${rawData}" -k 0xc34b5d9d44ac7b754806f761d3d4d2c4fe5214f6b074c19f069c4f5c2a29c8cc)
    hash=$(${CLI} wallet send -d "${signData}")

    echo "${hash}"
    sleep ${dplatformBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

}

function start_autotest() {

    echo "=========== #run autotest, make sure saving autotest.last.log file============="

    if [ -e autotest.log ]; then
        cat autotest.log >autotest.last.log
        rm autotest.log
    fi

    ./autotest -f ${autotestTempConfig}

}

function stop_dplatform() {

    rv=$?
    echo "=========== #stop dplatform ============="
    ${CLI} close || true
    #wait close
    sleep ${dplatformBlockTime}
    echo "==========================================local-auto-test-shell-end========================================================="
    exit ${rv}
}

function main() {
    echo "==========================================local-auto-test-shell-begin========================================================"
    config_autotest "$@"
    config_dplatform
    start_dplatform
    start_autotest

}

trap "stop_dplatform" INT TERM EXIT

main "$@"


## 安装

```
$ sudo install dplatformos-cli.bash  /usr/share/bash-completion/completions/dplatformos-cli
```

不安装
```
. dplatformos-cli.bash
```

```
# 重新开个窗口就有用了
$ ./dplatformos/dplatformos-cli 
account   dpom       coins     evm       hashlock  mempool   privacy   retrieve  send      ticket    trade     version   
block     close     config    exec      help      net       relay     seed      stat      token     tx        wallet    
```

## 演示
```
linj@linj-TM1701:~$ ./dplatformos/dplatformos-cli 
account   dpom       coins     evm       hashlock  mempool   privacy   retrieve  send      ticket    trade     version   
block     close     config    exec      help      net       relay     seed      stat      token     tx        wallet    
linj@linj-TM1701:~$ ./dplatformos/dplatformos-cli b
block  dpom    
linj@linj-TM1701:~$ ./dplatformos/dplatformos-cli dpom 
priv2priv  priv2pub   pub2priv   send       transfer   txgroup    withdraw   
linj@linj-TM1701:~$ ./dplatformos/dplatformos-cli dpom t
transfer  txgroup   
linj@linj-TM1701:~$ ./dplatformos/dplatformos-cli dpom transfer -
-a        --amount  -h        --help    -n        --note    --para    --rpc     -t        --to  
```


## 原理

 地址 https://confluence.33.cn/pages/viewpage.action?pageId=5967798

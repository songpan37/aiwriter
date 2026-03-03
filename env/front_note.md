
查看配置：
```
npm config list
; "builtin" config from E:\app\nodejs\node_modules\npm\npmrc

; prefix = "C:\\Users\\Song Pan\\AppData\\Roaming\\npm" ; overridden by user

; "user" config from C:\Users\Song Pan\.npmrc

cache = "E:\\app\\nodejs\\node_cache_repo"
prefix = "E:\\app\\nodejs\\node_global_repo"

; node bin location = E:\app\nodejs\node.exe
; node version = v24.14.0
; npm local prefix = C:\Users\Song Pan
; npm version = 11.9.0
; cwd = C:\Users\Song Pan
; HOME = C:\Users\Song Pan
; Run `npm config ls -l` to show all defaults.
```


设置本地存储位置和国内镜像：
```
# 1. 设置国内淘宝镜像源（加速下载）
npm config set registry https://registry.npmmirror.com

# 2. 设置本地存储位置（以D盘为例，请替换为你想要的路径）
npm config set prefix "D:\nodejs\node_global"
npm config set cache "D:\nodejs\node_cache"
```


初始化react项目：
```
npm create vite@latest my-react-app -- --template react-ts
```

初始化后，启动vite：
```
npm run dev
```

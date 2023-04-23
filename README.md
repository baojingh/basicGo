# basicGo

# 当前功能
1. 可以进入/bin/sh并执行操作
2. 内存限制功能完成


# 问题
1. 限制内存是100m
    ```
    md run  -ti  -m  100m /bin/sh
    ```
2. 容器中执行stress测试内存是否限制成功。使用99m是正常的，但是不能超过100m。原因未知。
    ```
        sh-4.2# stress --vm-bytes 200m --vm 1 
        stress: info: [5] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd
        stress: FAIL: [5] (415) <-- worker 6 got signal 9
        stress: WARN: [5] (417) now reaping child worker processes
        stress: FAIL: [5] (451) failed run completed in 0s
        sh-4.2# stress --vm-bytes 100m --vm 1 
        stress: info: [7] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd
        stress: FAIL: [7] (415) <-- worker 8 got signal 9
        stress: WARN: [7] (417) now reaping child worker processes
        stress: FAIL: [7] (451) failed run completed in 0s
        sh-4.2# stress --vm-bytes 99m --vm 1 
        stress: info: [9] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd
    ```

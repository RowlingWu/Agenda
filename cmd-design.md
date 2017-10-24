#agenda#
agenda [命令] [参数]

##命令  参数  用途##

`help`  列出命令说明

`register -uUsername  -pPassword  -eEmail  -tTelephone`    用户注册 

`login -uUsername -pPassword` 用户登录

`logout` 用户登出 

`qu` 用户查询 

`del` 删除用户 

`cm -tTitle  -pParticipant  -sStart_time  -eEnd_time` 创建会议 

`mod -tTitle  -aParticipant_added  -dParticipant_deleted` 增删会议参与者 

`qm -sStart_time  -eEnd_time -m -a`  查询会议（`-m`表示已登录的用户查询自己的议程在某一时间段内的所有会议安排；`-a`表示在指定时间范围内找到的所有会议）

`cancel -tTitle` 取消会议 

`quit -tTitle` 退出会议 

`clear` 清空会议 

const ws=require('nodejs-websocket')
const PORT=9000
const TYPEENTER=0
const TYPELEAVE=1
const TYPEMSG =2
const TYPEREGISTER=3
const TYPESENDMSG=4
//创建一个count，表示用户个数
//每次只要有人连接，就会有一个connect被创建
const sever=ws.createServer(connect=>{
    //console.log('有用户连接成功')
    connect.on('text',function(data){//每当有用户连接时，这个事件就会被触发
        console.log('接受到了用户的数据',data)
        //广播给所有用户某人发的消息(包括自己)
       var data01=JSON.parse(data)
        if(data01.type===TYPESENDMSG){
            var msgdata={
                type:TYPEMSG,
                msg:connect.username+'(用户名)'+':'+data01.connext,
                time:new Date().toLocaleTimeString()
            }
            console.log('返回数据',msgdata)
            broadcast(msgdata)
        }else if(data01.type===TYPEREGISTER){
            connect.username=data01.name
            var num=loginprocess(data,connect)
            console.log(num)
            connect.send(JSON.stringify({
                type:TYPEREGISTER,
                msg:num,
                name:connect.username
            }))
            if(num===200){
                if(connect.username!=''){
                    broadcast({
                        type:TYPEENTER,
                        msg:`${connect.username}(用户名)进入了聊天室`,
                        time:new Date().toLocaleTimeString()
                    })
                }
            }
           
        }
    })
    connect.on('error',function(){//错误处理----防止出现错误服务器崩掉
        console.log('有异常错误')
    })
    connect.on('close',function(){//关闭处理
        broadcast({
            type:TYPELEAVE,
            msg:`${connect.username}(用户名)离开了聊天室`,
            time:new Date().toLocaleTimeString()//获取时间
        })
        console.log('有人断开连接')
    })
})
//监听端口
sever.listen(PORT)
//广播函数
function broadcast(msg){
    //拿到所有用户的连接
    sever.connections.forEach(item=>{
        item.send(JSON.stringify(msg))
    })
}
function loginprocess(msg, connect) {
    msg = JSON.parse(msg)
    var conflict = sever.connections.some(item => {
        return msg.name === item.username && item != connect
    })
    if (conflict) {
        return 404
    } else {
        return 200
    }
}

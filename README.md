作业问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

```go
需要Wrap这个error进行包装给上层 原因在于：
1.防止sentinel error 哨兵error判定的==出现不用error的问题
2.为了对该error的错误提供上下文以及相关详细信息
3.详细的找出错误error并打印调用堆栈信息 方便排查并快速的定位问题
```

代码思路:

```go
1.initsql 初始化sql  
  sql.open 和 db.ping  初始对于mysql操作 对于出现问题errors.Wrap包装详细信息

2.建立相应数据库表User 字段Id Name

3.queryUser 方法查询用户
对于用户信息查询
for rows.Next(){
		user := User{}
		if err = rows.Scan(&user.Id, &user.Name); err != nil{
			switch {
				case err == sql.ErrNoRows:
					err = errors.Wrap(err, "ErrNoRows")
				default:
					err = errors.Wrap(err, "Unknown Errors")
			}
			return err
		}
 对于相应的sql查询错误进行特定包装  该处switch想通过断言判定但失败了 但根据pkg/errors包中的一些案例是通过指针判定 希望助教帮忙修改一下
  
    
    
4. main 函数
   测试样例
   详细的堆栈信息"%+v"
```


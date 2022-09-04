//
//  Types.swift
//  hollow
//
//  Created by ozline on 2022/8/28.
//

/* 这部分保存一些数据的结构体 */

import Foundation

struct User {
    var id : Int64          // 用户id
    var status : Int64      // 用户状态
    var username : String   // 用户名
    var password : String   // 密码
    var email : String      // 邮箱
    var phone : String      // 手机号
    var created_at : Int64  // 创建时间
    var deleted_at : Int64  // 删除时间
    var updated_at : Int64  // 更新时间
    var nickname : String   // 昵称
    var mfa_enabled : Bool  // 是否启用MFA认证
    var mfa_secret : String // MFA秘钥
}

struct Leaf {
    var id : Int64          // 帖子ID
    var owner : Int64       // 发表者id (若匿名此处为0)
    var created_at : Int64  // 创建时间
    var status : Int64      // 消息状态 0=匿名 1=实名 2=封禁
    var message : String    // 消息
    var liked : Int64       // 点赞数
}

struct Comment {
    var id : Int64          // 评论id
    var ownder : Int64      // 用户
    var root : Int64        // 归属帖子id
    var father : Int64      // 父楼id 0=直接归属root 其他数字则为上一层id
    var created_at : Int64  // 创建时间
    var status : Int64      // 状态 0=匿名 1=实名 2=封禁（删除）
    var message : String    // 评论内容
    var liked : Int64       // 点赞数
}

struct Like {
    var timestamp : Int64   // 评论时间（当做点赞id)
    var user : Int64        // 用户
    var comment : Int64     // 评论id
}

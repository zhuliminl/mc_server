package database

const createUserTable = `
  CREATE TABLE IF NOT EXISTS user (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    userId VARCHAR(50) COMMENT '用户 userId',
    username VARCHAR(32) COMMENT '随机默认用户名 | 用户自己设定名',
    email VARCHAR(50) COMMENT '邮箱',
    phone VARCHAR(20) COMMENT '手机号码',
    password VARCHAR(100) COMMENT '密码',
    wechatNickname VARCHAR(32) COMMENT '微信昵称',
    wechatNumber VARCHAR(20) COMMENT '微信号'
  );
`

const CreateUser = `INSERT INTO user(userId, username, email, phone, password ) VALUES ( ?, ?, ?, ?, ? )`
const DeleteUserByUserId = `DELETE FROM user WHERE userId=?`
const FindUserByUserId = `SELECT userId, username, email, phone, wechatNickname, wechatNumber from user WHERE userId=?`
const FindUserByEmail = `SELECT userId, username, email, phone, password, wechatNickname, wechatNumber from user WHERE email=?`
const FindUserAll = `SELECT userId, username, email, phone, wechatNickname from user`
const UpdateUser = ``

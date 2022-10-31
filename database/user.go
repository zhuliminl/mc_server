package database

const createUserTable = `
  CREATE TABLE IF NOT EXISTS user (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(32) COMMENT '随机默认用户名 | 用户自己设定名',
    email VARCHAR(50) COMMENT '邮箱',
    phone VARCHAR(20) COMMENT '手机号码',
    password VARCHAR(100) COMMENT '密码',
    wechatNickname VARCHAR(32) COMMENT '微信昵称',
    wechatNumber VARCHAR(20) COMMENT '微信号'
  );
`

// 创建用户
const createUser = ``

// 获取用户
const findUser = ``

// 更新用户
const updateUser = ``

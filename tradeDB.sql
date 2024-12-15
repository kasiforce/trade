create database trade_db;
use trade_db;

-- 用户表
create table users(
	userID int primary key auto_increment,
    userName varchar(30) not null unique,
    passwords varchar(30) not null,
    schoolID int not null,
    picture varchar(256),
    tel varchar(20),
    mail varchar(40) not null unique,
    gender tinyint,
    userStatus tinyint not null default 0
);

-- 商品类别表
 CREATE TABLE category (
    categoryID INT NOT NULL AUTO_INCREMENT,
    categoryName VARCHAR(20) NOT NULL,
    descriptions TEXT,
    PRIMARY KEY (categoryID)
);

-- 商品表 
CREATE TABLE goods (
    goodsID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    goodsName VARCHAR(30) NOT NULL,
    userID INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    categoryID INT NOT NULL,
    details TEXT,
    isSold TINYINT NOT NULL DEFAULT 0,
    goodsImages VARCHAR(256),
    createdTime DATETIME NOT NULL,
    deliveryMethod INT NOT NULL,  -- 新增字段 deliveryMethod
    shippingCost DECIMAL(10, 2) NULL DEFAULT 0,    -- 新增字段 shippingCost
    addrID INT,                            -- 新增字段 addrID 作为外键
    view INT NOT NULL DEFAULT 0,  -- 新增字段 view
    FOREIGN KEY (userID) REFERENCES users(userID), 
    FOREIGN KEY (categoryID) REFERENCES category(categoryID),
    FOREIGN KEY (addrID) REFERENCES address(addrID)  -- 新增外键约束
) DEFAULT CHARSET=utf8mb4;

-- 收藏 
CREATE TABLE collection (
    goodsID INT NOT NULL,
    userID INT NOT NULL,
    createdTime DATETIME NOT NULL,
    PRIMARY KEY (goodsID, userID),
    FOREIGN KEY (goodsID) REFERENCES goods(goodsID),
    FOREIGN KEY (userID) REFERENCES users(userID)
);

-- 收货地址 
CREATE TABLE address (
    addrID INT NOT NULL AUTO_INCREMENT,
    userID INT NOT NULL,
    province VARCHAR(30) NOT NULL,
    city VARCHAR(30) NOT NULL,
    districts VARCHAR(30) NOT NULL,
    address VARCHAR(256) NOT NULL,
    tel VARCHAR(20) NOT NULL,
    receiver VARCHAR(100) NOT NULL,
    isDefault TINYINT NOT NULL DEFAULT 1,
    PRIMARY KEY (addrID),
    FOREIGN KEY (userID) REFERENCES users(userID)
); 

-- 交易记录 
CREATE TABLE trade_records (
    tradeID INT NOT NULL AUTO_INCREMENT,
    sellerID INT NOT NULL,
    buyerID INT NOT NULL,
    goodsID INT NOT NULL,
    turnoverAmount DECIMAL(10,2) NOT NULL CHECK (turnoverAmount >= 0),
    shippingAddrID INT,
    deliveryAddrID INT,
    orderTime DATETIME NOT NULL,
    payTime DATETIME,
    shippingTime DATETIME,
    shippingCost DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (shippingCost >= 0),
    turnoverTime DATETIME,
    payMethod TINYINT NOT NULL DEFAULT 0,
    trackingNumber VARCHAR(100),
    isReturn TINYINT NOT NULL DEFAULT 0,
    status VARCHAR(20) NOT NULL,
    PRIMARY KEY (tradeID),
    FOREIGN KEY (sellerID) REFERENCES users(userID),
    FOREIGN KEY (buyerID) REFERENCES users(userID),
    FOREIGN KEY (goodsID) REFERENCES goods(goodsID),
    FOREIGN KEY (shippingAddrID) REFERENCES address(addrID),
    FOREIGN KEY (deliveryAddrID) REFERENCES address(addrID)
);

-- 退货记录 
CREATE TABLE refund_records (
    refundID INT NOT NULL AUTO_INCREMENT,
    tradeID INT NOT NULL,
    payMethod TINYINT NOT NULL,
    refundAgreedTime DATETIME NOT NULL,
    refundShippedTime DATETIME,
    refundArrivalTime DATETIME,
    trackingNumber VARCHAR(100),
    PRIMARY KEY (refundID),
    FOREIGN KEY (tradeID) REFERENCES trade_records(tradeID)
); 

-- 管理员
CREATE TABLE admins (
    adminID INT NOT NULL AUTO_INCREMENT,
    passwords VARCHAR(30) NOT NULL,
    adminName varchar(30) not null unique,
    tel VARCHAR(20),
    mail VARCHAR(40) NOT NULL,
    gender TINYINT,
    age INT CHECK (age >= 0),
    PRIMARY KEY (adminID)
);

-- 举报 
CREATE TABLE report (
    reportID INT NOT NULL AUTO_INCREMENT,
    userID INT NOT NULL,
    goodsID INT NOT NULL,
    reason TEXT NOT NULL,
    createdTime DATETIME NOT NULL,
    PRIMARY KEY (reportID),
    FOREIGN KEY (userID) REFERENCES users(userID),
    FOREIGN KEY (goodsID) REFERENCES goods(goodsID)
)DEFAULT CHARSET=utf8mb4;

-- 聊天记录 
CREATE TABLE chat_records (
    chatID INT NOT NULL AUTO_INCREMENT,
    senderID INT NOT NULL,
    receiverID INT NOT NULL,
    message TEXT NOT NULL,
    sendTime DATETIME,
    PRIMARY KEY (chatID),
    FOREIGN KEY (senderID) REFERENCES users(userID),
    FOREIGN KEY (receiverID) REFERENCES users(userID)
)DEFAULT CHARSET=utf8mb4; 

-- 退货申诉 
CREATE TABLE refund_complaint (
    complaintID INT NOT NULL AUTO_INCREMENT,
    tradeID INT NOT NULL,
    buyerReason TEXT NOT NULL,
    cTime DATETIME NOT NULL,
    cStatus TINYINT NOT NULL DEFAULT 0,
    sellerReason Text,
    PRIMARY KEY (complaintID),
    FOREIGN KEY (tradeID) REFERENCES trade_records(tradeID)
)DEFAULT CHARSET=utf8mb4;

-- 学校 
CREATE TABLE school (
    schoolID INT NOT NULL AUTO_INCREMENT,
    schoolName VARCHAR(80) NOT NULL,
    mailSuffix VARCHAR(30) NOT NULL,
    PRIMARY KEY (schoolID)
);

-- 评论
CREATE TABLE comment (
    commentID INT NOT NULL AUTO_INCREMENT,
    goodsID INT NOT NULL,
    commentatorID INT NOT NULL,
    commentContent TEXT NOT NULL,
    commentTime DATETIME NOT NULL,
    PRIMARY KEY (commentID),
    FOREIGN KEY (goodsID) REFERENCES goods(goodsID),
    FOREIGN KEY (commentatorID) REFERENCES users(userID)
)DEFAULT CHARSET=utf8mb4;

-- 公告
CREATE TABLE announcement (
    announcementID INT NOT NULL AUTO_INCREMENT,
    anTitle TEXT NOT NULL,
    anContent TEXT NOT NULL,
    anTime DATETIME NOT NULL,
    PRIMARY KEY (announcementID)
)DEFAULT CHARSET=utf8mb4;
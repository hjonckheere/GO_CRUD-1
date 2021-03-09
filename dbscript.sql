
drop database if exists custRegFormDB;

create database if not exists custRegFormDB;

use custRegFormDB;

CREATE TABLE IF NOT EXISTS Customer (
  cusid VARCHAR(20) NOT NULL,
  fullname VARCHAR(100) NULL,
  nic VARCHAR(20) NULL,
  contact int(10) NULL,
  address VARCHAR(100) NULL,
  CONSTRAINT PRIMARY KEY (cusid)
);
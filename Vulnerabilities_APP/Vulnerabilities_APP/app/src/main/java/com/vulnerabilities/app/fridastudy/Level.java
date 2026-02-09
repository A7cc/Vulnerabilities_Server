package com.vulnerabilities.app.fridastudy;

// 定义关卡类
public class Level {
    // 当前关卡等级
    public int id;
    // 该关卡解锁情况
    public boolean unlocked;
    // 该关卡通过情况
    public boolean passed;

    public Level(int id, boolean unlocked, boolean passed) {
        this.id = id;
        this.unlocked = unlocked;
        this.passed = passed;
    }
}

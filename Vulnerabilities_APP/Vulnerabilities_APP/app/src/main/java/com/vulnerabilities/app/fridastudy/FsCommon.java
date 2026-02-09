package com.vulnerabilities.app.fridastudy;

import android.util.Log;
import java.util.Objects;
import static com.vulnerabilities.app.fridastudy.FridaStudyActivity.GenerateFlag;
import static com.vulnerabilities.app.fridastudy.FridaStudyActivity.showResultsStatic;

public class FsCommon {
    int num;
    String str;
    FsCommon(int num, String str) {
        this.num = num;
        this.str = str;
    }
    static int code = 0;
    public static void increase() {
        code += 2;
    }
    // 第四关：主动调用动态函数
    public void LevelFour(String text) {
        Log.d("FridaStudy", "LevelFour");
        if (Objects.equals(text, "a7cc")) {
            String flag = GenerateFlag("04_"+text);
            // 直接调用 Activity 的显示逻辑
            showResultsStatic(flag);
        }
    }
}

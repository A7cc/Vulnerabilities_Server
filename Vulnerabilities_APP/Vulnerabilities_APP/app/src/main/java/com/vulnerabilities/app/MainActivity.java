package com.vulnerabilities.app;

import android.content.Intent;
import android.os.Bundle;
import androidx.appcompat.app.AppCompatActivity;
import com.vulnerabilities.app.crypto.CryptoActivity;
import com.vulnerabilities.app.databinding.ActivityMainBinding;
import com.vulnerabilities.app.http.HttpActivity;
import com.vulnerabilities.app.other.OtherCheckActivity;
import com.vulnerabilities.app.rootcheck.RootCheckActivity;

import java.io.BufferedReader;
import java.io.FileReader;

public class MainActivity extends AppCompatActivity {



    private ActivityMainBinding mainBinding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 初始化 ViewBinding
        mainBinding = ActivityMainBinding.inflate(getLayoutInflater());
        // 2. 设置布局
        setContentView(mainBinding.getRoot());
        // 3. 绑定点击事件（核心功能入口）

        // HTTP请求
        mainBinding.cardHttp.setOnClickListener(v -> {
            startActivity(new Intent(this, HttpActivity.class));
        });

        // 加密算法
        mainBinding.cardCrypto.setOnClickListener(v -> {
            startActivity(new Intent(this, CryptoActivity.class));
        });

        // ROOT检测
        mainBinding.cardRoot.setOnClickListener(v -> {
            startActivity(new Intent(this, RootCheckActivity.class));
        });

        // Other检测
        mainBinding.cardOther.setOnClickListener(v -> {
            startActivity(new Intent(this, OtherCheckActivity.class));
        });
    }
    /**
     * 检测 TracerPid（是否被调试）：java层
     */
    public static boolean isDebuggerAttached() {
        try (BufferedReader br = new BufferedReader(
                new FileReader("/proc/self/status"))) {

            String line;
            while ((line = br.readLine()) != null) {
                if (line.startsWith("TracerPid:")) {
                    String tracerPid = line.split(":")[1].trim();
                    return !tracerPid.equals("0");
                }
            }
        } catch (Exception ignored) {
        }
        return false;
    }
    // native层
//    #include <sys/ptrace.h>
//            #include <errno.h>
//
//    JNIEXPORT jboolean JNICALL
//    Java_com_vulnerabilities_app_RootCheckUtil_nativeCheckRoot(
//                    JNIEnv *env,
//                    jclass clazz) {
//
//        // ptrace 反调试
//        if (ptrace(PTRACE_TRACEME, 0, NULL, 0) == -1) {
//            return JNI_TRUE; // 被调试
//        }
//
//        // uid 检测
//        if (getuid() == 0) {
//            return JNI_TRUE;
//        }
//
//        // su 文件
//        struct stat st;
//        if (stat("/system/bin/su", &st) == 0) {
//            return JNI_TRUE;
//        }
//
//        return JNI_FALSE;
//    }

}
package com.vulnerabilities.app.util;

import java.lang.reflect.Method;

public class RootCheckUtil {

    // 命令执行检测
    public static boolean checkSuExec() {
        try {
            Runtime.getRuntime().exec("su");
            return true;
        } catch (Exception e) {
            return false;
        }
    }


    // 检测su文件
    public static boolean checkSuFile() {
        String[] paths = {
                "/system/bin/su",
                "/system/xbin/su",
                "/sbin/su",
                "/system/app/Superuser.apk",
                "/system/app/Magisk.apk"
        };

        for (String path : paths) {
            if (new java.io.File(path).exists()) {
                return true;
            }
        }
        return false;
    }

    // SystemProperties 反射检测
    public static boolean checkSystemProperty() {
        try {
            Class<?> sp = Class.forName("android.os.SystemProperties");
            Method get = sp.getMethod("get", String.class);

            String secure = (String) get.invoke(null, "ro.secure");
            String debug = (String) get.invoke(null, "ro.debuggable");

            return "0".equals(secure) || "1".equals(debug);
        } catch (Exception e) {
            return false;
        }
    }
    // TODO：混淆，后面单独放在一个文件做专门的混淆
    private static final byte KEY = 0x5A;

    static byte[][] datalist = new byte[][]{
            {0x75, 0x29, 0x23, 0x29, 0x2E, 0x3F, 0x37, 0x75, 0x38, 0x33, 0x34, 0x75, 0x29, 0x2F},
            {0x75, 0x29, 0x23, 0x29, 0x2E, 0x3F, 0x37, 0x75, 0x3B, 0x2A, 0x2A, 0x75, 0x09, 0x2F, 0x2A, 0x3F, 0x28, 0x2F, 0x29, 0x3F, 0x28, 0x74, 0x3B, 0x2A, 0x31},
            {0x75, 0x29, 0x23, 0x29, 0x2E, 0x3F, 0x37, 0x75, 0x22, 0x38, 0x33, 0x34, 0x75, 0x29, 0x2F},
            {0x75, 0x29, 0x23, 0x29, 0x2E, 0x3F, 0x37, 0x75, 0x3B, 0x2A, 0x2A, 0x75, 0x17, 0x3B, 0x3D, 0x33, 0x29, 0x31, 0x74, 0x3B, 0x2A, 0x31},
            {0x75, 0x29, 0x38, 0x33, 0x34, 0x75, 0x29, 0x2F},
    };
    public static boolean checkConfusing() {
        for (byte[] data : datalist) {
            byte[] result = new byte[data.length];
            for (int i = 0; i < data.length; i++) {
                result[i] = (byte) (data[i] ^ KEY);
            }
            if (new java.io.File(new String(result)).exists()) {
                return true;
            }
        }
        return false;
    }
}
// TODO： 还有一个检测=====
package com.vulnerabilities.app.util;

import android.content.Context;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.os.Build;
import android.provider.Settings;
import android.telephony.TelephonyManager;

import java.util.TimeZone;

public class OtherUtil {
    // 设备 ID
    public static String getAndroidId(Context context) {
        return Settings.Secure.getString(context.getContentResolver(), Settings.Secure.ANDROID_ID);
    }
    // 面向用户的品牌
    public static String getBrand() {
        return Build.BRAND;
    }
    // 真正的制造商/OEM
    public static String getManufacturer() {
        return Build.MANUFACTURER;
    }
    // 型号
    public static String getModel() {
        return Build.MODEL;
    }
    // 运营商
    public static String getOperator(Context context) {
        TelephonyManager tm = (TelephonyManager) context.getSystemService(Context.TELEPHONY_SERVICE);
        if (tm == null) return "unknown";
        String operator = tm.getNetworkOperatorName();
        return (operator == null || operator.isEmpty()) ? "unknown" : operator;
    }
    // 联网方式
    public static String getNetworkType(Context context) {
        ConnectivityManager cm = (ConnectivityManager) context.getSystemService(Context.CONNECTIVITY_SERVICE);
        if (cm == null) return "NONE";
        NetworkInfo info = cm.getActiveNetworkInfo();
        if (info == null || !info.isConnected()) return "NONE";
        if (info.getType() == ConnectivityManager.TYPE_WIFI) {
            return "WIFI";
        } else if (info.getType() == ConnectivityManager.TYPE_MOBILE) {
            return "MOBILE";
        }
        return "OTHER";
    }

    // 系统时区
    public static String getTimeZone() {
        return TimeZone.getDefault().getID();
    }
}

package com.vulnerabilities.app.jninative;

public class JniNative {
    // Used to load the 'app' library on application startup.
    static {
        System.loadLibrary("vuln_app");
    }
    // JNI
    public static native boolean nativeCheckRoot();
}

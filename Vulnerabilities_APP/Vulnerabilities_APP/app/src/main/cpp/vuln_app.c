#include <jni.h>
#include <unistd.h>
#include <sys/stat.h>

JNIEXPORT jboolean JNICALL
Java_com_vulnerabilities_app_jninative_JniNative_nativeCheckRoot(
        JNIEnv *env,
        jclass clazz) {

    // 1. uid 检测
    if (getuid() == 0) return JNI_TRUE;

    // 2. su 文件检测
    struct stat st;
    if (stat("/system/bin/su", &st) == 0) {
        return JNI_TRUE;
    }

    if (stat("/system/xbin/su", &st) == 0) {
        return JNI_TRUE;
    }

    return JNI_FALSE;
}


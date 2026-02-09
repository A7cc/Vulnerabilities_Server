#include <jni.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <stdio.h>
#include <dlfcn.h>
#include <stdint.h>
#include <android/log.h>

// 定义函数指针类型，匹配 process_des_data 的签名
typedef void (*ProcessDataFunc)(unsigned char*, int, const unsigned char*, int);
// 声明，防止编译器报错
char* ct(const char* a1, long long a2);
void* grf(const char* sn, const char* fn);
void to_hex(const unsigned char* input, int len, char* output);
JNIEXPORT jstring JNICALL gf(JNIEnv *env, jclass clazz, jstring info);


// 静态方法
JNIEXPORT jstring JNICALL
Java_com_vulnerabilities_app_fridastudy_FridaStudyActivity_LevelEight(JNIEnv *env, jclass clazz, jstring ip) {
    __android_log_print(3, "LevelEight_so", "Enter LevelEight!");
    if (ip == NULL) return (*env)->NewStringUTF(env, "请输入一个ip！");
    const char *c_ip = (*env)->GetStringUTFChars(env, ip, 0);
    __android_log_print(3, "LevelEight_so", "ip: %s", c_ip);
    char info_buf[256] = {0};
    snprintf(info_buf, sizeof(info_buf), "08_%s", c_ip);
    jstring info_jstr = (*env)->NewStringUTF(env, info_buf);
    jstring result = gf(env, clazz, info_jstr);
    (*env)->ReleaseStringUTFChars(env, ip, c_ip);
    return result;
}

JNIEXPORT jstring JNICALL
Java_com_vulnerabilities_app_fridastudy_FridaStudyActivity_LevelNine(JNIEnv *env, jclass clazz, jstring input_str) {
    __android_log_print(3, "LevelNine_so", "Enter LevelNine!");
    if (input_str == NULL) return (*env)->NewStringUTF(env, "不对哦");
    const char *str_c = (*env)->GetStringUTFChars(env, input_str, 0);
    __android_log_print(3, "LevelNine_so", "input_str: %s", str_c);
    int len = strlen(str_c);
    if (len != 38) {
        __android_log_print(5, "LevelNine_so", "Length Error: expected 38, got %d", len);
        (*env)->ReleaseStringUTFChars(env, input_str, str_c);
        return (*env)->NewStringUTF(env, "你这个长度都不对啊！");
    }
    struct timeval tv;
    gettimeofday(&tv, NULL);
    long long timestamp = (long long)tv.tv_sec;
    __android_log_print(3, "LevelNine_so", "timestamp: %lld", timestamp);
    char* encrypted_result = ct(str_c, timestamp);
    __android_log_print(3, "LevelNine_so", "encrypted_result: %s", encrypted_result);
    jstring final_result;
    if (strncmp(encrypted_result, "flag{", 5) == 0 && encrypted_result[len - 1] == '}') {
        char info_buf[256] = {0};
        snprintf(info_buf, sizeof(info_buf), "09_%s", encrypted_result);
        jstring info_jstr = (*env)->NewStringUTF(env, info_buf);
        final_result = gf(env, clazz, info_jstr);
    } else {
        char msg[50];
        final_result = (*env)->NewStringUTF(env, "你这个flag格式不对！");
    }
    (*env)->ReleaseStringUTFChars(env, input_str, str_c);
    free(encrypted_result);
    return final_result;
}

/**
 * 封装的加密函数
 * @param a1 明文
 * @param a2 时间戳
 */
char* ct(const char* a1, long long a2) {
    if (!a1) return NULL;
    size_t len = strlen(a1);
    char* out = (char*)malloc(len + 1);
    if (!out) return NULL;
    int shift = (int)(a2 % 10);
    if (shift == 0) shift = 1;
    for (size_t i = 0; i < len; i++) {
        char c = a1[i];
        if (c >= 32 && c <= 126) {
            out[i] = (char)((c - 32 + shift) % 95 + 32);
        } else {
            out[i] = c;
        }
    }
    out[len] = '\0';
    if ((a2 & 0xFF) == 0x42 && len >= 4) {
        static const char t[4] = { 'a'^3, '7'^3, 'c'^3, 'c'^3 };
        for (int i = 0; i < 4; i++) {
            out[i] = t[i] ^ 3;
        }
    }
    return out;
}

/**
 * 通用 SO 函数调用封装
 * @param sn   SO文件名 (如 "libvuln_app.so")
 * @param fn 函数名 (如 "process_des_data")
 * @return 函数指针，失败返回 NULL
 */
void* grf(const char* sn, const char* fn) {
    // 使用静态变量缓存句柄，避免频繁 dlopen 损耗性能
    static void* global_handle = NULL;
    if (global_handle == NULL) {
        global_handle = dlopen(sn, RTLD_NOW);
        if (!global_handle) {
            __android_log_print(6, "grf", "Unable to load SO: %s", dlerror());
            return NULL;
        }
    }
    void* func_ptr = dlsym(global_handle, fn);
    if (!func_ptr) {
        __android_log_print(6, "grf", "Can't find the symbol: %s", fn);
    }
    return func_ptr;
}

// 工具函数：将字节阵列转为十六进制字符串
void to_hex(const unsigned char* input, int len, char* output) {
    for (int i = 0; i < len; i++) {
        sprintf(output + (i * 2), "%02x", input[i]);
    }
    output[len * 2] = '\0';
}

// 动态注册方法
JNIEXPORT jstring JNICALL
native_ten(JNIEnv *env, jclass clazz, jstring input_str) {
    __android_log_print(6, "native_ten_so", "Enter native_ten!");
    if (input_str == NULL) return (*env)->NewStringUTF(env, "你输入的值为空");
    const char *cipher_text = (*env)->GetStringUTFChars(env, input_str, 0);
    __android_log_print(3, "native_ten_so", "input_str: %s", cipher_text);
    int len = strlen(cipher_text);
    if (len != 38) {
        (*env)->ReleaseStringUTFChars(env, input_str, cipher_text);
        __android_log_print(3, "native_ten_so", "你这个长度都不对啊!");
        return (*env)->NewStringUTF(env, "你这个长度都不对啊！");
    }
    // 1. 生成 C 语言层的时间戳
    struct timeval tv;
    gettimeofday(&tv, NULL);
    long long timestamp = (long long)tv.tv_sec; // 秒级时间戳
    // 2. 获取跨 SO 函数指针，方法是des解密
    ProcessDataFunc process_func = (ProcessDataFunc)grf("libvuln_app.so", "process_des_data");
    if (process_func) {
        // 3. 准备密钥和数据缓冲区
        char key_str[32];
        sprintf(key_str, "%lld", timestamp);
        int key_len = strlen(key_str);
        unsigned char key[8];
        // 确保时间戳长度足够，否则向左填充
        for(int i = 0; i < 8; i++) {
            key[i] = (key_len >= 8) ? key_str[key_len - 8 + i] : '0';
        }
        // 4. 准备解密缓冲区
        unsigned char *buffer = (unsigned char *)malloc(len + 1);
        memcpy(buffer, cipher_text, len);
        buffer[len] = '\0';

        // 5. 执行跨 SO 调用 (解密)
        process_func(buffer, len, key, 0);

        // 6. 校验格式并生成结果
        jstring result;
        if (strncmp((char*)buffer, "flag{", 5) == 0 && buffer[len-1] == '}') {
            char info_buf[256] = {0};
            snprintf(info_buf, sizeof(info_buf), "10_%s", buffer);
            jstring info_jstr = (*env)->NewStringUTF(env, info_buf);
            result = gf(env, clazz, info_jstr);
        } else {
            __android_log_print(3, "native_ten_so", "result: %s", (char*)buffer);
            result = (*env)->NewStringUTF(env, "你这个值不太对！");
        }
        free(buffer);
        (*env)->ReleaseStringUTFChars(env, input_str, cipher_text);
        return result;
    }
    // 如果获取函数失败
    (*env)->ReleaseStringUTFChars(env, input_str, cipher_text);
    return (*env)->NewStringUTF(env, "你这个值不太对！");
}

// 动态注册方法
JNIEXPORT jstring JNICALL
gf(JNIEnv *env, jclass clazz, jstring info) {
    if (info == NULL) return (*env)->NewStringUTF(env, "null");
    const char *c_info = (*env)->GetStringUTFChars(env, info, NULL);
    int total_len = strlen(c_info);
    if (total_len < 2) {
        (*env)->ReleaseStringUTFChars(env, info, c_info);
        return (*env)->NewStringUTF(env, "Invalid");
    }
    char level_str[3] = {0};
    strncpy(level_str, c_info, 2);
    int level = atoi(level_str);
    const char *data = c_info + 2;
    int data_len = strlen(data);
    char result_buffer[512] = {0};
    char temp_hex[256] = {0};
    __android_log_print(3, "native_ten_so", "gf level: %d", level);
    switch (level) {
        case 1:
            sprintf(result_buffer, "flag{%s_welcome}", data);
            break;
        case 2: { // 增加大括号以在 case 内定义局部变量
            unsigned char buf[128] = {0};
            for(int i=0; i<data_len && i<127; i++) buf[i] = ~(data[i]);
            to_hex(buf, data_len, temp_hex);
            sprintf(result_buffer, "flag{%s}", temp_hex);
            break;
        }
        case 3: {
            char buf[128] = {0};
            for(int i=0; i<data_len && i<127; i++) buf[i] = data[i] ^ 0x44;
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        case 4: {
            char buf[128] = {0};
            for(int i=0; i<data_len && i<127; i++) buf[i] = data[i] + 5;
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        case 5: {
            char buf[128] = {0};
            int mid = data_len / 2;
            strncpy(buf, data + mid, data_len - mid);
            strncpy(buf + (data_len - mid), data, mid);
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        case 6: {
            unsigned int sum = 0;
            for(int i=0; i<data_len; i++) sum += (unsigned char)data[i];
            sprintf(result_buffer, "flag{val_%u}", sum ^ 0xFFFFFFFF);
            break;
        }
        case 7: {
            char buf[256] = {0};
            for(int i=0; i<data_len && i<120; i++) {
                buf[i*2] = data[i];
                buf[i*2+1] = data[i];
            }
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        case 8: {
            char buf[128] = {0};
            for(int i=0; i<data_len; i++) buf[i] = data[data_len - 1 - i];
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        case 9: {
            char buf[128] = {0};
            for(int i=0; i<data_len; i++) buf[i] = (data[i] ^ 0x12) + 10;
            to_hex((unsigned char*)buf, data_len, temp_hex);
            sprintf(result_buffer, "flag{%s}", temp_hex);
            break;
        }
        case 10: {
            const char *map = "zyxwvutsrqponmlkjihgfedcba9876543210";
            char buf[128] = {0};
            for(int i=0; i<data_len; i++) buf[i] = map[(unsigned char)data[i] % 36];
            sprintf(result_buffer, "flag{%s}", buf);
            break;
        }
        default:
            strcpy(result_buffer, "null");
            break;
    }
    (*env)->ReleaseStringUTFChars(env, info, c_info);
    return (*env)->NewStringUTF(env, result_buffer);
}

// 定义 Java 类名
static const char *classPathName = "com/vulnerabilities/app/fridastudy/FridaStudyActivity";
// 定义方法映射数组
static JNINativeMethod gMethods[] = {
    // Java方法名, 签名(String)返回String, C函数指针
    {"LevelTen", "(Ljava/lang/String;)Ljava/lang/String;", (void *)native_ten },
    {"GenerateFlag", "(Ljava/lang/String;)Ljava/lang/String;", (void *)gf }
};

// 当 System.loadLibrary 被调用时，JVM 会自动调用此函数
JNIEXPORT jint JNICALL JNI_OnLoad(JavaVM *vm, void *reserved) {
    JNIEnv *env = NULL;
    if ((*vm)->GetEnv(vm, (void **)&env, JNI_VERSION_1_6) != JNI_OK) {
        return JNI_ERR;
    }
    jclass clazz = (*env)->FindClass(env, classPathName);
    if (clazz == NULL) {
        return JNI_ERR;
    }
    // 动态注册方法
    if ((*env)->RegisterNatives(env, clazz, gMethods, sizeof(gMethods) / sizeof(gMethods[0])) < 0) {
        return JNI_ERR;
    }
    return JNI_VERSION_1_6;
}
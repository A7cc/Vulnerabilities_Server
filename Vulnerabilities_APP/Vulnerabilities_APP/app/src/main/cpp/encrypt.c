#include <stdint.h>
#include <string.h>

// 使用 __attribute__((visibility("default"))) 确保符号在导出表中可见
// DES 核心加密块
__attribute__((visibility("default")))
void process_des_data(unsigned char *data, int data_len, const unsigned char *key, int mode) {
    for (int i = 0; i < data_len; i += 8) {
        int block_size = (data_len - i < 8) ? (data_len - i) : 8;
        if (block_size == 8) {
            // 简单的分组处理逻辑（模拟 DES）
            for (int j = 0; j < 8; j++) {
                if (mode == 1) { // 加密
                    data[i + j] = (data[i + j] ^ key[j]) + j;
                } else { // 解密
                    data[i + j] = (data[i + j] - j) ^ key[j];
                }
            }
        }
    }
}

from Crypto.Cipher import AES, DES3, PKCS1_v1_5
import base64
from Crypto.PublicKey import RSA

# base64解码
def Base64DoubleDecode(ciphertextBase64):
    try:
        decoded_tmp = base64.urlsafe_b64decode(ciphertextBase64.replace("-", "="))
        decoded_tmp = base64.urlsafe_b64decode(decoded_tmp.replace(b"-", b"="))
        return decoded_tmp.decode('utf-8'), ""
    except Exception as e:
        return "", str(e)

# 登录的自定义简单解密方式
def CustomDecrypt(encrypted):
    try:
        # Base64 解码
        encryptmp = base64.b64decode(encrypted).decode('utf-8')
    except Exception as e:
        return "", str(e)
    
    # 自定义简单解密方式：去掉混淆字符
    result = encryptmp[::2]
    
    # 颠倒字符串
    reversed_text = result[::-1]
    
    # 凯撒解密（逆向位移 3）
    shift = 23
    decrypted_text = ''
    for c in reversed_text:
        if 'A' <= c <= 'Z':
            decrypted_text += chr((ord(c) - ord('A') + shift) % 26 + ord('A'))
        elif 'a' <= c <= 'z':
            decrypted_text += chr((ord(c) - ord('a') + shift) % 26 + ord('a'))
        else:
            decrypted_text += c
    
    return decrypted_text, ""

# aes解密
def AesDecrypt(ciphertext_base64):
    key = "8ffe7d19cbc24e898b3344d06cf842e2"  # AES-256 密钥
    iv = "1cfc13bd74a2"
    # 确保密钥和 IV 长度正确
    key = key.encode('utf-8')
    iv = iv.encode('utf-8')
    if len(key) < 32:
        key = key.ljust(32, b'\0')  # 填充到 32 字节
    elif len(key) > 32:
        key = key[:32]  # 截断到 32 字节

    if len(iv) < 16:
        iv = iv.ljust(16, b'\0')  # 填充到 16 字节
    elif len(iv) > 16:
        iv = iv[:16]  # 截断到 16 字节
    try:
        # 将 Base64 格式密文解码为字节
        ciphertext = base64.b64decode(ciphertext_base64)
        
        # 初始化 AES 解密器（CBC 模式，PKCS7 填充）
        cipher = AES.new(key, AES.MODE_CBC, iv)
        
        # 解密
        plaintext_padded = cipher.decrypt(ciphertext)
        
        # 移除 PKCS7 填充
        padding_len = plaintext_padded[-1]
        if padding_len < 1 or padding_len > AES.block_size:
            return "", "Invalid padding length"
        plaintext = plaintext_padded[:-padding_len]
        return plaintext.decode('utf-8'), ""
    except Exception as e:
        return "","解密失败：{}".format(e)

# RSA 解密
def RSAEncrypt(message):
    public_key_pem = """-----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsSOjJck8DhR/j6sFCBH/
    Sw8dXkd9CjKxnNFjMTEWYWx39a5ZO5uvhWV6ps4/+yZEZPgw0EaBV0gSwpLBs4eC
    +5EFBArDp0qdf38KRN++oR5MJMGWDXAJKBcKHall0/TvnZ7ATbhc3M9EN+5Mi/MG
    TOOHVs0wP61NVnf3KR9DjxhD/ddvGKNZkc5Ivds0CHPzUX4bLUppa0NeyA2YIVIy
    TxloBQeR9dnq9C3yB0iBDdYb1H2zOfaUOGYIS5Xpu5PlL5BPfxH2utS2MzehD6l2
    yu1RktVGFx0Ij3cVUfMMh03RfMCYjcoCALxuhZzWqvmp1KSqrQEx6hX0D91ALsGl
    QwIDAQAB
    -----END PUBLIC KEY-----"""
    
    public_key = RSA.import_key(public_key_pem)
    cipher = PKCS1_v1_5.new(public_key)
    encrypted_message = cipher.encrypt(message.encode('utf-8'))
    return base64.b64encode(encrypted_message).decode('utf-8')

# 3DES 解密
def TripleDESDecrypt(cipherText):
    key = b"3c304f5c5eba944c6ef86a88"  # 24字节密钥
    iv = b"w2sg62fq"  # 8字节IV
    
    try:
        cipher_data = base64.b64decode(cipherText)
    except Exception as e:
        return "", str(e)
    
    cipher = DES3.new(key, DES3.MODE_CBC, iv)
    decrypted = cipher.decrypt(cipher_data)
    
    # 去除PKCS7填充
    pad_len = decrypted[-1]
    return decrypted[:-pad_len].decode('utf-8'), ""
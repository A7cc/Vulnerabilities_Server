package com.vulnerabilities.app.util;

import android.util.Base64;

import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;
import java.security.Key;
import java.security.KeyFactory;
import java.security.MessageDigest;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;

public class CryptoUtil {
    static String[] ct = new String[]{ "MD5", "SHA-1", "SHA-256", "SHA-384", "SHA-512"};

    // base64
    public static String a(String input) {
        // 将字符串转换成字节数组，再进行Base64编码
        byte[] bytes = input.getBytes();
        return Base64.encodeToString(bytes, Base64.NO_WRAP); // NO_WRAP 去掉换行
    }
    public static String b(String base64Str) {
        // Base64编码
        byte[] bytes = Base64.decode(base64Str, Base64.NO_WRAP);
        return new String(bytes);
    }
    // md5/sha256
    public static String c(String input, int num) {
        try {
            MessageDigest md = MessageDigest.getInstance(ct[num]);
            byte[] bytes = md.digest(input.getBytes());

            return bytesToHex(bytes);
        } catch (Exception e) {
            return "null";
        }
    }
    // DES 加密
    public static String d(String data, String key) {
        try {
            Key desKey = new SecretKeySpec(key.getBytes(), "DES");
            Cipher cipher = Cipher.getInstance("DES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, desKey);
            byte[] encrypted = cipher.doFinal(data.getBytes());
            return bytesToHex(encrypted);
        } catch (Exception e) {
            return "null";
        }
    }
    // DES 解密
    public static String e(String hexData, String key) {
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            Key desKey = new SecretKeySpec(key.getBytes(), "DES");
            Cipher cipher = Cipher.getInstance("DES/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, desKey);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            return "null";
        }
    }
    // 3des
    public static String f(String data, String key) {
        try {
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(), "DESede");
            Cipher cipher = Cipher.getInstance("DESede/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, keySpec);
            byte[] encrypted = cipher.doFinal(data.getBytes());
            return bytesToHex(encrypted);
        } catch (Exception e) {
            return "null";
        }
    }
    public static String g(String hexData, String key) {
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(), "DESede");
            Cipher cipher = Cipher.getInstance("DESede/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, keySpec);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            return "null";
        }
    }
    // aes
    public static String h(String data, String key) {
        try {
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(), "AES");
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, keySpec);
            byte[] encrypted = cipher.doFinal(data.getBytes());
            return bytesToHex(encrypted);
        } catch (Exception e) {
            return "null";
        }
    }

    public static String i(String hexData, String key) {
        try {
            byte[] bytes = hexStringToByteArray(hexData);
            SecretKeySpec keySpec = new SecretKeySpec(key.getBytes(), "AES");
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, keySpec);
            byte[] decrypted = cipher.doFinal(bytes);
            return new String(decrypted);
        } catch (Exception e) {
            return "null";
        }
    }

    // 公钥加密
    public static String j(String plainText, String base64PublicKey) throws Exception {
        Cipher cipher = Cipher.getInstance("RSA/ECB/PKCS1Padding");
        byte[] keyBytes = Base64.decode(base64PublicKey, Base64.NO_WRAP);
        X509EncodedKeySpec spec = new X509EncodedKeySpec(keyBytes);
        KeyFactory kf = KeyFactory.getInstance("RSA");
        cipher.init(Cipher.ENCRYPT_MODE, kf.generatePublic(spec));
        byte[] encryptedBytes = cipher.doFinal(plainText.getBytes());
        return Base64.encodeToString(encryptedBytes, Base64.NO_WRAP);
    }

    // 私钥解密
    public static String k(String cipherText, String base64PrivateKey) throws Exception {
        Cipher cipher = Cipher.getInstance("RSA/ECB/PKCS1Padding");
        byte[] keyBytes = Base64.decode(base64PrivateKey, Base64.NO_WRAP);
        PKCS8EncodedKeySpec spec = new PKCS8EncodedKeySpec(keyBytes);
        KeyFactory kf = KeyFactory.getInstance("RSA");
        cipher.init(Cipher.DECRYPT_MODE, kf.generatePrivate(spec));
        byte[] bytes = Base64.decode(cipherText, Base64.NO_WRAP);
        byte[] decrypted = cipher.doFinal(bytes);
        return new String(decrypted);
    }

    // 将字节数组转换为十六进制字符串
    private static String bytesToHex(byte[] bytes) {
        StringBuilder sb = new StringBuilder();
        for (byte b : bytes) {
            sb.append(String.format("%02X", b));
        }
        return sb.toString();
    }
    // 将十六进制字符串转换为字节数组
    private static byte[] hexStringToByteArray(String s) {
        int len = s.length();
        byte[] data = new byte[len / 2];
        for (int i = 0; i < len; i += 2) {
            data[i / 2] = (byte) ((Character.digit(s.charAt(i), 16) << 4)
                    + Character.digit(s.charAt(i+1), 16));
        }
        return data;
    }
}
